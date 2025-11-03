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
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/helios"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/rpcs"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
	wrappers "github.com/Helios-Chain-Labs/hyperion/solidity/wrappers/Hyperion.sol"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
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

	ethKeyFromAddress gethcommon.Address
	accAddress        cosmostypes.AccAddress

	runners                   map[uint64]context.CancelFunc
	orchestrators             map[uint64]*orchestrator.Orchestrator
	lastTimeResetHeliosClient time.Time
}

func NewGlobal(cfg *Config) *Global {
	return &Global{cfg: cfg, runners: make(map[uint64]context.CancelFunc, 0), orchestrators: make(map[uint64]*orchestrator.Orchestrator, 0), lastTimeResetHeliosClient: time.Now()}
}

func (g *Global) GetConfig() *Config {
	return g.cfg
}

func (g *Global) GetHeliosNetwork() *helios.Network {
	if g.heliosNetwork == nil {
		_, err := g.InitHeliosNetwork()
		if err != nil {
			return nil
		}
	}
	return g.heliosNetwork
}

func (g *Global) ResetHeliosClient() {
	if g.heliosNetwork != nil {
		if time.Since(g.lastTimeResetHeliosClient) < 10*time.Second {
			fmt.Println("Helios client already reset in the last 10 seconds")
			return
		}
		heliosNetwork := *g.heliosNetwork
		heliosNetwork.Close()
		_, err := g.InitHeliosNetwork()
		if err != nil {
			fmt.Println("Error resetting helios client:", err)
		}
		g.lastTimeResetHeliosClient = time.Now()
	}
}

func (g *Global) GetAddress() gethcommon.Address {
	return g.ethKeyFromAddress
}

func (g *Global) GetCosmosAddress() cosmostypes.AccAddress {
	return g.accAddress
}

func (g *Global) GetMinBatchFeeHLS(chainId uint64) float64 {
	hyperionSettings, err := storage.GetChainSettings(chainId)
	if err != nil {
		return 0.0
	}
	if hyperionSettings["min_batch_fee_hls"] == nil {
		return 0.0
	}
	return hyperionSettings["min_batch_fee_hls"].(float64)
}

