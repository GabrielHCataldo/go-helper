package helper

import (
	"reflect"
	"strings"
)

// IsNil check value is nil
func IsNil(v any) bool {
	elem := reflect.ValueOf(v)
	switch elem.Kind() {
	case reflect.Pointer, reflect.Chan, reflect.Map, reflect.Interface, reflect.Slice, reflect.Array:
		return elem.IsNil()
	default:
		return !elem.IsValid()
	}
}

// IsNotNil check value is not nil
func IsNotNil(v any) bool {
	return !IsNil(v)
}

// AllNil check all values are nil
func AllNil(v ...any) bool {
	for _, a := range v {
		if IsNotNil(a) {
			return false
		}
	}
	return true
}

// AllNotNil check all values are not nil
func AllNotNil(v ...any) bool {
	return !AllNil(v...)
}

// IsEmpty check value is empty
func IsEmpty(v any) bool {
	return isReflectZero(v)
}

// IsNotEmpty check value is not empty
func IsNotEmpty(v any) bool {
	return !IsEmpty(v)
}

// AllEmpty check all values are empty
func AllEmpty(v ...any) bool {
	for _, a := range v {
		if IsNotEmpty(a) {
			return false
		}
	}
	return true
}

// AllNotEmpty check all values are not empty
func AllNotEmpty(v ...any) bool {
	return !AllEmpty(v...)
}

// IsPointerNil check value pointer is nil
func IsPointerNil(v any) bool {
	if IsNotPointer(v) {
		panic("value type is not a pointer")
	}
	return isReflectNil(v)
}

// IsPointerNonNil check value pointer is not nil
func IsPointerNonNil(v any) bool {
	return !IsPointerNil(v)
}

// IsJsonEmpty check value struct, map or slice is empty
func IsJsonEmpty(v any) bool {
	if IsNotJson(v) {
		panic("value type is not struct, map or slice")
	}
	return isReflectZero(v)
}

// IsJsonNotEmpty check value struct, map or slice is not empty
func IsJsonNotEmpty(v any) bool {
	return !IsJsonEmpty(v)
}

// IsMapEmpty check value map is empty
func IsMapEmpty(v any) bool {
	if IsNotMap(v) {
		panic("value type is not a map")
	}
	return isReflectZero(v)
}

// IsMapNotEmpty check value map is not empty
func IsMapNotEmpty(v any) bool {
	return !IsMapEmpty(v)
}

// IsStructEmpty check value struct is empty
func IsStructEmpty(v any) bool {
	if IsNotStruct(v) {
		panic("value type is not a struct")
	}
	return isReflectZero(v)
}

// IsStructNotEmpty check value struct is not empty
func IsStructNotEmpty(v any) bool {
	return !IsStructEmpty(v)
}

// IsSliceEmpty check value slice is empty
func IsSliceEmpty(v any) bool {
	if IsNotSlice(v) {
		panic("value type is not a slice or array")
	}
	return isReflectZero(v)
}

// IsSliceNotEmpty check value slice is not empty
func IsSliceNotEmpty(v any) bool {
	return !IsSliceEmpty(v)
}

// IsStringEmpty check value string is empty (value blank return true)
func IsStringEmpty(v any) bool {
	if IsNotString(v) {
		panic("value type is not a string")
	}
	str := GetRealValue(v).(string)
	str = strings.TrimSpace(CleanAllRepeatSpaces(str))
	return len(str) == 0
}

// IsStringNotEmpty check value string is not empty
func IsStringNotEmpty(v any) bool {
	return !IsStringEmpty(v)
}

// IsIntEmpty check value int is empty
func IsIntEmpty(v any) bool {
	if IsNotInt(v) {
		panic("value type is not a int")
	}
	return isReflectZero(v)
}

// IsIntNotEmpty check value int is not empty
func IsIntNotEmpty(v any) bool {
	return !IsIntEmpty(v)
}

// IsFloatEmpty check value float is empty
func IsFloatEmpty(v any) bool {
	if IsNotFloat(v) {
		panic("value type is not a float")
	}
	return isReflectZero(v)
}

// IsFloatNotEmpty check value float is not empty
func IsFloatNotEmpty(v any) bool {
	return !IsFloatEmpty(v)
}

// IsBoolEmpty check value bool is empty
func IsBoolEmpty(v any) bool {
	if IsNotBool(v) {
		panic("value type is not a bool")
	}
	return isReflectZero(v)
}

// IsBoolNotEmpty check value bool is not empty
func IsBoolNotEmpty(v any) bool {
	return !IsBoolEmpty(v)
}

func isReflectZero(v any) bool {
	elem := reflect.ValueOf(v)
	if IsPointer(v) || IsInterface(v) {
		elem = elem.Elem()
	}
	if IsString(v) {
		return len(strings.TrimSpace(CleanAllRepeatSpaces(elem.String()))) == 0
	} else if IsSlice(v) {
		return elem.Len() == 0
	} else if IsMap(v) {
		return len(elem.MapKeys()) == 0
	}
	return !elem.IsValid() || elem.IsZero() || IsNil(v)
}

func isReflectNil(v any) bool {
	elem := reflect.ValueOf(v)
	return !elem.IsValid() || elem.IsNil()
}
