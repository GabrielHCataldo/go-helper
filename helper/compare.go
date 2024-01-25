package helper

import "strings"

// Equals compare values if are equals return true, otherwise return false
func Equals(a, b any, c ...any) bool {
	c = append(c, a)
	c = append(c, b)
	c2 := c
	for _, v := range c {
		for _, v2 := range c2 {
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
func IsNotEqualTo(a, b any, c ...any) bool {
	return !Equals(a, b, c...)
}

// EqualsIgnoreCase compare values if are equals ignoring case return true, otherwise return false
func EqualsIgnoreCase(a, b any, c ...any) bool {
	c = append(c, a)
	c = append(c, b)
	c2 := c
	for _, v := range c {
		for _, v2 := range c2 {
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
func IsNotEqualToIgnoreCase(a, b any, c ...any) bool {
	return !EqualsIgnoreCase(a, b, c...)
}

// IsGreaterThan compares whether A is greater than all values passed in others params, If the value is not numeric,
// we will convert them to string and compare the size
func IsGreaterThan(a, b any, c ...any) bool {
	c = append(c, b)
	for _, cv := range c {
		fa, err := ConvertToFloat(a)
		if IsNotNil(err) {
			s, _ := ConvertToString(a)
			fa = float64(len(s))
		}
		fb, err := ConvertToFloat(cv)
		if IsNotNil(err) {
			s, _ := ConvertToString(cv)
			fb = float64(len(s))
		}
		if fb >= fa {
			return false
		}
	}
	return true
}

// IsGreaterThanOrEqual compares whether A is greater than or equal to all values passed in others params, If the value
// is not numeric, we will convert them to string and compare the size
func IsGreaterThanOrEqual(a, b any, c ...any) bool {
	c = append(c, b)
	for _, cv := range c {
		fa, err := ConvertToFloat(a)
		if IsNotNil(err) {
			s, _ := ConvertToString(a)
			fa = float64(len(s))
		}
		fb, err := ConvertToFloat(cv)
		if IsNotNil(err) {
			s, _ := ConvertToString(cv)
			fb = float64(len(s))
		}
		if fb != fa && fb > fa {
			return false
		}
	}
	return true
}

// IsLessThan compares whether A is less than all values passed in others params, If the value is not numeric, we will
// convert them to string and compare the size
func IsLessThan(a, b any, c ...any) bool {
	c = append(c, b)
	for _, cv := range c {
		fa, err := ConvertToFloat(a)
		if IsNotNil(err) {
			s, _ := ConvertToString(a)
			fa = float64(len(s))
		}
		fb, _ := ConvertToFloat(cv)
		if IsNotNil(err) {
			s, _ := ConvertToString(cv)
			fb = float64(len(s))
		}
		if fa >= fb {
			return false
		}
	}
	return true
}

// IsLessThanOrEqual compares whether A is less than or equal to all values passed in others params, If the value is
// not numeric, we will convert them to string and compare the size
func IsLessThanOrEqual(a, b any, c ...any) bool {
	c = append(c, b)
	for _, cv := range c {
		fa, err := ConvertToFloat(a)
		if IsNotNil(err) {
			s, _ := ConvertToString(a)
			fa = float64(len(s))
		}
		fb, err := ConvertToFloat(cv)
		if IsNotNil(err) {
			s, _ := ConvertToString(cv)
			fb = float64(len(s))
		}
		if fa != fb && fa > fb {
			return false
		}
	}
	return true
}
