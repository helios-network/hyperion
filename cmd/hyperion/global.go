package main

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"time"

	"cosmossdk.io/math"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/committer"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/keystore"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/provider"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/helios"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/rpcchainlist"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

type Global struct {
	cfg           *Config
	heliosNetwork *helios.Network
	targetNetwork *ethereum.Network

	ethKeyFromAddress gethcommon.Address
	accAddress        cosmostypes.AccAddress
	signerFn          bind.SignerFn
	personalSignFn    keystore.PersonalSignFn
}

func NewGlobal(cfg *Config) *Global {
	return &Global{cfg: cfg}
}

func (g *Global) GetConfig() *Config {
	return g.cfg
}

func (g *Global) GetHeliosNetwork() *helios.Network {
	return g.heliosNetwork
}

func (g *Global) InitHeliosNetwork(linkChainId uint64) (*helios.Network, error) {

	heliosKeyring, err := helios.NewKeyringFromPrivateKey(*g.cfg.heliosPrivKey)
	if err != nil {
		return nil, err
	}

	g.accAddress = heliosKeyring.Addr

	var (
		heliosNetworkCfg = helios.NetworkConfig{
			ChainID:          *g.cfg.heliosChainID,
			HeliosGRPC:       *g.cfg.heliosGRPC,
			TendermintRPC:    *g.cfg.tendermintRPC,
			GasPrice:         *g.cfg.heliosGasPrices,
			Gas:              *g.cfg.heliosGas,
			ValidatorAddress: heliosKeyring.Addr.String(),
		}
	)

	ethKeyFromAddress, signerFn, personalSignFn, err := initEthereumAccountsManagerWithPrivateKey(g.cfg.heliosPrivKey, linkChainId)
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

	ethKeyFromAddress, signerFn, _, _ := initEthereumAccountsManagerWithRandomKey(chainId)

	ethNetwork, _ := ethereum.NewNetwork(gethcommon.HexToAddress("0x0000000000000000000000000000000000000000"), ethKeyFromAddress, signerFn, ethereum.NetworkConfig{
		EthNodeRPCs:           rpcs,
		GasPriceAdjustment:    *g.cfg.ethGasPriceAdjustment,
		MaxGasPrice:           *g.cfg.ethMaxGasPrice,
		PendingTxWaitDuration: *g.cfg.pendingTxWaitDuration,
	})

	gasPrice, err := ethNetwork.GetGasPrice(context.Background())
	if err != nil {
		return "0"
	}
	return gasPrice.String()
}

