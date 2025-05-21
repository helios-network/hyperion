package helios

import (
	"context"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum"

	sdkmath "cosmossdk.io/math"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
	log "github.com/xlab/suplog"
)

func TestnetForceUpdateValset(
	ctx context.Context,
	hyperionID int,
	heliosNetwork Network,
	ethNetwork ethereum.Network,
) error {
	// check if the helios hyperion is not synchronized
	nonce, err := heliosNetwork.QueryGetLastObservedEventNonce(ctx, uint64(hyperionID))
	if err != nil {
		return err
	}
	lastEventNonce, err := ethNetwork.GetLastEventNonce(ctx)
	if err != nil {
		return err
	}
	if nonce == 0 && lastEventNonce.Uint64() > 1 { // not firstime

		height, err := ethNetwork.GetLastValsetUpdatedEventHeight(ctx)
		if err != nil {
			return err
		}

		lastEventBlockHeight, err := ethNetwork.GetLastEventHeight(ctx)
		if err != nil {
			return err
		}

		events, err := ethNetwork.GetValsetUpdatedEventsAtSpecificBlock(height.Uint64())
		if err != nil {
			return err
		}

		if len(events) == 0 {
			log.Fatalln("helios hyperion is not synchronized, please wait for it to be synchronized")
		}

		event := events[0]

		valset := &hyperiontypes.Valset{
			Nonce:        event.NewValsetNonce.Uint64(),
			Members:      make([]*hyperiontypes.BridgeValidator, 0, len(event.Powers)),
			RewardAmount: sdkmath.NewIntFromBigInt(event.RewardAmount),
			RewardToken:  event.RewardToken.Hex(),
		}

		for idx, p := range event.Powers {
			valset.Members = append(valset.Members, &hyperiontypes.BridgeValidator{
				Power:           p.Uint64(),
				EthereumAddress: event.Validators[idx].Hex(),
			})
		}

		err = heliosNetwork.SendForceSetValsetAndLastObservedEventNonce(ctx, uint64(hyperionID), lastEventNonce.Uint64(), lastEventBlockHeight.Uint64(), valset)
		if err != nil {
			return err
		}

		log.Infoln("helios hyperion is now forcefully synchronized with ethereum hyperion")
	}
	return nil
}
