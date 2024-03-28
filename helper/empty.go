package helper

import (
	"reflect"
	"strings"
)

// IsNil check all values are nil
func IsNil(a any, b ...any) bool {
	b = append(b, a)
	for _, v := range b {
		if !isReflectNil(v) {
			return false
		}
	}
	return true
}

// IsNotNil check all values are not nil
func IsNotNil(a any, b ...any) bool {
	return !IsNil(a, b...)
}

// IfNilReturns if A is nil return value B, otherwise return A value
func IfNilReturns[T any](a *T, b T) T {
	if IsNil(a) {
		return b
	}
	return *a
}

// ReturnNonNilValue returns the information value that is not nil
func ReturnNonNilValue[T any](a ...T) T {
	var result T
	for _, v := range a {
		if IsNotNil(v) {
			result = v
			break
		}
	}
	return result
}

// IsEmpty check all value are empty
func IsEmpty(a any, b ...any) bool {
	b = append(b, a)
	for _, v := range b {
		if !isReflectZero(v) {
			return false
		}
	}
	return true
}

// IsNotEmpty check all values are not empty
func IsNotEmpty(a any, b ...any) bool {
	return !IsEmpty(a, b...)
}

// IfEmptyReturns if A is empty return B value, otherwise return A value
func IfEmptyReturns[T any](a T, b T) T {
	if IsEmpty(a) {
		return b
	}
	return a
}

// ReturnNonEmptyValue returns the information value that is not empty
func ReturnNonEmptyValue[T any](a ...T) T {
	var result T
	for _, v := range a {
		if IsNotNil(v) && IsNotEmpty(v) {
			result = v
			break
		}
	}
	return result
}

func isReflectZero(a any) bool {
	elem := reflect.ValueOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		elem = elem.Elem()
	}
	if IsStringType(a) {
		return len(strings.TrimSpace(CleanAllRepeatSpaces(elem.String()))) == 0
	} else if IsSliceType(a) {
		return elem.Len() == 0
	} else if IsMapType(a) {
		return len(elem.MapKeys()) == 0
	} else if IsObjectIdType(a) {
		return SimpleConvertToObjectId(a).IsZero()
	} else if IsPrimitiveDateTimeType(a) {
		return SimpleConvertToPrimitiveDateTime(a) == 0
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
