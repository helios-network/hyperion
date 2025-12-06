package queries

import (
	"context"
	"fmt"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/utils"
)

func UpdateChainSettings(ctx context.Context, global *global.Global, chainId uint64, settings map[string]interface{}) error {
	baseSettings, err := GetChainSettings(ctx, global, chainId)
	if err != nil {
		fmt.Println("Error getting chain settings: ", err)
		return err
	}

	if settings["min_tx_fee_hls"] != nil && baseSettings["min_tx_fee_hls"] != nil && settings["min_tx_fee_hls"].(float64) != baseSettings["min_tx_fee_hls"].(float64) || settings["min_batch_fee_hls"] != nil && baseSettings["min_batch_fee_hls"] != nil && settings["min_batch_fee_hls"].(float64) != baseSettings["min_batch_fee_hls"].(float64) {
		err := UpdateFeeHyperion(ctx, global, settings["min_tx_fee_hls"].(float64), settings["min_batch_fee_hls"].(float64), chainId)
		if err != nil {
			fmt.Println("Error updating fee hyperion: ", err)
			return err
		}
	}
	err = storage.SetChainSettings(chainId, settings)
	if err != nil {
		fmt.Println("Error updating chain settings: ", err)
		return err
	}

	return nil
}

func GetChainSettings(ctx context.Context, global *global.Global, chainId uint64) (map[string]interface{}, error) {
	settings, err := storage.GetChainSettings(chainId)
	if err != nil {
		return nil, err
	}

	for key, value := range storage.DefaultChainSettingsMap {
		if _, ok := settings[key]; !ok {
			settings[key] = value
		}
	}

	network := *global.GetHeliosNetwork()
	registeredNetworks, err := network.GetValidatorHyperionData(ctx, global.GetAddress())
	if err != nil {
		return settings, nil
	}

	for _, orchestratorHyperionData := range registeredNetworks.OrchestratorHyperionData {
		if orchestratorHyperionData.HyperionId == chainId {
			minTxFeeHLS, err := utils.ParseAmount(orchestratorHyperionData.MinimumTxFee, 18)
			if err != nil {
				fmt.Println("Error parsing min tx fee hls: ", err)
				return nil, err
			}
			settings["min_tx_fee_hls"] = minTxFeeHLS
			minBatchFeeHLS, err := utils.ParseAmount(orchestratorHyperionData.MinimumBatchFee, 18)
			if err != nil {
				fmt.Println("Error parsing min batch fee hls: ", err)
				return nil, err
			}
			settings["min_batch_fee_hls"] = minBatchFeeHLS
			break
		}
	}

	return settings, nil
}
