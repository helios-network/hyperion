package queries

import (
	"context"
	"fmt"
	"time"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
)

func ProposeHyperion(ctx context.Context, global *global.Global, title string, description string, bridgeChainId uint64, bridgeChainName string, averageCounterpartyBlockTime uint64) (map[string]interface{}, error) {
	network := *global.GetHeliosNetwork()
	hyperionParams, err := network.HyperionParams(ctx)
	if err != nil {
		return nil, err
	}

	counterpartyChainParams := hyperionParams.CounterpartyChainParams

	for _, counterpartyChainParam := range counterpartyChainParams {
		if counterpartyChainParam.BridgeChainId == bridgeChainId {
			return nil, fmt.Errorf("chainId %d already exists", bridgeChainId)
		}
	}

	hyperionContractInfo, err := storage.GetHyperionContractInfo(bridgeChainId)
	if err != nil {
		return nil, fmt.Errorf("please create hyperion contract for chainId %d before proposing", bridgeChainId)
	}

	hyperionAddress := hyperionContractInfo["hyperionAddress"].(string)
	startHeight := uint64(hyperionContractInfo["initializedAtBlockNumber"].(float64))

	proposalId, err := global.CreateNewBlockchainProposal(title, description, &hyperiontypes.CounterpartyChainParams{
		HyperionId:                   bridgeChainId,
		BridgeChainId:                bridgeChainId,
		BridgeChainName:              bridgeChainName,
		BridgeChainLogo:              "",
		AverageCounterpartyBlockTime: averageCounterpartyBlockTime,
		BridgeCounterpartyAddress:    hyperionAddress,
		BridgeContractStartHeight:    startHeight + 1, // +1 because the start height is the first block who hyperion will start listening
	})

	if err != nil {
		return nil, err
	}

	storage.UpdateHyperionContractInfo(bridgeChainId, hyperionAddress, map[string]interface{}{
		"proposed": true,
	})

	time.Sleep(4 * time.Second)

	global.VoteOnProposal(proposalId)

	return map[string]interface{}{
		"proposalId": proposalId,
	}, nil
}
