package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/helios"
	"github.com/ethereum/go-ethereum/common"
	cli "github.com/jawher/mow.cli"
	"github.com/pkg/errors"
	"github.com/xlab/closer"
)

// queryCmdSubset contains actions that query stuff from Hyperion module
// and the Ethereum contract
//
// $ hyperion q
func queryCmdSubset(cmd *cli.Cmd) {
	cmd.Command("unset-orchestrator", "Unset the orchestrator addresses for a chain", unsetOrchestratorCmd)
	cmd.Command("set-orchestrator", "Set the orchestrator addresses for a chain", setOrchestratorCmd)
	cmd.Command("list-operative-chains", "List the chains that are operative", listOperativeChainsCmd)
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
