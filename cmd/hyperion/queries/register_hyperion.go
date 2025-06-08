package queries

import (
	"context"
	"fmt"
	"slices"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/helios"
)

func RegisterHyperion(ctx context.Context, global *global.Global, chainId uint64) error {
	network := *global.GetHeliosNetwork()

	registeredNetworks, _ := helios.GetListOfNetworksWhereRegistered(*global.GetHeliosNetwork(), global.GetAddress())
	if slices.Contains(registeredNetworks, chainId) {
		return fmt.Errorf("hyperion already registered for chain %d", chainId)
	}

	err := network.SendSetOrchestratorAddresses(ctx, chainId, global.GetAddress().Hex())
	if err != nil {
		return err
	}
	return nil
}
