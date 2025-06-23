package global

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"cosmossdk.io/math"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/committer"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/keystore"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/provider"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/helios"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/rpcchainlist"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
	wrappers "github.com/Helios-Chain-Labs/hyperion/solidity/wrappers/Hyperion.sol"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethcommon "github.com/ethereum/go-ethereum/common"

	keys "github.com/Helios-Chain-Labs/hyperion/orchestrator/keys"
)

type Config struct {
	PrivateKey      string
	HeliosChainID   string
	HeliosGRPC      string
	TendermintRPC   string
	HeliosGasPrices string
	HeliosGas       string

	EthGasPriceAdjustment float64
	EthMaxGasPrice        string
	PendingTxWaitDuration string
}

type Global struct {
	cfg           *Config
	heliosNetwork *helios.Network
	targetNetwork *ethereum.Network

	ethKeyFromAddress gethcommon.Address
	accAddress        cosmostypes.AccAddress
	signerFn          bind.SignerFn
	personalSignFn    keystore.PersonalSignFn

	runners       map[uint64]context.CancelFunc
	orchestrators map[uint64]*orchestrator.Orchestrator
}

func NewGlobal(cfg *Config) *Global {
	return &Global{cfg: cfg, runners: make(map[uint64]context.CancelFunc, 0), orchestrators: make(map[uint64]*orchestrator.Orchestrator, 0)}
}

func (g *Global) GetConfig() *Config {
	return g.cfg
}

func (g *Global) GetHeliosNetwork() *helios.Network {
	if g.heliosNetwork == nil {
		g.InitHeliosNetwork(0)
	}
	return g.heliosNetwork
}

func (g *Global) GetAddress() gethcommon.Address {
	return g.ethKeyFromAddress
}

func (g *Global) GetCosmosAddress() cosmostypes.AccAddress {
	return g.accAddress
}

func (g *Global) GetMinBatchFeeUSD() float64 {
	return 0.0
}

func (g *Global) StartRunnersAtStartUp(runHyperion func(ctx context.Context, g *Global, chainId uint64) error) {
	runners, err := storage.GetRunners()
	if err != nil {
		return
	}
	for _, runner := range runners {
		err := runHyperion(context.Background(), g, uint64(runner["chainId"].(float64)))
		if err != nil {
			fmt.Println("Error running hyperion:", err)
		}
	}
}

func (g *Global) SetRunner(chainId uint64, cancel context.CancelFunc, orchestrator *orchestrator.Orchestrator) {
	storage.SetRunner(chainId)
	g.runners[chainId] = cancel
	g.orchestrators[chainId] = orchestrator
}

func (g *Global) GetRunner(chainId uint64) context.CancelFunc {
	return g.runners[chainId]
}

func (g *Global) GetRunners() map[uint64]context.CancelFunc {
	return g.runners
}

func (g *Global) GetOrchestrator(chainId uint64) *orchestrator.Orchestrator {
	return g.orchestrators[chainId]
}

func (g *Global) GetOrchestrators() map[uint64]*orchestrator.Orchestrator {
	return g.orchestrators
}

func (g *Global) CancelRunner(chainId uint64) {
	g.runners[chainId]()
	delete(g.runners, chainId)
	storage.RemoveRunner(chainId)
	delete(g.orchestrators, chainId)
}

