package util

import "math"

const epsilon float64 = 1.0e-10

var (
	// MinNormal is the smallest positive normal value of type float64.
	MinNormal = math.Float64frombits(0x0010000000000000)
)

func AlmostEqual(a float64, b float64, epsilon float64) bool {
	absA := math.Abs(a)
	absB := math.Abs(b)
	diff := math.Abs(a - b)

	if a == b {
		return true
	} else if a == 0 || b == 0 || (absA+absB < MinNormal) {
		return diff < epsilon*MinNormal
	} else {
		return diff/math.Min(absA+absB, math.MaxFloat64) < epsilon
	}
}

func Greater(a float64, b float64) bool {
	return a > b && !AlmostEqual(a, b, epsilon)
}
