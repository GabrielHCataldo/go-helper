package helper

import (
	"math"
)

// RoundUp rounding up ex: 1.5 -> 2
func RoundUp(v float64, precision int) float64 {
	return math.Ceil(v*(math.Pow10(precision))) / math.Pow10(precision)
}

// RoundDown rounding down ex: 1.5 -> 1
func RoundDown(v float64, precision int) float64 {
	return math.Floor(v*(math.Pow10(precision))) / math.Pow10(precision)
}

// Round rounding by value ex: 1.3 -> 1 or 1.6 -> 2
func Round(v float64, precision int) float64 {
	return math.Round(v*(math.Pow10(precision))) / math.Pow10(precision)
}

// MinFloat32 min float, if v is greater than min, min will prevail
func MinFloat32(v, min float32) float32 {
	if v < min {
		return min
	}
	return v
}

// MaxFloat32 max float, if v is less than max, max will prevail
func MaxFloat32(v, max float32) float32 {
	if v > max {
		return max
	}
	return v
}

// MinFloat64 min float, if v is greater than min, min will prevail
func MinFloat64(v, min float64) float64 {
	if v < min {
		return min
	}
	return v
}

// MaxFloat64 max float, if v is less than max, max will prevail
func MaxFloat64(v, max float64) float64 {
	if v > max {
		return max
	}
	return v
}
