package helper

import (
	"reflect"
	"strings"
)

// IsNil check value is nil
func IsNil(a any) bool {
	elem := reflect.ValueOf(a)
	switch elem.Kind() {
	case reflect.Pointer, reflect.Chan, reflect.Map, reflect.Interface, reflect.Slice:
		return elem.IsNil() || !elem.IsValid()
	default:
		return !elem.IsValid()
	}
}

// IsNotNil check value is not nil
func IsNotNil(a any) bool {
	return !IsNil(a)
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
func IsEmpty(a any) bool {
	return isReflectZero(a)
}

// IsNotEmpty check value is not empty
func IsNotEmpty(a any) bool {
	return !IsEmpty(a)
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
func IsPointerNil(a any) bool {
	if IsNotPointer(a) {
		panic("value type is not a pointer")
	}
	return isReflectNil(a)
}

// IsPointerNonNil check value pointer is not nil
func IsPointerNonNil(a any) bool {
	return !IsPointerNil(a)
}

// IsJsonEmpty check value struct, map or slice is empty
func IsJsonEmpty(a any) bool {
	if IsNotJson(a) {
		panic("value type is not struct, map or slice")
	}
	return isReflectZero(a)
}

// IsJsonNotEmpty check value struct, map or slice is not empty
func IsJsonNotEmpty(a any) bool {
	return !IsJsonEmpty(a)
}

// IsMapEmpty check value map is empty
func IsMapEmpty(a any) bool {
	if IsNotMap(a) {
		panic("value type is not a map")
	}
	return isReflectZero(a)
}

// IsMapNotEmpty check value map is not empty
func IsMapNotEmpty(a any) bool {
	return !IsMapEmpty(a)
}

// IsStructEmpty check value struct is empty
func IsStructEmpty(a any) bool {
	if IsNotStruct(a) {
		panic("value type is not a struct")
	}
	return isReflectZero(a)
}

// IsStructNotEmpty check value struct is not empty
func IsStructNotEmpty(a any) bool {
	return !IsStructEmpty(a)
}

// IsSliceEmpty check value slice is empty
func IsSliceEmpty(a any) bool {
	if IsNotSlice(a) {
		panic("value type is not a slice or array")
	}
	return isReflectZero(a)
}

// IsSliceNotEmpty check value slice is not empty
func IsSliceNotEmpty(a any) bool {
	return !IsSliceEmpty(a)
}

// IsStringEmpty check value string is empty (value blank return true)
func IsStringEmpty(a any) bool {
	if IsNotString(a) {
		panic("value type is not a string")
	}
	str := getRealValue(a).(string)
	str = strings.TrimSpace(CleanAllRepeatSpaces(str))
	return len(str) == 0
}

// IsStringNotEmpty check value string is not empty
func IsStringNotEmpty(a any) bool {
	return !IsStringEmpty(a)
}

// IsIntEmpty check value int is empty
func IsIntEmpty(a any) bool {
	if IsNotInt(a) {
		panic("value type is not a int")
	}
	return isReflectZero(a)
}

// IsIntNotEmpty check value int is not empty
func IsIntNotEmpty(a any) bool {
	return !IsIntEmpty(a)
}

// IsFloatEmpty check value float is empty
func IsFloatEmpty(a any) bool {
	if IsNotFloat(a) {
		panic("value type is not a float")
	}
	return isReflectZero(a)
}

// IsFloatNotEmpty check value float is not empty
func IsFloatNotEmpty(a any) bool {
	return !IsFloatEmpty(a)
}

// IsBoolEmpty check value bool is empty
func IsBoolEmpty(a any) bool {
	if IsNotBool(a) {
		panic("value type is not a bool")
	}
	return isReflectZero(a)
}

// IsBoolNotEmpty check value bool is not empty
func IsBoolNotEmpty(a any) bool {
	return !IsBoolEmpty(a)
}

func isReflectZero(a any) bool {
	elem := reflect.ValueOf(a)
	if IsPointer(a) || IsInterface(a) {
		elem = elem.Elem()
	}
	if IsString(a) {
		return len(strings.TrimSpace(CleanAllRepeatSpaces(elem.String()))) == 0
	} else if IsSlice(a) {
		return elem.Len() == 0
	} else if IsMap(a) {
		return len(elem.MapKeys()) == 0
	}
	return !elem.IsValid() || elem.IsZero() || IsNil(a)
}

func isReflectNil(a any) bool {
	elem := reflect.ValueOf(a)
	return !elem.IsValid() || elem.IsNil()
}
