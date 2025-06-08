package queries

import (
	"context"
	"fmt"
	"slices"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/helios"
)

func GetListHyperions(ctx context.Context, global *global.Global) (map[string]interface{}, error) {
	network := *global.GetHeliosNetwork()
	hyperionParams, err := network.HyperionParams(ctx)
	if err != nil {
		return nil, err
	}

	counterpartyChainParams := hyperionParams.CounterpartyChainParams
	registeredNetworks, _ := helios.GetListOfNetworksWhereRegistered(*global.GetHeliosNetwork(), global.GetAddress())

	hyperions := map[string]interface{}{}
	for _, counterpartyChainParam := range counterpartyChainParams {
		registered := false
		if slices.Contains(registeredNetworks, counterpartyChainParam.BridgeChainId) {
			registered = true
		}
		running := false
		if global.GetRunner(counterpartyChainParam.BridgeChainId) != nil {
			running = true
		}
		hyperions[fmt.Sprintf("%d", counterpartyChainParam.BridgeChainId)] = map[string]interface{}{
			"address":    counterpartyChainParam.BridgeCounterpartyAddress,
			"chainId":    counterpartyChainParam.BridgeChainId,
			"registered": registered,
			"running":    running,
			"paused":     counterpartyChainParam.Paused,
		}
	}
	return hyperions, nil
}
