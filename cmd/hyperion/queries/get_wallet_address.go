package queries

import (
	"context"
	"errors"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
)

func GetWalletAddress(ctx context.Context, global *global.Global) (map[string]interface{}, error) {
	network := global.GetHeliosNetwork()
	if network == nil {
		return nil, errors.New("network not initialized")
	}
	address := global.GetAddress()

	return map[string]interface{}{
		"address": address.Hex(),
	}, nil
}
