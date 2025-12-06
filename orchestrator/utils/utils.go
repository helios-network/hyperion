package utils

import (
	"fmt"
	"math/big"

	sdkmath "cosmossdk.io/math"
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
	return fmt.Sprintf("%.18f", floatVal)
}

func FormatAmount(amount float64, decimals uint64) (sdkmath.Int, error) {
	dec, err := sdkmath.LegacyNewDecFromStr(fmt.Sprintf("%g", amount))
	if err != nil {
		return sdkmath.NewInt(0), err
	}

	tenPow := sdkmath.LegacyNewDec(1)
	for i := 0; i < int(decimals); i++ {
		tenPow = tenPow.MulInt64(10)
	}

	amountDec := dec.Mul(tenPow)
	amountMath := amountDec.TruncateInt()
	if amountMath.IsNegative() {
		return sdkmath.NewInt(0), fmt.Errorf("amount is negative")
	}
	return amountMath, nil
}

func ParseAmount(amount sdkmath.Int, decimals uint64) (float64, error) {
	dec, err := sdkmath.LegacyNewDecFromStr(amount.String())
	if err != nil {
		return 0, err
	}

	tenPow := sdkmath.LegacyNewDec(1)
	for i := 0; i < int(decimals); i++ {
		tenPow = tenPow.MulInt64(10)
	}

	amountDec := dec.Quo(tenPow)
	amountFloat, err := amountDec.Float64()
	if err != nil {
		return 0, err
	}
	return amountFloat, nil
}
