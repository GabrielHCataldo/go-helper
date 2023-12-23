package helper

func MinInt(v, min int) int {
	if v < min {
		return min
	}
	return v
}

func MinInt32(v, min int32) int32 {
	if v < min {
		return min
	}
	return v
}

func MinInt64(v, min int64) int64 {
	if v < min {
		return min
	}
	return v
}

func MaxInt(v, max int) int {
	if v > max {
		return max
	}
	return v
}

func MaxInt32(v, max int32) int32 {
	if v > max {
		return max
	}
	return v
}

func MaxInt64(v, max int64) int64 {
	if v > max {
		return max
	}
	return v
}
