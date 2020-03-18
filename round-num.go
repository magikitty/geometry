package geometry

import "math"

// RoundNumber returns number rounded to specified number of decimal places (between 0 - 10)
func RoundNumber(number float64, decimals int) float64 {
	precision := 1.0
	for i := 0; i < decimals; i++ {
		precision = precision * 10
	}
	roundedNum := math.Round(number*precision) / precision

	return roundedNum
}
