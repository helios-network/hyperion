package queries

import (
	"context"
	"fmt"
	"time"

	"cosmossdk.io/errors"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
)

func DeployHeliosTokenToChain(ctx context.Context, global *global.Global, chainId uint64, denom string, name string, symbol string, decimals uint8) (map[string]interface{}, error) {
	network := *global.GetHeliosNetwork()
	counterpartyChainParams, err := network.GetCounterpartyChainParamsByChainId(ctx, chainId)
	if err != nil {
		return nil, err
	}

	heliosNetwork := global.GetHeliosNetwork()
	if heliosNetwork == nil {
		return nil, fmt.Errorf("helios network not initialized")
	}

	targetNetwork, err := global.InitTargetNetwork(counterpartyChainParams)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to initialize target network for chain %d", counterpartyChainParams.BridgeChainId))
	}

	tx, blockNumber, err := (*targetNetwork).DeployERC20(ctx, global.GetAddress(), denom, name, symbol, decimals)
	if err != nil {
		return nil, errors.Wrap(err, "failed to deploy ERC20")
	}

	erc20DeploymentEvents, err := (*targetNetwork).GetHyperionERC20DeployedEvents(blockNumber, blockNumber)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get ERC20Deployed events")
	}

	erc20DeploymentEvent := erc20DeploymentEvents[0]

	erc20Address := erc20DeploymentEvent.TokenContract

	erc20ContractInfo := map[string]interface{}{
		"erc20Address":  erc20Address,
		"chainId":       counterpartyChainParams.BridgeChainId,
		"createdAt":     time.Now().Format(time.RFC3339),
		"atBlockNumber": blockNumber,
		"txHash":        tx.Hash().Hex(),
		"denom":         denom,
		"name":          name,
		"symbol":        symbol,
		"decimals":      decimals,
	}
	return erc20ContractInfo, nil
}
