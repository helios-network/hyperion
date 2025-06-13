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
}

func UpdateChainSettings(ctx context.Context, global *global.Global, chainId uint64, settings map[string]interface{}) error {
	err := storage.SetChainSettings(chainId, settings)
	if err != nil {
		return err
	}

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
