package helios

import (
	"context"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	ethcmn "github.com/ethereum/go-ethereum/common"
	log "github.com/xlab/suplog"
)

func TestnetAutoRegisterValidator(
	ctx context.Context,
	hyperionID int,
	heliosNetwork Network,
	isValidator bool,
	addr cosmostypes.AccAddress,
	ethKeyFromAddress ethcmn.Address,
) (bool, error) {
	if isValidator {
		log.Debugln("provided ETH address is registered with an orchestrator", addr.String())
	} else {
		err := heliosNetwork.SendSetOrchestratorAddresses(ctx, uint64(hyperionID), ethKeyFromAddress.String())
		if err != nil {
			return false, err
		}
		addr, isValidator = HasRegisteredOrchestrator(heliosNetwork, uint64(hyperionID), ethKeyFromAddress)
		if isValidator {
			log.Debugln("provided ETH address is registered with an orchestrator", addr.String())
		}
	}
	return true, nil
}
