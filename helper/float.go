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
