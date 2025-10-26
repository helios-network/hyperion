package queries

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
)

func formatBigStringToFloat64(cost string) string {
	// convertir la string en big.Int
	costBigInt, ok := new(big.Int).SetString(cost, 10)
	if !ok {
		return cost
	}

	// convertir big.Int → float64 avec 18 décimales
	// big.Rat permet de garder la précision avant conversion
	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil) // 10^18
	rat := new(big.Rat).SetFrac(costBigInt, divisor)

	// conversion en float64 pour affichage
	floatVal, _ := rat.Float64()

	// format final avec 6 décimales visibles
	return fmt.Sprintf("%.6f", floatVal)
}

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
		fmt.Println("fee", fee)
		if fee["cost"] != nil && fee["tx_type"].(string) == "CLAIM" {
			fee["cost"] = formatBigStringToFloat64(fee["cost"].(string)) + " HLS"
		} else if fee["cost"] != nil && fee["tx_type"].(string) == "BATCH" {
			nativeCurrency := ""
			if fee["chain_id"] != nil && chainIdToNativeCurrency[uint64(fee["chain_id"].(float64))] != "" {
				nativeCurrency = chainIdToNativeCurrency[uint64(fee["chain_id"].(float64))]
			}
			fee["cost"] = formatBigStringToFloat64(fee["cost"].(string)) + " " + nativeCurrency
		}
	}

	return map[string]interface{}{
		"transactions": fees,
		"total":        total,
	}, nil
}