func (g *Global) GetMinTxFeeHLS(chainId uint64) float64 {
	hyperionSettings, err := storage.GetChainSettings(chainId)
	if err != nil {
		return 0.0
	}
	if hyperionSettings["min_tx_fee_hls"] == nil {
		return 0.0
	}
	return hyperionSettings["min_tx_fee_hls"].(float64)
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
	if orchestrator != nil {
		g.orchestrators[chainId] = orchestrator
	}
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

func (g *Global) InitHeliosNetwork() (*helios.Network, error) {

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

	ethKeyFromAddress, _, _, err := keys.InitEthereumAccountsManagerWithPrivateKey(&g.cfg.PrivateKey, 1)
	if err != nil {
		fmt.Println("Error initializing ethereum accounts manager:", err)
		return nil, err
	}

	heliosNetwork, err := helios.NewNetworkWithBroadcast(heliosKeyring, heliosNetworkCfg)
	if err != nil {
		fmt.Println("Error creating helios network:", err)
		return nil, err
	}

	g.ethKeyFromAddress = ethKeyFromAddress
	g.heliosNetwork = &heliosNetwork

	return &heliosNetwork, nil
}

func (g *Global) GetAnonymousEVMNetwork(chainId uint64, rpc *rpcs.Rpc, options ...committer.EVMCommitterOption) (*ethereum.Network, error) {
	ethKeyFromAddress, signerFn, personalSignFn, err := keys.InitEthereumAccountsManagerWithRandomKey(chainId)
	if err != nil {
		return nil, err
	}

	ethNetwork, err := ethereum.NewNetwork(gethcommon.HexToAddress("0x0000000000000000000000000000000000000000"), ethKeyFromAddress, signerFn, personalSignFn, ethereum.NetworkConfig{
		EthNodeRPC:            rpc,
		GasPriceAdjustment:    g.cfg.EthGasPriceAdjustment,
		MaxGasPrice:           g.cfg.EthMaxGasPrice,
		PendingTxWaitDuration: g.cfg.PendingTxWaitDuration,
	}, options...)
	if err != nil {
		return nil, err
	}
	return &ethNetwork, nil
}

func (g *Global) GetAnonymousEVMNetworks(chainId uint64, rpcs []*rpcs.Rpc) ([]*ethereum.Network, error) {
	ethNetworks := make([]*ethereum.Network, 0)
	for _, rpc := range rpcs {
		ethNetwork, err := g.GetAnonymousEVMNetwork(chainId, rpc)
		if err != nil {
			fmt.Println("Error getting EVM network:", err, "for rpc:", rpc.Url)
			continue
		}
		ethNetworks = append(ethNetworks, ethNetwork)
	}
	return ethNetworks, nil
}

func (g *Global) GetEVMNetwork(counterpartyChainParams *hyperiontypes.CounterpartyChainParams, rpc *rpcs.Rpc) (*ethereum.Network, error) {
	hyperionContractAddr := gethcommon.HexToAddress(counterpartyChainParams.BridgeCounterpartyAddress)

	settings, err := storage.GetChainSettings(counterpartyChainParams.BridgeChainId)
	if err != nil {
		return nil, err
	}

	ethKeyFromAddress, signerFn, personalSignFn, err := keys.InitEthereumAccountsManagerWithPrivateKey(&g.cfg.PrivateKey, counterpartyChainParams.BridgeChainId)
	if err != nil {
		fmt.Println("Error initializing ethereum accounts manager:", err)
		return nil, err
	}

	options := make([]committer.EVMCommitterOption, 0)
	gasLimit := 5000000

	if settings["gas_limit"] != nil {
		gasLimit = int(settings["gas_limit"].(float64))
	}

	if settings["estimate_gas"] != nil && !settings["estimate_gas"].(bool) {
		gasPrice := strconv.FormatInt(committer.ParseGasPrice(settings["eth_gas_price"].(string)), 10)
		options = append(options, committer.OptionGasPriceFromString(gasPrice))
		options = append(options, committer.OptionGasLimit(uint64(gasLimit)))
		options = append(options, committer.OptionEstimateGas(false))
	} else {
		gasPrice := g.GetGasPrice(counterpartyChainParams.BridgeChainId)
		options = append(options, committer.OptionGasPriceFromString(gasPrice))
		options = append(options, committer.OptionGasLimit(uint64(gasLimit)))
	}

	ethNetwork, err := ethereum.NewNetwork(hyperionContractAddr, ethKeyFromAddress, signerFn, personalSignFn, ethereum.NetworkConfig{
		EthNodeRPC:            rpc,
		GasPriceAdjustment:    g.cfg.EthGasPriceAdjustment,
		MaxGasPrice:           g.cfg.EthMaxGasPrice,
		PendingTxWaitDuration: g.cfg.PendingTxWaitDuration,
	}, options...)

	if err != nil {
		return nil, err
	}
	return &ethNetwork, nil
}

func (g *Global) GetEVMNetworks(counterpartyChainParams *hyperiontypes.CounterpartyChainParams, rpcs []*rpcs.Rpc) ([]*ethereum.Network, error) {
	ethNetworks := make([]*ethereum.Network, 0)
	for _, rpc := range rpcs {
		ethNetwork, err := g.GetEVMNetwork(counterpartyChainParams, rpc)
		if err != nil {
			fmt.Println("Error getting EVM network:", err, "for rpc:", rpc.Url)
			continue
		}
		ethNetworks = append(ethNetworks, ethNetwork)
	}
	return ethNetworks, nil
}

func (g *Global) GetGasPrice(chainId uint64) string {
	rpcs, _, err := storage.GetRpcsFromStorge(chainId)
	if err != nil {
		return "0"
	}

	if len(rpcs) == 0 {
		return "0"
	}

	ethNetwork, err := g.GetAnonymousEVMNetwork(chainId, rpcs[0])
	if err != nil {
		return "0"
	}

	gasPrice, err := (*ethNetwork).GetGasPrice(context.Background())
	if err != nil {
		return "0"
	}
	return gasPrice.String()
}

func (g *Global) TestChainListRpcsAndSaveForChain(chainId uint64) ([]*rpcs.Rpc, error) {
	rpcChainListFeed := rpcs.NewRpcChainListFeed()
	rpcsFromChainListStrings, err := rpcChainListFeed.QueryRpcs(chainId)
	if err != nil {
		return nil, err
	}
	rpcsFromChainList := make([]*rpcs.Rpc, 0)
	for _, rpc := range rpcsFromChainListStrings {
		rpcsFromChainList = append(rpcsFromChainList, &rpcs.Rpc{
			Url: rpc,
		})
	}
	ethNetworks, err := g.GetEVMNetworks(&hyperiontypes.CounterpartyChainParams{
		BridgeChainId:             chainId,
		BridgeCounterpartyAddress: "0x0000000000000000000000000000000000000000",
		BridgeChainName:           "Simulation",
	}, rpcsFromChainList)
	if err != nil {
		return nil, err
	}

	for _, ethNetwork := range ethNetworks {
		ok := (*ethNetwork).TestRpc(context.Background())
		if !ok {
			fmt.Println("Error testing rpc:", (*ethNetwork).GetRpc().Url)
			continue
		}
		rpc := (*ethNetwork).GetRpc()
		rpc.IsTested = true
		rpc.Usages = append(rpc.Usages, rpcs.RpcUsage{
			Success: true,
			Error:   "",
			Time:    time.Now(),
		})
		storage.AddRpcToStorge(chainId, rpc)
		rpcsFromChainList = append(rpcsFromChainList, rpc)
	}
	return rpcsFromChainList, nil
}

func (g *Global) InitTargetNetworks(counterpartyChainParams *hyperiontypes.CounterpartyChainParams) ([]*ethereum.Network, error) {
	rpcs, _, err := storage.GetRpcsFromStorge(counterpartyChainParams.BridgeChainId)
	if err != nil {
		return nil, err
	}

	if len(rpcs) == 0 {
		return nil, fmt.Errorf("no rpcs found for chainId: %d", counterpartyChainParams.BridgeChainId)
	}

	ethNetworks, err := g.GetEVMNetworks(counterpartyChainParams, rpcs)
	if err != nil {
		return nil, err
	}
	return ethNetworks, nil
}

// func (g *Global) TestRpcsAndGetRpcs(chainId uint64, rpcsOptional []string) ([]*hyperiontypes.Rpc, error) {
// 	rpcsFromStorage, timeSinceLastUpdate, err := storage.GetRpcsFromStorge(chainId)
// 	listOfRpcs := make([]*hyperiontypes.Rpc, 0)
// 	if err == nil && len(rpcsFromStorage) > 0 {
// 		for _, rpc := range rpcsFromStorage {
// 			listOfRpcs = append(listOfRpcs, &hyperiontypes.Rpc{
// 				Url:            rpc,
// 				Reputation:     3,
// 				LastHeightUsed: 1,
// 			})
// 		}
// 	}

// 	staticRpcs, err := storage.GetStaticRpcs(chainId)
// 	if err != nil {
// 		staticRpcs = make([]storage.Rpc, 0)
// 	}

// 	for _, rpc := range staticRpcs {
// 		exists := false
// 		for _, r := range listOfRpcs {
// 			if r.Url == rpc.Url {
// 				exists = true
// 				r.Reputation = 10
// 				break
// 			}
// 		}
// 		if !exists {
// 			listOfRpcs = append(listOfRpcs, &hyperiontypes.Rpc{
// 				Url:            rpc.Url,
// 				Reputation:     10,
// 				LastHeightUsed: 1,
// 			})
// 		}
// 	}

// 	if timeSinceLastUpdate < 60*time.Minute && len(listOfRpcs) > 0 && timeSinceLastUpdate != 0 {
// 		return listOfRpcs, nil
// 	}

// 	if len(rpcsOptional) > 0 {
// 		notInRpcs := make([]string, 0)
// 		for _, rpc := range rpcsOptional {
// 			exists := false
// 			for _, r := range listOfRpcs {
// 				if r.Url == rpc {
// 					exists = true
// 					break
// 				}
// 			}
// 			if !exists {
// 				notInRpcs = append(notInRpcs, rpc)
// 			}
// 		}
// 		if len(notInRpcs) > 0 {
// 			for _, rpc := range notInRpcs {
// 				listOfRpcs = append(listOfRpcs, &hyperiontypes.Rpc{
// 					Url:            rpc,
// 					Reputation:     1,
// 					LastHeightUsed: 1,
// 				})
// 			}
// 		}
// 	}

// 	isStaticRpcOnly := false
// 	settings, err := storage.GetChainSettings(chainId, map[string]interface{}{})
// 	if err == nil {
// 		if settings["static_rpc_only"] != nil && settings["static_rpc_only"].(bool) {
// 			isStaticRpcOnly = true
// 		}
// 	}

// 	if !isStaticRpcOnly {
// 		// search rpc from chainlist
// 		rpcChainListFeed := rpcs.NewRpcChainListFeed()
// 		rpcsFromChainList, err := rpcChainListFeed.QueryRpcs(chainId)
// 		if err == nil {
// 			notInRpcs := make([]string, 0)
// 			for _, rpc := range rpcsFromChainList {
// 				exists := false
// 				for _, r := range listOfRpcs {
// 					if r.Url == rpc {
// 						exists = true
// 						break
// 					}
// 				}
// 				if !exists {
// 					notInRpcs = append(notInRpcs, rpc)
// 				}
// 			}
// 			if len(notInRpcs) > 0 {
// 				for _, rpc := range notInRpcs {
// 					listOfRpcs = append(listOfRpcs, &hyperiontypes.Rpc{
// 						Url:            rpc,
// 						Reputation:     1,
// 						LastHeightUsed: 1,
// 					})
// 				}
// 			}
// 		}
// 	}

// 	ethKeyFromAddress, signerFn, personalSignFn, _ := keys.InitEthereumAccountsManagerWithRandomKey(chainId)

// 	ethNetwork, _ := ethereum.NewNetwork(gethcommon.HexToAddress("0x0000000000000000000000000000000000000000"), ethKeyFromAddress, signerFn, personalSignFn, ethereum.NetworkConfig{
// 		EthNodeRPCs:           listOfRpcs,
// 		GasPriceAdjustment:    g.cfg.EthGasPriceAdjustment,
// 		MaxGasPrice:           g.cfg.EthMaxGasPrice,
// 		PendingTxWaitDuration: g.cfg.PendingTxWaitDuration,
// 	})

// 	ethNetwork.TestRpcs(context.Background())

// 	testedRpcs := ethNetwork.GetRpcs()

// 	rpcsToSave := make([]string, 0)
// 	for _, rpc := range testedRpcs {
// 		rpcsToSave = append(rpcsToSave, rpc.Url)
// 	}

// 	storage.UpdateRpcsToStorge(chainId, rpcsToSave)

// 	rpcFinalList := make([]*hyperiontypes.Rpc, 0)
// 	for _, rpc := range rpcsToSave {
// 		rpcFinalList = append(rpcFinalList, &hyperiontypes.Rpc{
// 			Url: rpc,
// 		})
// 	}

// 	return rpcFinalList, nil
// }

// func (g *Global) GetRpcs(chainId uint64) ([]*hyperiontypes.Rpc, error) {
// 	rpcList, err := g.TestRpcsAndGetRpcs(chainId, []string{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return rpcList, nil
// }

// func (g *Global) GetEVMProvider(chainId uint64) (provider.EVMProvider, error) {
// 	rpcs, err := g.TestRpcsAndGetRpcs(chainId, []string{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return provider.NewEVMProvider(rpcs), nil
// }

func (g *Global) ProposeHyperionUpdate(title string, description string, counterpartyChainParams *hyperiontypes.CounterpartyChainParams) (uint64, error) {
	heliosNetwork := g.GetHeliosNetwork()

	if heliosNetwork == nil {
		return 0, fmt.Errorf("helios network not initialized")
	}

	hash := sha256.Sum256([]byte(wrappers.HyperionBin))
	hashString := hex.EncodeToString(hash[:])

	contentI := map[string]interface{}{
		"@type":                        "/helios.hyperion.v1.MsgUpdateChainSmartContract",
		"chain_id":                     counterpartyChainParams.BridgeChainId,
		"bridge_contract_address":      counterpartyChainParams.BridgeCounterpartyAddress,
		"bridge_contract_start_height": counterpartyChainParams.BridgeContractStartHeight,
		"contract_source_hash":         hashString,
		"first_orchestrator_address":   g.ethKeyFromAddress.Hex(),
	}
	content, _ := json.Marshal(contentI)
	proposalId, err := (*g.heliosNetwork).SendProposal(context.Background(), title, description, string(content), g.accAddress, math.NewInt(1000000000000000000))
	if err != nil {
		return 0, err
	}
	return proposalId, nil
}

func (g *Global) ProposeMsgAddOneWhitelistedAddress(title string, description string, hyperionId uint64, address string) (uint64, error) {
	heliosNetwork := g.GetHeliosNetwork()

	if heliosNetwork == nil {
		return 0, fmt.Errorf("helios network not initialized")
	}

	contentI := map[string]interface{}{
		"@type":       "/helios.hyperion.v1.MsgAddOneWhitelistedAddress",
		"hyperion_id": hyperionId,
		"address":     address,
	}
	content, _ := json.Marshal(contentI)
	proposalId, err := (*g.heliosNetwork).SendProposal(context.Background(), title, description, string(content), g.accAddress, math.NewInt(1000000000000000000))
	if err != nil {
		return 0, err
	}
	return proposalId, nil
}

func (g *Global) CreateNewBlockchainProposal(title string, description string, counterpartyChainParams *hyperiontypes.CounterpartyChainParams) (uint64, error) {
	heliosNetwork := g.GetHeliosNetwork()

	if heliosNetwork == nil {
		return 0, fmt.Errorf("helios network not initialized")
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
	fmt.Println("contentI: ", contentI)
	content, _ := json.Marshal(contentI)
	proposalId, err := (*g.heliosNetwork).SendProposal(context.Background(), title, description, string(content), g.accAddress, math.NewInt(1000000000000000000))
	if err != nil {
		return 0, err
	}
	return proposalId, nil
}

func (g *Global) DeployNewHyperionContract(chainId uint64) (gethcommon.Address, uint64, bool) {
	rpcs, _, err := storage.GetRpcsFromStorge(chainId)
	if err != nil {
		return gethcommon.Address{}, 0, false
	}
	if len(rpcs) == 0 {
		// setup rpcs
		fmt.Println("No rpcs found for chainId:", chainId, "in global.DeployNewHyperionContract")
		return gethcommon.Address{}, 0, false
	}
	heliosNetwork := g.GetHeliosNetwork()
	if heliosNetwork == nil {
		return gethcommon.Address{}, 0, false
	}

	ethKeyFromAddress, signerFn, _, err := keys.InitEthereumAccountsManagerWithPrivateKey(&g.cfg.PrivateKey, chainId)
	if err != nil {
		fmt.Println("Error initializing ethereum accounts manager:", err)
		return gethcommon.Address{}, 0, false
	}

	gasPrice := g.GetGasPrice(chainId)

	gasLimit := 5000000
	settings, err := storage.GetChainSettings(chainId)
	if err != nil {
		return gethcommon.Address{}, 0, false
	}
	if settings["gas_limit"] != nil {
		gasLimit = int(settings["gas_limit"].(float64))
	}
	fmt.Println("gasLimit:", gasLimit)
	options := []committer.EVMCommitterOption{
		committer.OptionGasPriceFromString(gasPrice),
		committer.OptionGasLimit(uint64(gasLimit)),
	}

	fmt.Println("Deploying Hyperion contract...", ethKeyFromAddress.Hex())
	address, blockNumber, err, success := ethereum.DeployNewHyperionContract(g.ethKeyFromAddress, signerFn, ethereum.NetworkConfig{
		EthNodeRPC:            rpcs[0],
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
	heliosNetwork := g.GetHeliosNetwork()
	if heliosNetwork == nil {
		return 0, fmt.Errorf("helios network not initialized")
	}
	hyperionContractInfo, err := storage.GetHyperionContractInfo(chainId)
	if err != nil {
		return 0, err
	}

	targetNetworks, err := g.InitTargetNetworks(&hyperiontypes.CounterpartyChainParams{
		BridgeChainId:             chainId,
		BridgeCounterpartyAddress: hyperionContractInfo["hyperionAddress"].(string),
	})
	if err != nil {
		return 0, err
	}
	if len(targetNetworks) == 0 {
		return 0, fmt.Errorf("no target networks found for chainId: %d", chainId)
	}
	targetNetwork := targetNetworks[0]

	// default valset only for initial deployment
	hyperionIdHash := gethcommon.HexToHash(strconv.FormatUint(chainId, 16))
	powerThreshold := big.NewInt(1431655765) // ≈ 33.3% of 4294967295 who is int32 normalized power
	validators := make([]gethcommon.Address, 1)
	powers := make([]*big.Int, 1)
	validators[0] = g.ethKeyFromAddress
	powers[0] = big.NewInt(4294967295) // ≈ 100% of 4294967295 for the first validator

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

	storage.UpdateHyperionContractInfo(chainId, map[string]interface{}{
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
	heliosNetwork := g.GetHeliosNetwork()
	if heliosNetwork == nil {
		return fmt.Errorf("helios network not initialized")
	}
	err := (*g.heliosNetwork).VoteOnProposal(context.Background(), proposalId, g.accAddress)
	if err != nil {
		return err
	}
	return nil
}

func (g *Global) VoteOnProposalWithOption(proposalId uint64, voteOption govtypes.VoteOption) error {
	heliosNetwork := g.GetHeliosNetwork()
	if heliosNetwork == nil {
		return fmt.Errorf("helios network not initialized")
	}
	err := (*g.heliosNetwork).VoteOnProposalWithOption(context.Background(), proposalId, g.accAddress, voteOption)
	if err != nil {
		return err
	}
	return nil
}
