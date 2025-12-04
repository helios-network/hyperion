package utils

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
)

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