func (g *Global) InitHeliosNetwork(linkChainId uint64) (*helios.Network, error) {

	heliosKeyring, err := helios.NewKeyringFromPrivateKey(g.cfg.PrivateKey)
	if err != nil {
		return nil, err
	}

	g.accAddress = heliosKeyring.Addr

	var (
		heliosNetworkCfg = helios.NetworkConfig{
			ChainID:          g.cfg.HeliosChainID,
			HeliosGRPC:       g.cfg.HeliosGRPC,
			TendermintRPC:    g.cfg.TendermintRPC,
			GasPrice:         g.cfg.HeliosGasPrices,
			Gas:              g.cfg.HeliosGas,
			ValidatorAddress: heliosKeyring.Addr.String(),
		}
	)

	ethKeyFromAddress, signerFn, personalSignFn, err := keys.InitEthereumAccountsManagerWithPrivateKey(&g.cfg.PrivateKey, linkChainId)
	if err != nil {
		fmt.Println("Error initializing ethereum accounts manager:", err)
		return nil, err
	}

	heliosNetwork, err := helios.NewNetworkWithBroadcast(heliosKeyring, personalSignFn, heliosNetworkCfg)
	if err != nil {
		fmt.Println("Error creating helios network:", err)
		return nil, err
	}

	g.ethKeyFromAddress = ethKeyFromAddress
	g.signerFn = signerFn
	g.personalSignFn = personalSignFn
	g.heliosNetwork = &heliosNetwork

	return &heliosNetwork, nil
}

func (g *Global) GetGasPrice(chainId uint64) string {

	rpcs, err := g.TestRpcsAndGetRpcs(chainId, []string{})
	if err != nil {
		return "0"
	}

	ethKeyFromAddress, signerFn, _, _ := keys.InitEthereumAccountsManagerWithRandomKey(chainId)

	ethNetwork, _ := ethereum.NewNetwork(gethcommon.HexToAddress("0x0000000000000000000000000000000000000000"), ethKeyFromAddress, signerFn, ethereum.NetworkConfig{
		EthNodeRPCs:           rpcs,
		GasPriceAdjustment:    g.cfg.EthGasPriceAdjustment,
		MaxGasPrice:           g.cfg.EthMaxGasPrice,
		PendingTxWaitDuration: g.cfg.PendingTxWaitDuration,
	})

	gasPrice, err := ethNetwork.GetGasPrice(context.Background())
	if err != nil {
		return "0"
	}
	return gasPrice.String()
}

func (g *Global) TestRpcsAndGetRpcs(chainId uint64, rpcsOptional []string) ([]*hyperiontypes.Rpc, error) {
	rpcsFromStorage, timeSinceLastUpdate, err := storage.GetRpcsFromStorge(chainId)
	rpcs := make([]*hyperiontypes.Rpc, 0)
	if err == nil && len(rpcsFromStorage) > 0 {
		for _, rpc := range rpcsFromStorage {
			rpcs = append(rpcs, &hyperiontypes.Rpc{
				Url:            rpc,
				Reputation:     3,
				LastHeightUsed: 1,
			})
		}
	}

	staticRpcs, err := storage.GetStaticRpcs(chainId)
	if err != nil {
		staticRpcs = make([]string, 0)
	}

	for _, rpc := range staticRpcs {
		exists := false
		for _, r := range rpcs {
			if r.Url == rpc {
				exists = true
				r.Reputation = 10
				break
			}
		}
		if !exists {
			rpcs = append(rpcs, &hyperiontypes.Rpc{
				Url:            rpc,
				Reputation:     10,
				LastHeightUsed: 1,
			})
		}
	}

	if timeSinceLastUpdate < 60*time.Minute && len(rpcs) > 0 && timeSinceLastUpdate != 0 {
		return rpcs, nil
	}

	if len(rpcsOptional) > 0 {
		notInRpcs := make([]string, 0)
		for _, rpc := range rpcsOptional {
			exists := false
			for _, r := range rpcs {
				if r.Url == rpc {
					exists = true
					break
				}
			}
			if !exists {
				notInRpcs = append(notInRpcs, rpc)
			}
		}
		if len(notInRpcs) > 0 {
			for _, rpc := range notInRpcs {
				rpcs = append(rpcs, &hyperiontypes.Rpc{
					Url:            rpc,
					Reputation:     1,
					LastHeightUsed: 1,
				})
			}
		}
	}

	rpcChainListFeed := rpcchainlist.NewRpcChainListFeed()
	rpcsFromChainList, err := rpcChainListFeed.QueryRpcs(chainId)
	if err == nil {
		notInRpcs := make([]string, 0)
		for _, rpc := range rpcsFromChainList {
			exists := false
			for _, r := range rpcs {
				if r.Url == rpc {
					exists = true
					break
				}
			}
			if !exists {
				notInRpcs = append(notInRpcs, rpc)
			}
		}
		if len(notInRpcs) > 0 {
			for _, rpc := range notInRpcs {
				rpcs = append(rpcs, &hyperiontypes.Rpc{
					Url:            rpc,
					Reputation:     1,
					LastHeightUsed: 1,
				})
			}
		}
	}

	ethKeyFromAddress, signerFn, _, _ := keys.InitEthereumAccountsManagerWithRandomKey(chainId)

	ethNetwork, _ := ethereum.NewNetwork(gethcommon.HexToAddress("0x0000000000000000000000000000000000000000"), ethKeyFromAddress, signerFn, ethereum.NetworkConfig{
		EthNodeRPCs:           rpcs,
		GasPriceAdjustment:    g.cfg.EthGasPriceAdjustment,
		MaxGasPrice:           g.cfg.EthMaxGasPrice,
		PendingTxWaitDuration: g.cfg.PendingTxWaitDuration,
	})

	ethNetwork.TestRpcs(context.Background())

	testedRpcs := ethNetwork.GetRpcs()

	rpcsToSave := make([]string, 0)
	for _, rpc := range testedRpcs {
		rpcsToSave = append(rpcsToSave, rpc.Url)
	}

	storage.UpdateRpcsToStorge(chainId, rpcsToSave)

	rpcFinalList := make([]*hyperiontypes.Rpc, 0)
	for _, rpc := range rpcsToSave {
		rpcFinalList = append(rpcFinalList, &hyperiontypes.Rpc{
			Url: rpc,
		})
	}

	return rpcFinalList, nil
}

