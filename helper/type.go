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

func IsChan(v any) bool {
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
	return t.Kind() == reflect.Struct || t.Kind() == reflect.Map || t.Kind() == reflect.Slice || t.Kind() == reflect.Array
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
	return t.Kind() == reflect.Slice || t.Kind() == reflect.Array
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
	return t.Kind() == reflect.Int || t.Kind() == reflect.Int8 || t.Kind() == reflect.Int16 ||
		t.Kind() == reflect.Int32 || t.Kind() == reflect.Int64
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
	return t.Kind() == reflect.Uint || t.Kind() == reflect.Uint8 || t.Kind() == reflect.Uint16 ||
		t.Kind() == reflect.Uint32 || t.Kind() == reflect.Uint64
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
	return t.Kind() == reflect.Float32 || t.Kind() == reflect.Float64
}

func IsFloat32(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t.Kind() == reflect.Float32
}

func IsFloat64(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t.Kind() == reflect.Float64
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
	return ok
}
