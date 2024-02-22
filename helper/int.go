package helper

import (
	"math/rand"
)

// MinInt min int, if v is greater than min, min will prevail
func MinInt(v, min int) int {
	if v < min {
		return min
	}
	return v
}

// MinInt32 min int, if v is greater than min, min will prevail
func MinInt32(v, min int32) int32 {
	if v < min {
		return min
	}
	return v
}

// MinInt64 min int, if v is greater than min, min will prevail
func MinInt64(v, min int64) int64 {
	if v < min {
		return min
	}
	return v
}

// MaxInt max int, if v is less than max, max will prevail
func MaxInt(v, max int) int {
	if v > max {
		return max
	}
	return v
}

// MaxInt32 max int, if v is less than max, max will prevail
func MaxInt32(v, max int32) int32 {
	if v > max {
		return max
	}
	return v
}

// MaxInt64 max int, if v is less than max, max will prevail
func MaxInt64(v, max int64) int64 {
	if v > max {
		return max
	}
	return v
}

// RandomNumber random number generator with range min and max
func RandomNumber(min, max int) int {
	return rand.Intn(max-min) + min
}
