package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/helios"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
	"github.com/ethereum/go-ethereum/common"
	cli "github.com/jawher/mow.cli"
	"github.com/pkg/errors"
	"github.com/xlab/closer"
	log "github.com/xlab/suplog"
)

// queryCmdSubset contains actions that query stuff from Hyperion module
// and the Ethereum contract
//
// $ hyperion q
func queryCmdSubset(cmd *cli.Cmd) {
	cmd.Command("unset-orchestrator", "Unset the orchestrator addresses for a chain", unsetOrchestratorCmd)
	cmd.Command("set-orchestrator", "Set the orchestrator addresses for a chain", setOrchestratorCmd)
	cmd.Command("list-operative-chains", "List the chains that are operative", listOperativeChainsCmd)
	cmd.Command("initialize-blockchain", "Initialize the blockchain for a chain", initializeBlockchainCmd)
}

func unsetOrchestratorCmd(cmd *cli.Cmd) {

	cfg := initConfig(cmd)

	chainId := cmd.String(cli.StringOpt{
		Name:   "chain-id",
		Desc:   "Specify which chain to query",
		EnvVar: "CHAIN_ID",
		Value:  "",
	})

	cmd.Action = func() {

		if *chainId == "" {
			orShutdown(errors.New("chain-id is required"))
		}

		chainIdUint, err := strconv.ParseUint(*chainId, 10, 64)
		orShutdown(errors.Wrap(err, "failed to parse chain id"))

		ctx, cancelFn := context.WithCancel(context.Background())
		closer.Bind(cancelFn)

		heliosNetwork, err := getHeliosNetworkFromConfig(&cfg)
		orShutdown(errors.Wrap(err, "failed to get Helios network"))

		hyperionParams, err := (*heliosNetwork).HyperionParams(ctx)
		orShutdown(errors.Wrap(err, "failed to query hyperion params, is heliades running?"))

		for _, counterpartyChainParams := range hyperionParams.CounterpartyChainParams {
			if counterpartyChainParams.BridgeChainId == chainIdUint {
				fmt.Println("unsetting orchestrator addresses for chain", counterpartyChainParams.BridgeChainId)
				(*heliosNetwork).SendUnSetOrchestratorAddresses(ctx, counterpartyChainParams.HyperionId, common.HexToAddress(*cfg.ethKeyFrom).Hex())
			}
		}

		fmt.Println("done")
	}
}

func setOrchestratorCmd(cmd *cli.Cmd) {
	cfg := initConfig(cmd)

	chainId := cmd.String(cli.StringOpt{
		Name:   "chain-id",
		Desc:   "Specify which chain to query",
		EnvVar: "CHAIN_ID",
		Value:  "",
	})

	cmd.Action = func() {

		if *chainId == "" {
			orShutdown(errors.New("chain-id is required"))
		}

		ctx, cancelFn := context.WithCancel(context.Background())
		closer.Bind(cancelFn)

		chainIdUint, err := strconv.ParseUint(*chainId, 10, 64)
		orShutdown(errors.Wrap(err, "failed to parse chain id"))

		heliosNetwork, err := getHeliosNetworkFromConfig(&cfg)
		orShutdown(errors.Wrap(err, "failed to get Helios network"))

		hyperionParams, err := (*heliosNetwork).HyperionParams(ctx)
		orShutdown(errors.Wrap(err, "failed to query hyperion params, is heliades running?"))

		for _, counterpartyChainParams := range hyperionParams.CounterpartyChainParams {
			if counterpartyChainParams.BridgeChainId == chainIdUint {
				fmt.Println("setting orchestrator addresses for chain", counterpartyChainParams.BridgeChainId)
				(*heliosNetwork).SendSetOrchestratorAddresses(ctx, counterpartyChainParams.HyperionId, common.HexToAddress(*cfg.ethKeyFrom).Hex())
			}
		}

		fmt.Println("done")
	}
}

func listOperativeChainsCmd(cmd *cli.Cmd) {
	cfg := initConfig(cmd)

	cmd.Action = func() {

		heliosNetwork, err := getHeliosNetworkFromConfig(&cfg)
		orShutdown(errors.Wrap(err, "failed to get Helios network"))

		var chains []string

		networks, _ := helios.GetListOfNetworksWhereRegistered(*heliosNetwork, common.HexToAddress(*cfg.ethKeyFrom))

		fmt.Println("networkId: ", networks, "ssss", *cfg.ethKeyFrom)
		for _, networkId := range networks {
			fmt.Println("networkId: ", networkId)
			chains = append(chains, fmt.Sprintf("%d", networkId))
		}

		if len(chains) == 0 {
			fmt.Println("[]")
			return
		}

		data, _ := json.Marshal(chains)
		fmt.Println(string(data))
	}
}

