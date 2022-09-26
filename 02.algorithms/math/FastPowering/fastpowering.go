package FastPowering

import (
	"math"
)

func FastPowering(base float64, power int) float64 {
	// Anything raised to the power of zero is 1
	if power == 0 {
		return 1
	}
	// Even powers are redefined via two smaller powers
	if power%2 == 0 {
		multiplier := FastPowering(base, power/2)
		return multiplier * multiplier
	}
	// Odd powers
	multiplier := FastPowering(base, int(math.Floor(float64(power/2))))
	return multiplier * multiplier * base
}