func (g *Global) GetRpcs(chainId uint64) ([]*hyperiontypes.Rpc, error) {
	rpcList, err := g.TestRpcsAndGetRpcs(chainId, []string{})
	if err != nil {
		return nil, err
	}
	return rpcList, nil
}

func (g *Global) InitTargetNetwork(counterpartyChainParams *hyperiontypes.CounterpartyChainParams) (*ethereum.Network, error) {

	hyperionContractAddr := gethcommon.HexToAddress(counterpartyChainParams.BridgeCounterpartyAddress)
	rpcs, err := g.TestRpcsAndGetRpcs(counterpartyChainParams.BridgeChainId, []string{})
	if err != nil {
		return nil, err
	}

	settings, err := storage.GetChainSettings(counterpartyChainParams.BridgeChainId, map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	ethKeyFromAddress, signerFn, _, err := keys.InitEthereumAccountsManagerWithPrivateKey(&g.cfg.PrivateKey, counterpartyChainParams.BridgeChainId)
	if err != nil {
		fmt.Println("Error initializing ethereum accounts manager:", err)
		return nil, err
	}

	options := make([]committer.EVMCommitterOption, 0)

	if settings["estimate_gas"] != nil && !settings["estimate_gas"].(bool) {
		gasPrice := strconv.FormatInt(committer.ParseGasPrice(settings["eth_gas_price"].(string)), 10)
		options = append(options, committer.OptionGasPriceFromString(gasPrice))
		options = append(options, committer.OptionGasLimit(5000000))
		options = append(options, committer.OptionEstimateGas(false))
	} else {
		gasPrice := g.GetGasPrice(counterpartyChainParams.BridgeChainId)
		options = append(options, committer.OptionGasPriceFromString(gasPrice))
		options = append(options, committer.OptionGasLimit(5000000))
	}

	ethNetwork, err := ethereum.NewNetwork(hyperionContractAddr, ethKeyFromAddress, signerFn, ethereum.NetworkConfig{
		EthNodeRPCs:           rpcs,
		GasPriceAdjustment:    g.cfg.EthGasPriceAdjustment,
		MaxGasPrice:           g.cfg.EthMaxGasPrice,
		PendingTxWaitDuration: g.cfg.PendingTxWaitDuration,
	}, options...)

	if err != nil {
		return nil, err
	}
	return &ethNetwork, nil
}

func (g *Global) GetEVMProvider(chainId uint64) (provider.EVMProvider, error) {
	rpcs, err := g.TestRpcsAndGetRpcs(chainId, []string{})
	if err != nil {
		return nil, err
	}
	return provider.NewEVMProvider(rpcs), nil
}

func (g *Global) CreateNewBlockchainProposal(title string, description string, counterpartyChainParams *hyperiontypes.CounterpartyChainParams) (uint64, error) {
	_, err := g.InitHeliosNetwork(0)
	if err != nil {
		return 0, err
	}

	hash := sha256.Sum256([]byte(wrappers.HyperionBin))
	hashString := hex.EncodeToString(hash[:])

	contentI := map[string]interface{}{
		"@type": "/helios.hyperion.v1.MsgAddCounterpartyChainParams",
		"counterparty_chain_params": map[string]interface{}{
			"hyperion_id":                      counterpartyChainParams.HyperionId,
			"contract_source_hash":             hashString,
			"bridge_counterparty_address":      counterpartyChainParams.BridgeCounterpartyAddress,
			"bridge_chain_id":                  counterpartyChainParams.BridgeChainId,
			"bridge_chain_name":                counterpartyChainParams.BridgeChainName,
			"bridge_chain_logo":                counterpartyChainParams.BridgeChainLogo,
			"bridge_chain_type":                "evm",
			"signed_valsets_window":            25000,
			"signed_batches_window":            25000,
			"signed_claims_window":             25000,
			"target_batch_timeout":             36000000, // 10 hours
			"target_outgoing_tx_timeout":       36000000, // 10 hours
			"average_block_time":               2000,
			"average_counterparty_block_time":  counterpartyChainParams.AverageCounterpartyBlockTime,
			"slash_fraction_valset":            "0.001",
			"slash_fraction_batch":             "0.001",
			"slash_fraction_claim":             "0.001",
			"slash_fraction_conflicting_claim": "0.001",
			"unbond_slashing_valsets_window":   25000,
			"slash_fraction_bad_eth_signature": "0.001",
			"bridge_contract_start_height":     counterpartyChainParams.BridgeContractStartHeight,
			"valset_reward": map[string]interface{}{
				"denom":  "ahelios",
				"amount": "0",
			},
			"initializer":    g.ethKeyFromAddress.Hex(),
			"default_tokens": []interface{}{
				// map[string]interface{}{
				// 	"token_address_to_denom": map[string]interface{}{
				// 		"denom":                "ahelios",
				// 		"token_address":        "0x507ABEA5D8d39E1880E0fd7620fe433B5797A284",
				// 		"is_cosmos_originated": true,
				// 		"is_concensus_token":   true,
				// 		"symbol":               "HLS",
				// 		"decimals":             18,
				// 	},
				// 	"default_holders": []interface{}{},
				// },
				// map[string]interface{}{
				// 	"token_address_to_denom": map[string]interface{}{
				// 		"denom":                "ueth",
				// 		"token_address":        "0x959FA4351fA64aad2aE9e55FFd77f341459a012b",
				// 		"is_cosmos_originated": false,
				// 		"is_concensus_token":   false,
				// 		"symbol":               "ETH",
				// 		"decimals":             18,
				// 	},
				// 	"default_holders": []interface{}{},
				// },
			},
			"rpcs":                       []interface{}{},
			"offset_valset_nonce":        0,
			"min_call_external_data_gas": 10000000,
			"paused":                     false,
		},
	}
	content, _ := json.Marshal(contentI)
	proposalId, err := (*g.heliosNetwork).SendProposal(context.Background(), title, description, string(content), g.accAddress, math.NewInt(1000000000000000000))
	if err != nil {
		return 0, err
	}
	return proposalId, nil
}

func (g *Global) DeployNewHyperionContract(chainId uint64) (gethcommon.Address, uint64, bool) {
	rpcs, err := g.TestRpcsAndGetRpcs(chainId, []string{})
	if err != nil {
		return gethcommon.Address{}, 0, false
	}
	_, err = g.InitHeliosNetwork(chainId)
	if err != nil {
		return gethcommon.Address{}, 0, false
	}

	gasPrice := g.GetGasPrice(chainId)
	options := []committer.EVMCommitterOption{
		committer.OptionGasPriceFromString(gasPrice),
		committer.OptionGasLimit(5000000),
	}

	fmt.Println("Deploying Hyperion contract...", g.ethKeyFromAddress.Hex())
	address, blockNumber, err, success := ethereum.DeployNewHyperionContract(g.ethKeyFromAddress, g.signerFn, ethereum.NetworkConfig{
		EthNodeRPCs:           rpcs,
		GasPriceAdjustment:    g.cfg.EthGasPriceAdjustment,
		MaxGasPrice:           g.cfg.EthMaxGasPrice,
		PendingTxWaitDuration: g.cfg.PendingTxWaitDuration,
	}, options...)
	if err != nil {
		fmt.Println("Error deploying Hyperion contract:", err)
		return gethcommon.Address{}, 0, false
	}
	if !success {
		return gethcommon.Address{}, 0, false
	}
	return address, blockNumber, true
}

func (g *Global) InitializeHyperionContractWithDefaultValset(chainId uint64) (uint64, error) {
	_, err := g.InitHeliosNetwork(chainId)
	if err != nil {
		return 0, err
	}
	hyperionContractInfo, err := storage.GetHyperionContractInfo(chainId)
	if err != nil {
		return 0, err
	}

	targetNetwork, err := g.InitTargetNetwork(&hyperiontypes.CounterpartyChainParams{
		BridgeChainId:             chainId,
		BridgeCounterpartyAddress: hyperionContractInfo["hyperionAddress"].(string),
	})
	if err != nil {
		return 0, err
	}

	// default valset only for initial deployment
	hyperionIdHash := common.HexToHash(strconv.FormatUint(chainId, 16))
	powerThreshold := big.NewInt(1431655765)
	validators := make([]gethcommon.Address, 1)
	powers := make([]*big.Int, 1)
	validators[0] = g.ethKeyFromAddress
	powers[0] = big.NewInt(2147483647)

	_, blockNumber, err := (*targetNetwork).SendInitializeBlockchainTx(context.Background(), g.ethKeyFromAddress, hyperionIdHash, powerThreshold, validators, powers)

	if err != nil {
		lastEventNonce, err2 := (*targetNetwork).GetLastEventNonce(context.Background())
		if err2 != nil {
			return 0, err2
		}

		if lastEventNonce.Uint64() == 1 {
			lastEventHeight, err2 := (*targetNetwork).GetLastEventHeight(context.Background())
			if err2 != nil {
				return 0, err2
			}
			blockNumber = lastEventHeight.Uint64()
		} else {
			return 0, err
		}
	}

	storage.UpdateHyperionContractInfo(chainId, hyperionContractInfo["hyperionAddress"].(string), map[string]interface{}{
		"initializedAtBlockNumber": blockNumber,
	})

	return blockNumber, nil
}

func (g *Global) GetProposal(proposalId uint64) (*govtypes.Proposal, error) {
	proposal, err := (*g.heliosNetwork).GetProposal(context.Background(), proposalId)
	if err != nil {
		return nil, err
	}
	return proposal, nil
}

func (g *Global) VoteOnProposal(proposalId uint64) error {
	_, err := g.InitHeliosNetwork(0)
	if err != nil {
		return err
	}
	err = (*g.heliosNetwork).VoteOnProposal(context.Background(), proposalId, g.accAddress)
	if err != nil {
		return err
	}
	return nil
}
