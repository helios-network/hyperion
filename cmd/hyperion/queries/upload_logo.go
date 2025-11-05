package queries

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func UploadLogo(ctx context.Context, global *global.Global, logobase64 string) map[string]interface{} {
	network := *global.GetHeliosNetwork()
	msg, err := network.StoreLogoMsg(ctx, logobase64)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	// Generate a SHA-256 hash of the content
	hasher := sha256.New()
	hasher.Write([]byte(logobase64))
	logoHash := hex.EncodeToString(hasher.Sum(nil))

	err = network.SyncBroadcastMsgsSimulate(ctx, []sdk.Msg{msg})
	if err != nil {
		return map[string]interface{}{
			"error":     err.Error(),
			"logo_hash": logoHash,
		}
	}

	resp, err := global.SyncBroadcastMsgs(ctx, []sdk.Msg{msg})
	if err != nil {
		return map[string]interface{}{
			"error":     err.Error(),
			"logo_hash": logoHash,
		}
	}

	return map[string]interface{}{
		"tx_hash":   resp.TxHash,
		"logo_hash": logoHash,
	}
}
