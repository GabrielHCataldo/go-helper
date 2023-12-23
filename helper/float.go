package helper

import (
	"math"
)

func RoundUp(v float64, precision int) float64 {
	return math.Ceil(v*(math.Pow10(precision))) / math.Pow10(precision)
}

func RoundDown(v float64, precision int) float64 {
	return math.Floor(v*(math.Pow10(precision))) / math.Pow10(precision)
}

func Round(v float64, precision int) float64 {
	return math.Round(v*(math.Pow10(precision))) / math.Pow10(precision)
}
