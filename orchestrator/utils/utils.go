package utils

import (
	"fmt"
	"math/big"
)

func FormatBigStringToFloat64(cost string, decimals int64) string {
	// convertir la string en big.Int
	costBigInt, ok := new(big.Int).SetString(cost, 10)
	if !ok {
		return cost
	}

	// convertir big.Int → float64 avec <decimals> décimales
	// big.Rat permet de garder la précision avant conversion
	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(decimals), nil) // 10^18
	rat := new(big.Rat).SetFrac(costBigInt, divisor)

	// conversion en float64 pour affichage
	floatVal, _ := rat.Float64()

	// format final avec 6 décimales visibles
	return fmt.Sprintf("%.6f", floatVal)
}
