package helper

import (
	"reflect"
	"strings"
)

// IsNil check all values are nil
func IsNil(a ...any) bool {
	for _, v := range a {
		if !isReflectNil(v) {
			return false
		}
	}
	return true
}

// IsNotNil check all values are not nil
func IsNotNil(a ...any) bool {
	return !IsNil(a...)
}

// IsEmpty check all value are empty
func IsEmpty(a ...any) bool {
	for _, v := range a {
		if !isReflectZero(v) {
			return false
		}
	}
	return true
}

// IsNotEmpty check all values are not empty
func IsNotEmpty(a ...any) bool {
	return !IsEmpty(a...)
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
	} else if IsObjectId(a) {
		return SimpleConvertToObjectId(a).IsZero()
	} else if IsPrimitiveDateTime(a) {
		return SimpleConvertToPrimitiveDateTime(a).Time().IsZero()
	}
	return !elem.IsValid() || elem.IsZero() || IsNil(a)
}

func isReflectNil(a any) bool {
	elem := reflect.ValueOf(a)
	switch elem.Kind() {
	case reflect.Pointer, reflect.Chan, reflect.Map, reflect.Interface, reflect.Slice:
		return elem.IsNil() || !elem.IsValid()
	default:
		return !elem.IsValid()
	}
}
