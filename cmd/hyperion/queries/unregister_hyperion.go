package queries

import (
	"context"
	"fmt"
	"slices"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/helios"
)

func UnRegisterHyperion(ctx context.Context, global *global.Global, chainId uint64) error {
	network := *global.GetHeliosNetwork()

	// stop hyperion if it is running
	runner := global.GetRunner(chainId)
	if runner != nil {
		global.CancelRunner(chainId)
		fmt.Println("Hyperion stopped successfully for chain", chainId)
	}

	// check if hyperion is registered
	registeredNetworks, _ := helios.GetListOfNetworksWhereRegistered(*global.GetHeliosNetwork(), global.GetAddress())
	if !slices.Contains(registeredNetworks, chainId) {
		return fmt.Errorf("hyperion not registered for chain %d", chainId)
	}

	// unregister hyperion
	err := network.SendUnSetOrchestratorAddresses(ctx, chainId, global.GetAddress().Hex())
	if err != nil {
		return err
	}
	return nil
}
