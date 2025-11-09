package queries

import (
	"context"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
)

func DeleteChain(ctx context.Context, global *global.Global, chainId uint64) error {
	storage.RemoveHyperionContractInfo(chainId)
	return nil
}
