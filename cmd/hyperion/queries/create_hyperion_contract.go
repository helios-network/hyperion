package queries

import (
	"context"
	"fmt"
	"time"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
)

func CreateHyperionContract(ctx context.Context, global *global.Global, chainId uint64) (map[string]interface{}, error) {
	network := *global.GetHeliosNetwork()
	hyperionParams, err := network.HyperionParams(ctx)
	if err != nil {
		return nil, err
	}

	counterpartyChainParams := hyperionParams.CounterpartyChainParams
	contractType := "new"
	for _, counterpartyChainParam := range counterpartyChainParams {
		if counterpartyChainParam.BridgeChainId == chainId {
			contractType = "update"
			break
		}
	}

	hyperionAddress, atBlockNumber, success := global.DeployNewHyperionContract(chainId)
	if !success {
		return nil, fmt.Errorf("failed to deploy hyperion contract")
	}

	hyperionContractInfo := map[string]interface{}{
		"hyperionAddress":          hyperionAddress.Hex(),
		"chainId":                  chainId,
		"createdAt":                time.Now().Format(time.RFC3339),
		"atBlockNumber":            atBlockNumber,
		"initializedAtBlockNumber": atBlockNumber + 1000,
		"type":                     contractType,
		"proposed":                 false,
	}

	fmt.Println("hyperionContractInfo: ", hyperionContractInfo)

	err = storage.UpdateHyperionContractInfo(chainId, hyperionContractInfo)
	if err != nil {
		return nil, err
	}

	blockNumber, err := global.InitializeHyperionContractWithDefaultValset(chainId)
	if err != nil {
		storage.RemoveHyperionContractInfo(chainId)
		return nil, err
	}
	hyperionContractInfo["initializedAtBlockNumber"] = blockNumber
	err = storage.UpdateHyperionContractInfo(chainId, hyperionContractInfo)
	if err != nil {
		return nil, err
	}
	return hyperionContractInfo, nil
}
