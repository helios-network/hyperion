package queries

import (
	"context"
	"errors"
	"fmt"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
)

func GetValidator(ctx context.Context, global *global.Global) (map[string]interface{}, error) {
	network := global.GetHeliosNetwork()
	if network == nil {
		return nil, errors.New("network not initialized")
	}

	validator, err := (*network).GetValidator(ctx, cosmostypes.ValAddress(global.GetCosmosAddress().Bytes()).String())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to get validator: %s", err.Error()))
	}

	return map[string]interface{}{
		"validator": validator,
	}, nil
}
