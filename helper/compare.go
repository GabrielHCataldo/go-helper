package helper

import "strings"

// IsTrue compare any values if all true return true, otherwise return false
func IsTrue(a ...any) bool {
	for _, av := range a {
		b, _ := ConvertToBool(av)
		if !b {
			return false
		}
	}
	return true
}

// IsNotTrue compare any values if all false return true, otherwise return false
func IsNotTrue(a ...any) bool {
	for _, av := range a {
		b, _ := ConvertToBool(av)
		if b {
			return false
		}
	}
	return true
}

// IsEqual compare values if are equals return true, otherwise return false
func IsEqual(a ...any) bool {
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

// IsNotEqual compare values if aren't equals return true, otherwise return false
func IsNotEqual(a ...any) bool {
	return !IsEqual(a...)
}

// IsEqualIgnoreCase compare values if are equals ignoring case return true, otherwise return false
func IsEqualIgnoreCase(a ...any) bool {
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

// IsNotEqualIgnoreCase compare values if aren't equals ignoring case return true, otherwise return false
func IsNotEqualIgnoreCase(a ...any) bool {
	return !IsEqualIgnoreCase(a...)
}

// IsGreater compares whether A is greater than all values passed in B, If the value is not numeric, we will convert
// them to string and compare the size
func IsGreater(a any, b ...any) bool {
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

// IsGreaterEqual compares whether A is greater than or equal to all values passed in B, If the value is not numeric, we will convert
// them to string and compare the size
func IsGreaterEqual(a any, b ...any) bool {
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

// IsSmaller compares whether A is smaller than all values passed in B, If the value is not numeric, we will convert
// them to string and compare the size
func IsSmaller(a any, b ...any) bool {
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

// IsSmallerEqual compares whether A is smaller than or equal to all values passed in B, If the value is not numeric, we will convert
// them to string and compare the size
func IsSmallerEqual(a any, b ...any) bool {
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
