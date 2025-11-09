package queries

import (
	"context"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
)

func ResetHeliosClient(ctx context.Context, global *global.Global) error {
	global.ResetHeliosClient()
	return nil
}
