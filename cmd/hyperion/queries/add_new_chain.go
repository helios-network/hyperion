package queries

import (
	"context"
	"fmt"
	"time"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
)

func AddNewChain(ctx context.Context, global *global.Global, chainId uint64) (map[string]interface{}, error) {
	network := *global.GetHeliosNetwork()
	hyperionParams, err := network.HyperionParams(ctx)
	if err != nil {
		return nil, err
	}

	counterpartyChainParams := hyperionParams.CounterpartyChainParams

	for _, counterpartyChainParam := range counterpartyChainParams {
		if counterpartyChainParam.BridgeChainId == chainId {
			return nil, fmt.Errorf("chainId %d already exists", chainId)
		}
	}

	hyperionContractInfo := map[string]interface{}{
		"hyperionAddress":          "0x0000000000000000000000000000000000000000",
		"chainId":                  chainId,
		"createdAt":                time.Now().Format(time.RFC3339),
		"atBlockNumber":            0,
		"initializedAtBlockNumber": 0,
	}

	err = storage.AddOneNewHyperionDeployedAddress(hyperionContractInfo)
	if err != nil {
		return nil, err
	}
	return hyperionContractInfo, nil
}
