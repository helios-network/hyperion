package queries

import (
	"context"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/utils"
)

var chainIdToNativeCurrency = map[uint64]string{
	1:        "ETH",
	137:      "POL",
	56:       "BNB",
	10:       "ETH",
	42161:    "ETH",
	97:       "tBNB",
	11155111: "tETH",
	84532:    "ETH",
}

func GetListTransactions(ctx context.Context, global *global.Global, page int, size int) (map[string]interface{}, error) {
	fees, err := storage.GetFeesFile()
	if err != nil {
		return map[string]interface{}{
			"transactions": []map[string]interface{}{},
			"total":        0,
		}, nil
	}

	total := len(fees)

	// get page and size
	start := (page - 1) * size
	end := start + size
	if end > len(fees) {
		end = len(fees)
	}
	fees = fees[start:end]

	for _, fee := range fees {
		if fee["cost"] != nil && fee["tx_type"].(string) == "CLAIM" {
			fee["cost"] = utils.FormatBigStringToFloat64(fee["cost"].(string), 18) + " HLS"
		} else if fee["cost"] != nil && fee["tx_type"].(string) == "BATCH" {
			nativeCurrency := ""
			if fee["chain_id"] != nil && chainIdToNativeCurrency[uint64(fee["chain_id"].(float64))] != "" {
				nativeCurrency = chainIdToNativeCurrency[uint64(fee["chain_id"].(float64))]
			}
			fee["cost"] = utils.FormatBigStringToFloat64(fee["cost"].(string), 18) + " " + nativeCurrency
		}

		if fee["fees_taken"] != nil && fee["tx_type"].(string) == "BATCH" {
			fee["fees_taken"] = utils.FormatBigStringToFloat64(fee["fees_taken"].(string), 18) + " HLS"
		} else {
			fee["fees_taken"] = "0 HLS"
		}
	}

	return map[string]interface{}{
		"transactions": fees,
		"total":        total,
	}, nil
}