func (g *Global) TestRpcsAndGetRpcs(chainId uint64, rpcsOptional []string) ([]*hyperiontypes.Rpc, error) {
	rpcs, timeSinceLastUpdate, err := getRpcsFromStorge(chainId)
	if err != nil {
		rpcs = make([]string, 0)
	}

	if timeSinceLastUpdate < 60*time.Minute && len(rpcs) > 0 && timeSinceLastUpdate != 0 {
		rpcList := make([]*hyperiontypes.Rpc, 0)
		for _, rpc := range rpcs {
			rpcList = append(rpcList, &hyperiontypes.Rpc{
				Url:            rpc,
				Reputation:     1,
				LastHeightUsed: 1,
			})
		}
		return rpcList, nil
	}

	if len(rpcsOptional) > 0 {
		notInRpcs := make([]string, 0)
		for _, rpc := range rpcsOptional {
			if !slices.Contains(rpcs, rpc) {
				notInRpcs = append(notInRpcs, rpc)
			}
		}
		if len(notInRpcs) > 0 {
			rpcs = append(rpcs, notInRpcs...)
		}
	}

	rpcChainListFeed := rpcchainlist.NewRpcChainListFeed()
	rpcsFromChainList, err := rpcChainListFeed.QueryRpcs(chainId)
	if err == nil {
		notInRpcs := make([]string, 0)
		for _, rpc := range rpcsFromChainList {
			if !slices.Contains(rpcs, rpc) {
				notInRpcs = append(notInRpcs, rpc)
			}
		}
		if len(notInRpcs) > 0 {
			rpcs = append(rpcs, notInRpcs...)
		}
	}

	rpcList := make([]*hyperiontypes.Rpc, 0)

	for _, rpc := range rpcs {
		rpcList = append(rpcList, &hyperiontypes.Rpc{
			Url:            rpc,
			Reputation:     1,
			LastHeightUsed: 1,
		})
	}

	ethKeyFromAddress, signerFn, _, _ := initEthereumAccountsManagerWithRandomKey(chainId)

	ethNetwork, _ := ethereum.NewNetwork(gethcommon.HexToAddress("0x0000000000000000000000000000000000000000"), ethKeyFromAddress, signerFn, ethereum.NetworkConfig{
		EthNodeRPCs:           rpcList,
		GasPriceAdjustment:    *g.cfg.ethGasPriceAdjustment,
		MaxGasPrice:           *g.cfg.ethMaxGasPrice,
		PendingTxWaitDuration: *g.cfg.pendingTxWaitDuration,
	})

	ethNetwork.TestRpcs(context.Background())

	testedRpcs := ethNetwork.GetRpcs()

	rpcsToSave := make([]string, 0)
	for _, rpc := range testedRpcs {
		rpcsToSave = append(rpcsToSave, rpc.Url)
	}

	updateRpcsToStorge(chainId, rpcsToSave)

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

	ethNetwork, err := ethereum.NewNetwork(hyperionContractAddr, g.ethKeyFromAddress, g.signerFn, ethereum.NetworkConfig{
		EthNodeRPCs:           rpcs,
		GasPriceAdjustment:    *g.cfg.ethGasPriceAdjustment,
		MaxGasPrice:           *g.cfg.ethMaxGasPrice,
		PendingTxWaitDuration: *g.cfg.pendingTxWaitDuration,
	})

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

func (g *Global) CreateNewBlockchainProposal() (uint64, error) {
	title := "Add EthereumSepolia as a counterparty chain"
	description := "Add EthereumSepolia as a counterparty chain"
	contentI := map[string]interface{}{
		"@type": "/helios.hyperion.v1.MsgAddCounterpartyChainParams",
		"counterparty_chain_params": map[string]interface{}{
			"hyperion_id":                      11155111,
			"contract_source_hash":             "05649bbd4cea9b99354570631c1d935aad20faee",
			"bridge_counterparty_address":      "0xB8ed88AcD8b7ac80d9f546F4D75F33DD19dD5746",
			"bridge_chain_id":                  11155111,
			"bridge_chain_name":                "EthereumSepolia",
			"bridge_chain_logo":                "45fa0204dcbb461f9899168a8b56162ecc832919b0c8b81b85f7de2abba408aa",
			"bridge_chain_type":                "evm",
			"signed_valsets_window":            25000,
			"signed_batches_window":            25000,
			"signed_claims_window":             25000,
			"target_batch_timeout":             36000000, // 10 hours
			"target_outgoing_tx_timeout":       36000000, // 10 hours
			"average_block_time":               2000,
			"average_counterparty_block_time":  12000,
			"slash_fraction_valset":            "0.001",
			"slash_fraction_batch":             "0.001",
			"slash_fraction_claim":             "0.001",
			"slash_fraction_conflicting_claim": "0.001",
			"unbond_slashing_valsets_window":   25000,
			"slash_fraction_bad_eth_signature": "0.001",
			"bridge_contract_start_height":     8434556,
			"valset_reward": map[string]interface{}{
				"denom":  "ahelios",
				"amount": "0",
			},
			"initializer": g.ethKeyFromAddress.Hex(),
			"default_tokens": []interface{}{
				map[string]interface{}{
					"token_address_to_denom": map[string]interface{}{
						"denom":                "ahelios",
						"token_address":        "0x507ABEA5D8d39E1880E0fd7620fe433B5797A284",
						"is_cosmos_originated": true,
						"is_concensus_token":   true,
						"symbol":               "HLS",
						"decimals":             18,
					},
					"default_holders": []interface{}{},
				},
				map[string]interface{}{
					"token_address_to_denom": map[string]interface{}{
						"denom":                "ueth",
						"token_address":        "0x959FA4351fA64aad2aE9e55FFd77f341459a012b",
						"is_cosmos_originated": false,
						"is_concensus_token":   false,
						"symbol":               "ETH",
						"decimals":             18,
					},
					"default_holders": []interface{}{},
				},
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

func (g *Global) DeployNewHyperionContract(chainId uint64) (gethcommon.Address, bool) {
	rpcs, err := g.TestRpcsAndGetRpcs(chainId, []string{})
	if err != nil {
		return gethcommon.Address{}, false
	}
	if g.signerFn == nil {
		_, err := g.InitHeliosNetwork(chainId)
		if err != nil {
			return gethcommon.Address{}, false
		}
	}

	gasPrice := g.GetGasPrice(chainId)
	options := []committer.EVMCommitterOption{
		committer.OptionGasPriceFromString(gasPrice),
		committer.OptionGasLimit(5000000),
	}

	fmt.Println("Deploying Hyperion contract...", g.ethKeyFromAddress.Hex())
	address, err, success := ethereum.DeployNewHyperionContract(g.ethKeyFromAddress, g.signerFn, ethereum.NetworkConfig{
		EthNodeRPCs:           rpcs,
		GasPriceAdjustment:    *g.cfg.ethGasPriceAdjustment,
		MaxGasPrice:           *g.cfg.ethMaxGasPrice,
		PendingTxWaitDuration: *g.cfg.pendingTxWaitDuration,
	}, options...)
	if err != nil {
		fmt.Println("Error deploying Hyperion contract:", err)
		return gethcommon.Address{}, false
	}
	if !success {
		return gethcommon.Address{}, false
	}
	return address, true
}

func (g *Global) GetProposal(proposalId uint64) (*govtypes.Proposal, error) {
	proposal, err := (*g.heliosNetwork).GetProposal(context.Background(), proposalId)
	if err != nil {
		return nil, err
	}
	return proposal, nil
}

func (g *Global) VoteOnProposal(proposalId uint64) error {
	err := (*g.heliosNetwork).VoteOnProposal(context.Background(), proposalId, g.accAddress)
	if err != nil {
		return err
	}
	return nil
}
