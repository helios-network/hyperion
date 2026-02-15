package queries

import (
	"context"
	"fmt"
	"time"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
)

func ProposeUpdateAverageCounterpartyBlockTime(ctx context.Context, global *global.Global, title string, description string, chainId uint64, averageCounterpartyBlockTime uint64) (map[string]interface{}, error) {
	network := *global.GetHeliosNetwork()
	hyperionParams, err := network.HyperionParams(ctx)
	if err != nil {
		return nil, err
	}

	counterpartyChainParams := hyperionParams.CounterpartyChainParams

	alreadyExists := false
	for _, counterpartyChainParam := range counterpartyChainParams {
		if counterpartyChainParam.BridgeChainId == chainId {
			alreadyExists = true
			break
		}
	}

	if !alreadyExists {
		return nil, fmt.Errorf("chainId %d does not exist", chainId)
	}

	proposalId, err := global.ProposeUpdateAverageCounterpartyBlockTime(title, description, chainId, averageCounterpartyBlockTime)

	if err != nil {
		return nil, err
	}

	time.Sleep(4 * time.Second)

	global.VoteOnProposal(proposalId)

	return map[string]interface{}{
		"proposalId": proposalId,
	}, nil
}
