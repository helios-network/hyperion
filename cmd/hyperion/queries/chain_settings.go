package queries

import (
	"context"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
)

var defaultSettings = map[string]interface{}{
	"min_batch_fee_usd":        0,
	"eth_gas_price_adjustment": 1.3,
	"eth_max_gas_price":        "100gwei",
	"estimate_gas":             true,
	"eth_gas_price":            "10gwei",
	"valset_offset_dur":        "5m",
	"batch_offset_dur":         "2m",
	"static_rpc_anonymous":     true,
	"static_rpc_only":          false,
	"min_batch_fee_hls":        0.1,
	"min_tx_fee_hls":           0.1,
}

func UpdateChainSettings(ctx context.Context, global *global.Global, chainId uint64, settings map[string]interface{}) error {
	baseSettings, err := GetChainSettings(chainId)
	if err != nil {
		return err
	}

	if settings["min_tx_fee_hls"] != nil && baseSettings["min_tx_fee_hls"] != nil && settings["min_tx_fee_hls"].(float64) != baseSettings["min_tx_fee_hls"].(float64) || settings["min_batch_fee_hls"] != nil && baseSettings["min_batch_fee_hls"] != nil && settings["min_batch_fee_hls"].(float64) != baseSettings["min_batch_fee_hls"].(float64) {
		err := UpdateFeeHyperion(ctx, global, settings["min_tx_fee_hls"].(float64), settings["min_batch_fee_hls"].(float64), chainId)
		if err != nil {
			return err
		}
	}
	err = storage.SetChainSettings(chainId, settings)
	if err != nil {
		return err
	}
	storage.UpdateRpcsToStorge(chainId, []string{})

	return nil
}

func GetChainSettings(chainId uint64) (map[string]interface{}, error) {
	settings, err := storage.GetChainSettings(chainId, defaultSettings)
	if err != nil {
		return nil, err
	}

	for key, value := range defaultSettings {
		if _, ok := settings[key]; !ok {
			settings[key] = value
		}
	}

	return settings, nil
}
