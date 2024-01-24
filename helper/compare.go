package helper

import "strings"

// Equals compare values if are equals return true, otherwise return false
func Equals(a ...any) bool {
	a2 := a
	for _, v := range a {
		for _, v2 := range a2 {
			s, _ := ConvertToString(v)
			s2, _ := ConvertToString(v2)
			if s != s2 {
				return false
			}
		}
	}
	return true
}

// IsNotEqualTo compare values if aren't equals return true, otherwise return false
func IsNotEqualTo(a ...any) bool {
	return !Equals(a...)
}

// EqualsIgnoreCase compare values if are equals ignoring case return true, otherwise return false
func EqualsIgnoreCase(a ...any) bool {
	a2 := a
	for _, v := range a {
		for _, v2 := range a2 {
			s := strings.ToLower(SimpleConvertToString(v))
			s2 := strings.ToLower(SimpleConvertToString(v2))
			if s != s2 {
				return false
			}
		}
	}
	return true
}

// IsNotEqualToIgnoreCase compare values if aren't equals ignoring case return true, otherwise return false
func IsNotEqualToIgnoreCase(a ...any) bool {
	return !EqualsIgnoreCase(a...)
}

// IsGreaterThan compares whether A is greater than all values passed in B, If the value is not numeric, we will convert
// them to string and compare the size
func IsGreaterThan(a any, b ...any) bool {
	for _, bv := range b {
		fa, err := ConvertToFloat(a)
		if IsNotNil(err) {
			s, _ := ConvertToString(a)
			fa = float64(len(s))
		}
		fb, err := ConvertToFloat(bv)
		if IsNotNil(err) {
			s, _ := ConvertToString(bv)
			fb = float64(len(s))
		}
		if fb >= fa {
			return false
		}
	}
	return true
}

// IsGreaterThanOrEqual compares whether A is greater than or equal to all values passed in B, If the value is not numeric, we will convert
// them to string and compare the size
func IsGreaterThanOrEqual(a any, b ...any) bool {
	for _, bv := range b {
		fa, err := ConvertToFloat(a)
		if IsNotNil(err) {
			s, _ := ConvertToString(a)
			fa = float64(len(s))
		}
		fb, err := ConvertToFloat(bv)
		if IsNotNil(err) {
			s, _ := ConvertToString(bv)
			fb = float64(len(s))
		}
		if fb != fa && fb > fa {
			return false
		}
	}
	return true
}

// IsLessThan compares whether A is less than all values passed in B, If the value is not numeric, we will convert
// them to string and compare the size
func IsLessThan(a any, b ...any) bool {
	for _, bv := range b {
		fa, err := ConvertToFloat(a)
		if IsNotNil(err) {
			s, _ := ConvertToString(a)
			fa = float64(len(s))
		}
		fb, _ := ConvertToFloat(bv)
		if IsNotNil(err) {
			s, _ := ConvertToString(bv)
			fb = float64(len(s))
		}
		if fa >= fb {
			return false
		}
	}
	return true
}

// IsLessThanOrEqual compares whether A is less than or equal to all values passed in B, If the value is not numeric, we will convert
// them to string and compare the size
func IsLessThanOrEqual(a any, b ...any) bool {
	for _, bv := range b {
		fa, err := ConvertToFloat(a)
		if IsNotNil(err) {
			s, _ := ConvertToString(a)
			fa = float64(len(s))
		}
		fb, err := ConvertToFloat(bv)
		if IsNotNil(err) {
			s, _ := ConvertToString(bv)
			fb = float64(len(s))
		}
		if fa != fb && fa > fb {
			return false
		}
	}
	return true
}
