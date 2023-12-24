package helper

import (
	"reflect"
	"time"
)

func IsPointer(v any) bool {
	t := reflect.TypeOf(v)
	return t.Kind() == reflect.Pointer
}

func IsFunc(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t.Kind() == reflect.Func
}

func IsChannel(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t.Kind() == reflect.Chan
}

func IsInterface(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t.Kind() == reflect.Interface
}

func IsJson(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	switch t.Kind() {
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		return true
	default:
		return false
	}
}

func IsMap(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t.Kind() == reflect.Map
}

func IsStruct(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t.Kind() == reflect.Struct
}

func IsSlice(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	switch t.Kind() {
	case reflect.Slice, reflect.Array:
		return true
	default:
		return false
	}
}

func IsString(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t.Kind() == reflect.String
}

func IsInt(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	default:
		return false
	}
}

func IsInt8(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t.Kind() == reflect.Int8
}

func IsInt16(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t.Kind() == reflect.Int16
}

func IsInt32(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t.Kind() == reflect.Int32
}

func IsInt64(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t.Kind() == reflect.Int64
}

func IsUint(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	switch t.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	default:
		return false
	}
}

func IsUint8(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t.Kind() == reflect.Uint8
}

func IsUint16(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t.Kind() == reflect.Uint16
}

func IsUint32(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t.Kind() == reflect.Uint32
}

func IsUint64(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t.Kind() == reflect.Uint64
}

func IsFloat(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

func IsBool(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t.Kind() == reflect.Bool
}

func IsTime(v any) bool {
	t := reflect.TypeOf(v)
	vr := reflect.ValueOf(v)
	if IsPointer(v) {
		t = t.Elem()
		vr = vr.Elem()
	}
	_, ok := vr.Interface().(time.Time)
	return t.Kind() == reflect.Struct && ok
}
