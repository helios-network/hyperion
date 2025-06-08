package helios

import (
	"context"

	ethcmn "github.com/ethereum/go-ethereum/common"
)

func RegisterValidator(
	ctx context.Context,
	hyperionID int,
	heliosNetwork Network,
	ethKeyFromAddress ethcmn.Address,
) error {
	err := heliosNetwork.SendSetOrchestratorAddresses(ctx, uint64(hyperionID), ethKeyFromAddress.String())
	if err != nil {
		return err
	}
	return nil
}

func UnRegisterValidator(
	ctx context.Context,
	hyperionID int,
	heliosNetwork Network,
	ethKeyFromAddress ethcmn.Address,
) error {
	err := heliosNetwork.SendUnSetOrchestratorAddresses(ctx, uint64(hyperionID), ethKeyFromAddress.String())
	if err != nil {
		return err
	}
	return nil
}
