package utils

import "math"

func Sqrt(x float64) float64 {
	if x < 0 {
		return math.NaN()
	}
	halfx := x * 0.5
	x64 := math.Float64bits(x)
	x64 = 0x5FE6EB50C7B537A9 - (x64 >> 1)
	f := math.Float64frombits(x64)
	f *= float64(1.5) - (halfx * f * f)
	return f * x
}
