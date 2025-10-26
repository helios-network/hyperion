package queries

import (
	"context"
	"fmt"
	"time"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
)

func ProposeAddWhitelistedAddress(ctx context.Context, global *global.Global, title string, description string, chainId uint64, address string) (map[string]interface{}, error) {
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

	proposalId, err := global.ProposeMsgAddOneWhitelistedAddress(title, description, chainId, address)

	if err != nil {
		return nil, err
	}

	time.Sleep(4 * time.Second)

	global.VoteOnProposal(proposalId)

	return map[string]interface{}{
		"proposalId": proposalId,
	}, nil
}
