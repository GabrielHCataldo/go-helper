package helper

// Equal compare values if are equals return true, otherwise return false
func Equal(a ...any) bool {
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

// NotEqual compare values if aren't equals return true, otherwise return false
func NotEqual(a ...any) bool {
	return !Equal(a...)
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