func initializeBlockchainCmd(cmd *cli.Cmd) {
	cfg := initConfig(cmd)

	chainId := cmd.String(cli.StringOpt{
		Name:   "chain-id",
		Desc:   "Specify which chain to query",
		EnvVar: "CHAIN_ID",
		Value:  "",
	})

	cmd.Action = func() {
		ctx, cancelFn := context.WithCancel(context.Background())
		closer.Bind(cancelFn)

		chainIdUint, err := strconv.ParseUint(*chainId, 10, 64)
		orShutdown(errors.Wrap(err, "failed to parse chain id"))

		heliosNetwork, err := getHeliosNetworkFromConfig(&cfg)
		orShutdown(errors.Wrap(err, "failed to get Helios network"))

		counterpartyChainParams, err := (*heliosNetwork).GetCounterpartyChainParamsByChainId(ctx, chainIdUint)
		orShutdown(errors.Wrap(err, "failed to query hyperion params, is heliades running?"))

		if counterpartyChainParams == nil {
			orShutdown(errors.New("chain not found on helios please create a new proposal before using this command"))
		}

		ethKeyFromAddress, signerFn, _, err := initEthereumAccountsManager(
			uint64(counterpartyChainParams.BridgeChainId),
			cfg.ethKeystoreDir,
			cfg.ethKeyFrom,
			cfg.ethPassphrase,
			cfg.ethPrivKey,
		)

		if err != nil {
			orShutdown(errors.Wrap(err, fmt.Sprintf("failed to initialize ethereum accounts manager for chain %d", counterpartyChainParams.BridgeChainId)))
		}

		rpcs := counterpartyChainParams.Rpcs

		formattedRPCs := formatRPCs(*cfg.evmRPCs)

		if ok := formattedRPCs[fmt.Sprintf("%d", counterpartyChainParams.BridgeChainId)]; ok != nil {
			for _, rpc := range formattedRPCs[fmt.Sprintf("%d", counterpartyChainParams.BridgeChainId)] {
				log.Println("rpc", rpc)
				rpcs = append(rpcs, &hyperiontypes.Rpc{
					Url:            rpc,
					Reputation:     1,
					LastHeightUsed: 1,
				})
			}
		}

		// rpcChainListFeed := rpcchainlist.NewRpcChainListFeed()

		// rpcsFromChainList, err := rpcChainListFeed.QueryRpcs(counterpartyChainParams.BridgeChainId)
		// if err == nil {
		// 	for _, rpc := range rpcsFromChainList {
		// 		log.Println("rpc", rpc)
		// 		rpcs = append(rpcs, &hyperiontypes.Rpc{
		// 			Url:            rpc,
		// 			Reputation:     1,
		// 			LastHeightUsed: 1,
		// 		})
		// 	}
		// } else {
		// 	log.Println("Error fetching rpcs from chain list", err)
		// }

		fmt.Println("contract Address:", counterpartyChainParams.BridgeCounterpartyAddress)

		ethNetwork, err := ethereum.NewNetwork(common.HexToAddress(counterpartyChainParams.BridgeCounterpartyAddress), ethKeyFromAddress, signerFn, ethereum.NetworkConfig{
			EthNodeRPCs:           rpcs,
			GasPriceAdjustment:    *cfg.ethGasPriceAdjustment,
			MaxGasPrice:           *cfg.ethMaxGasPrice,
			PendingTxWaitDuration: *cfg.pendingTxWaitDuration,
		})
		if err != nil {
			orShutdown(errors.Wrap(err, fmt.Sprintf("failed to initialize ethereum network for chain %d", counterpartyChainParams.BridgeChainId)))
		}

		if !ethNetwork.TestRpcs(ctx) {
			orShutdown(errors.New("failed to test rpc"))
		}

		hyperionIdHash := common.HexToHash(strconv.FormatUint(counterpartyChainParams.HyperionId, 16))

		valset, err := (*heliosNetwork).CurrentValset(ctx, counterpartyChainParams.HyperionId)
		if err != nil {
			orShutdown(errors.Wrap(err, fmt.Sprintf("failed to query current valset for chain %d", counterpartyChainParams.BridgeChainId)))
		}
		if len(valset.Members) == 0 {
			orShutdown(errors.New("no members found in the valset, please subscribe before initializing the blockchain"))
		}

		myPower := big.NewInt(0)
		validators := make([]common.Address, len(valset.Members))
		powers := make([]*big.Int, len(valset.Members))
		for i, member := range valset.Members {
			validators[i] = common.HexToAddress(member.EthereumAddress)
			powers[i] = big.NewInt(int64(member.Power))
			if common.HexToAddress(member.EthereumAddress).Hex() == common.HexToAddress(*cfg.ethKeyFrom).Hex() {
				myPower = big.NewInt(int64(member.Power))
			}
		}
		powerThreshold := big.NewInt(int64(1431655765))
		if len(validators) == 1 {
			powerThreshold = big.NewInt(0).Div(myPower, big.NewInt(2))
		}

		fmt.Println("hyperionIdHash: ", hyperionIdHash)
		fmt.Println("powerThreshold: ", powerThreshold)
		fmt.Println("validators: ", validators)
		fmt.Println("powers: ", powers)

		if ethNetwork != nil {
			fmt.Println("ethNetwork: ", "OK")

		}

		balance, err := ethNetwork.GetNativeBalance(ctx)
		if err != nil {
			orShutdown(errors.Wrap(err, fmt.Sprintf("failed to get native balance for chain %d", counterpartyChainParams.BridgeChainId)))
		}
		fmt.Println("balance: ", balance)

		tx, err := ethNetwork.SendInitializeBlockchainTx(ctx, common.HexToAddress(*cfg.ethKeyFrom), hyperionIdHash, powerThreshold, validators, powers)
		if err != nil {
			orShutdown(errors.Wrap(err, fmt.Sprintf("failed to send initialize blockchain tx for chain %d", counterpartyChainParams.BridgeChainId)))
		}

		fmt.Println("tx Hash: ", tx.Hash().Hex())
	}
}
