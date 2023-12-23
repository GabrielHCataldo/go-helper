package helper

import (
	"reflect"
	"regexp"
	"strings"
)

func IsNil(v any) bool {
	elem := reflect.ValueOf(v)
	switch elem.Kind() {
	case reflect.Pointer, reflect.Chan, reflect.Map, reflect.Interface, reflect.Slice, reflect.Array:
		return elem.IsNil()
	default:
		return !elem.IsValid()
	}
}

func IsNonNil(v any) bool {
	elem := reflect.ValueOf(v)
	switch elem.Kind() {
	case reflect.Pointer, reflect.Chan, reflect.Map, reflect.Interface, reflect.Slice, reflect.Array:
		return elem.IsNil()
	default:
		return !elem.IsValid()
	}
}

func IsEmpty(v any) bool {
	return isReflectZero(v)
}

func IsNotEmpty(v any) bool {
	return !IsEmpty(v)
}

func IsPointerNil(v any) bool {
	if !IsPointer(v) {
		panic("value type is not a pointer")
	}
	return isReflectNil(v)
}

func IsPointerNonNil(v any) bool {
	return !IsPointerNil(v)
}

func IsJsonEmpty(v any) bool {
	if !IsJson(v) {
		panic("value type is not a struct or map")
	}
	return isReflectZero(v)
}

func IsJsonNotEmpty(v any) bool {
	return !IsJsonEmpty(v)
}

func IsMapEmpty(v any) bool {
	if !IsMap(v) {
		panic("value type is not a map")
	}
	return isReflectZero(v)
}

func IsMapNotEmpty(v any) bool {
	return !IsMapEmpty(v)
}

func IsStructEmpty(v any) bool {
	if !IsStruct(v) {
		panic("value type is not a struct")
	}
	return isReflectZero(v)
}

func IsStructNotEmpty(v any) bool {
	return !IsStructEmpty(v)
}

func IsSliceEmpty(v any) bool {
	if !IsSlice(v) {
		panic("value type is not a slice or array")
	}
	return isReflectZero(v)
}

func IsSliceNotEmpty(v any) bool {
	return !IsSliceEmpty(v)
}

func IsStringEmpty(v any) bool {
	if !IsString(v) {
		panic("value type is not a string")
	}
	elem := reflect.ValueOf(v)
	if IsPointer(v) || IsInterface(v) {
		elem = elem.Elem()
	}
	regex := regexp.MustCompile(`\s+`)
	str := strings.TrimSpace(regex.ReplaceAllString(elem.String(), ""))
	return len(str) == 0
}

func IsStringNotEmpty(v any) bool {
	return !IsStringEmpty(v)
}

func IsIntEmpty(v any) bool {
	if !IsInt(v) {
		panic("value type is not a int")
	}
	return isReflectZero(v)
}

func IsIntNotEmpty(v any) bool {
	return !IsIntEmpty(v)
}

func IsFloatEmpty(v any) bool {
	if !IsFloat(v) {
		panic("value type is not a float")
	}
	return isReflectZero(v)
}

func IsFloatNotEmpty(v any) bool {
	return !IsFloatEmpty(v)
}

func IsBoolEmpty(v any) bool {
	if !IsBool(v) {
		panic("value type is not a bool")
	}
	return isReflectZero(v)
}

func IsBoolNotEmpty(v any) bool {
	return !IsBoolEmpty(v)
}

func isReflectZero(v any) bool {
	elem := reflect.ValueOf(v)
	if IsPointer(v) || IsInterface(v) {
		elem = elem.Elem()
	}
	return !elem.IsValid() || elem.IsZero()
}

func isReflectNil(v any) bool {
	elem := reflect.ValueOf(v)
	return !elem.IsValid() || elem.IsNil()
}
