package queries

import (
	"context"
	"fmt"
	"slices"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/helios"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
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
			"address":                          counterpartyChainParam.BridgeCounterpartyAddress,
			"chainId":                          counterpartyChainParam.BridgeChainId,
			"name":                             counterpartyChainParam.BridgeChainName,
			"logo":                             counterpartyChainParam.BridgeChainLogo,
			"bridgeChainType":                  counterpartyChainParam.BridgeChainType,
			"contractSourceHash":               counterpartyChainParam.ContractSourceHash,
			"signedValsetsWindow":              counterpartyChainParam.SignedValsetsWindow,
			"signedBatchesWindow":              counterpartyChainParam.SignedBatchesWindow,
			"signedClaimsWindow":               counterpartyChainParam.SignedClaimsWindow,
			"targetBatchTimeout":               counterpartyChainParam.TargetBatchTimeout,
			"targetOutgoingTxTimeout":          counterpartyChainParam.TargetOutgoingTxTimeout,
			"averageBlockTime":                 counterpartyChainParam.AverageBlockTime,
			"averageCounterpartyBlockTime":     counterpartyChainParam.AverageCounterpartyBlockTime,
			"slashFractionValset":              counterpartyChainParam.SlashFractionValset.String(),
			"slashFractionBatch":               counterpartyChainParam.SlashFractionBatch.String(),
			"slashFractionClaim":               counterpartyChainParam.SlashFractionClaim.String(),
			"slashFractionConflictingClaim":    counterpartyChainParam.SlashFractionConflictingClaim.String(),
			"slashFractionBadEthSignature":     counterpartyChainParam.SlashFractionBadEthSignature.String(),
			"unbondSlashingValsetsWindow":      counterpartyChainParam.UnbondSlashingValsetsWindow,
			"bridgeContractStartHeight":        counterpartyChainParam.BridgeContractStartHeight,
			"valsetReward":                     counterpartyChainParam.ValsetReward.String(),
			"initializer":                      counterpartyChainParam.Initializer,
			"offsetValsetNonce":                counterpartyChainParam.OffsetValsetNonce,
			"minCallExternalDataGas":           counterpartyChainParam.MinCallExternalDataGas,
			"registered":                       registered,
			"running":                          running,
			"paused":                           counterpartyChainParam.Paused,
			"enabled":                          true,
			"proposed":                         true,
		}
	}

	hyperionsDeployedAddresses, err := storage.GetMyHyperionsDeployedAddresses()
	if err != nil {
		return nil, err
	}
	for _, hyperionDeployedAddress := range hyperionsDeployedAddresses {
		key := fmt.Sprintf("%d", uint64(hyperionDeployedAddress["chainId"].(float64)))
		if hyperions[key] == nil {
			proposed := false
			if hyperionDeployedAddress["proposed"] != nil {
				proposed = hyperionDeployedAddress["proposed"].(bool)
			}
			contractType := "new"
			if hyperionDeployedAddress["type"] != nil && hyperionDeployedAddress["type"].(string) == "update" {
				contractType = "update"
			}

			hyperions[fmt.Sprintf("%d", uint64(hyperionDeployedAddress["chainId"].(float64)))] = map[string]interface{}{
				"address":    hyperionDeployedAddress["hyperionAddress"],
				"chainId":    uint64(hyperionDeployedAddress["chainId"].(float64)),
				"registered": false,
				"running":    false,
				"paused":     false,
				"proposed":   proposed,
				"enabled":    false,
				"type":       contractType,
			}
		} else {
			if hyperionDeployedAddress["hyperionAddress"] != hyperions[key].(map[string]interface{})["address"] {
				proposed := false
				if hyperionDeployedAddress["proposed"] != nil {
					proposed = hyperionDeployedAddress["proposed"].(bool)
				}
				contractType := "new"
				if hyperionDeployedAddress["type"] != nil && hyperionDeployedAddress["type"].(string) == "update" {
					contractType = "update"
				}
				hyperions[key].(map[string]interface{})["type"] = contractType
				hyperions[key].(map[string]interface{})["proposed"] = proposed
				hyperions[key].(map[string]interface{})["oldAddress"] = hyperions[key].(map[string]interface{})["address"]
				hyperions[key].(map[string]interface{})["address"] = hyperionDeployedAddress["hyperionAddress"]
			}
		}
	}

	return hyperions, nil
}
