package helper

import (
	"reflect"
	"strings"
)

// Equals compare values if are equals return true, otherwise return false
func Equals(a, b any, c ...any) bool {
	return equals(false, a, b, c...)
}

// IsNotEqualTo compare values if aren't equals return true, otherwise return false
func IsNotEqualTo(a, b any, c ...any) bool {
	return !Equals(a, b, c...)
}

// EqualsIgnoreCase compare values if are equals ignoring case return true, otherwise return false
func EqualsIgnoreCase(a, b any, c ...any) bool {
	return equals(true, a, b, c...)
}

// IsNotEqualToIgnoreCase compare values if aren't equals ignoring case return true, otherwise return false
func IsNotEqualToIgnoreCase(a, b any, c ...any) bool {
	return !EqualsIgnoreCase(a, b, c...)
}

// EqualsLen compares whether the size of the value of parameter A is equal to the suggested len parameter.
func EqualsLen(a any, len int) bool {
	return Len(a) == len
}

// IsNotEqualsLen compares whether the size of the value of parameter A is not equal to the suggested len parameter.
func IsNotEqualsLen(a any, len int) bool {
	return !EqualsLen(a, len)
}

// Contains if values passed in parameters B and C contain the value of parameter A, it returns true, otherwise
// it returns false
func Contains(a, b any, c ...any) bool {
	if IsSlice(a) {
		return containsSlice(a, b, c...)
	}
	return contains(false, a, b, c...)
}

// NotContains if values passed in parameters B and C do not contain the value of parameter A, it returns true, otherwise
// it returns false
func NotContains(a, b any, c ...any) bool {
	return !Contains(a, b, c...)
}

// ContainsIgnoreCase if values passed in parameters B and C contain the value of parameter A, it returns true, otherwise
// it returns false
func ContainsIgnoreCase(a, b any, c ...any) bool {
	return contains(true, a, b, c...)
}

// NotContainsIgnoreCase if values passed in parameters B and C do not contain the value of parameter A, it returns true,
// otherwise it returns false
func NotContainsIgnoreCase(a, b any, c ...any) bool {
	return !ContainsIgnoreCase(a, b, c...)
}

// IsGreaterThan compares whether A is greater than all values passed in other parameters, if the value is not numeric,
// let's use the Len function and compare the size
func IsGreaterThan(a, b any, c ...any) bool {
	c = append(c, b)
	for _, cv := range c {
		fa, err := ConvertToFloat(a)
		if IsNotNil(err) {
			fa = float64(Len(a))
		}
		fb, err := ConvertToFloat(cv)
		if IsNotNil(err) {
			fb = float64(Len(cv))
		}
		if fb >= fa {
			return false
		}
	}
	return true
}

// IsGreaterThanOrEqual compares whether A is greater than or equal to all values passed in others params, If the value
// is not numeric, let's use the Len function and compare the size
func IsGreaterThanOrEqual(a, b any, c ...any) bool {
	c = append(c, b)
	for _, cv := range c {
		fa, err := ConvertToFloat(a)
		if IsNotNil(err) {
			fa = float64(Len(a))
		}
		fb, err := ConvertToFloat(cv)
		if IsNotNil(err) {
			fb = float64(Len(cv))
		}
		if fb != fa && fb > fa {
			return false
		}
	}
	return true
}

// IsLessThan compares whether A is less than all values passed in others params, If the value is not numeric,
// let's use the Len function and compare the size
func IsLessThan(a, b any, c ...any) bool {
	c = append(c, b)
	for _, cv := range c {
		fa, err := ConvertToFloat(a)
		if IsNotNil(err) {
			fa = float64(Len(a))
		}
		fb, err := ConvertToFloat(cv)
		if IsNotNil(err) {
			fb = float64(Len(cv))
		}
		if fa >= fb {
			return false
		}
	}
	return true
}

// IsLessThanOrEqual compares whether A is less than or equal to all values passed in others params, If the value is
// not numeric, let's use the Len function and compare the size
func IsLessThanOrEqual(a, b any, c ...any) bool {
	c = append(c, b)
	for _, cv := range c {
		fa, err := ConvertToFloat(a)
		if IsNotNil(err) {
			fa = float64(Len(a))
		}
		fb, err := ConvertToFloat(cv)
		if IsNotNil(err) {
			fb = float64(Len(cv))
		}
		if fa != fb && fa > fb {
			return false
		}
	}
	return true
}

func equals(ignoreCase bool, a, b any, c ...any) bool {
	c = append(c, a)
	c = append(c, b)
	c2 := c
	for _, v := range c {
		for _, v2 := range c2 {
			s := SimpleConvertToString(v)
			s2 := SimpleConvertToString(v2)
			if ignoreCase {
				s = strings.ToLower(s)
				s2 = strings.ToLower(s2)
			}
			if s != s2 {
				return false
			}
		}
	}
	return true
}

func contains(ignoreCase bool, a, b any, c ...any) bool {
	c = append(c, b)
	s := SimpleConvertToString(a)
	if ignoreCase {
		s = strings.ToLower(s)
	}
	for _, v := range c {
		s2 := SimpleConvertToString(v)
		if ignoreCase {
			s2 = strings.ToLower(s2)
		}
		if !strings.Contains(s, s2) {
			return false
		}
	}
	return true
}

func containsSlice(a, b any, c ...any) bool {
	c = append(c, b)
	a = getRealValue(a)
	ra := reflect.ValueOf(a)
	for _, v := range c {
		v = getRealValue(ra)
		rv := reflect.ValueOf(v)
		if ra.Equal(rv) {
			return false
		}
	}
	return true
}

func getRealValue(a any) any {
	ra := reflect.ValueOf(a)
	if IsPointer(a) {
		ra = ra.Elem()
	}
	return ra.Interface()
}
