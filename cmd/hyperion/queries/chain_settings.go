package queries

import (
	"context"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
)

func UpdateChainSettings(ctx context.Context, global *global.Global, chainId uint64, settings map[string]interface{}) error {
	err := storage.SetChainSettings(chainId, settings)
	if err != nil {
		return err
	}

	return nil
}

func GetChainSettings(ctx context.Context, global *global.Global, chainId uint64) (map[string]interface{}, error) {
	return storage.GetChainSettings(chainId)
}
