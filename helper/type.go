package helper

import (
	"reflect"
	"time"
)

// IsPointer If value is pointer return true, otherwise return false
func IsPointer(v any) bool {
	t := reflect.TypeOf(v)
	return t != nil && t.Kind() == reflect.Pointer
}

// IsFunc If value is func return true, otherwise return false
func IsFunc(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Func
}

// IsChan If value is chan return true, otherwise return false
func IsChan(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Chan
}

// IsInterface If value is interface return true, otherwise return false
func IsInterface(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Interface
}

// IsJson If value is struct, map, slice or array return true, otherwise return false
func IsJson(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && (t.Kind() == reflect.Struct || t.Kind() == reflect.Map || t.Kind() == reflect.Slice ||
		t.Kind() == reflect.Array)
}

// IsMap If value is map return true, otherwise return false
func IsMap(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Map
}

// IsStruct If value is struct return true, otherwise return false
func IsStruct(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Struct
}

// IsSlice If value is slice or array return true, otherwise return false
func IsSlice(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && (t.Kind() == reflect.Slice || t.Kind() == reflect.Array)
}

// IsString If value is string return true, otherwise return false
func IsString(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.String
}

// IsInt If value is int, int8, int16, int32 or int64 return true, otherwise return false
func IsInt(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && (t.Kind() == reflect.Int || t.Kind() == reflect.Int8 || t.Kind() == reflect.Int16 ||
		t.Kind() == reflect.Int32 || t.Kind() == reflect.Int64)
}

// IsInt8 If value is int8 return true, otherwise return false
func IsInt8(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Int8
}

// IsInt16 If value is int16 return true, otherwise return false
func IsInt16(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Int16
}

// IsInt32 If value is int32 return true, otherwise return false
func IsInt32(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Int32
}

// IsInt64 If value is int64 return true, otherwise return false
func IsInt64(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Int64
}

// IsUint If value is uint, uint8, uint16, uint32 or uint64 return true, otherwise return false
func IsUint(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && (t.Kind() == reflect.Uint || t.Kind() == reflect.Uint8 || t.Kind() == reflect.Uint16 ||
		t.Kind() == reflect.Uint32 || t.Kind() == reflect.Uint64)
}

// IsUint8 If value is uint8 return true, otherwise return false
func IsUint8(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Uint8
}

// IsUint16 If value is uint16 return true, otherwise return false
func IsUint16(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Uint16
}

// IsUint32 If value is uint32 return true, otherwise return false
func IsUint32(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Uint32
}

// IsUint64 If value is uint64 return true, otherwise return false
func IsUint64(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Uint64
}

// IsFloat If value is float32 or float64 return true, otherwise return false
func IsFloat(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && (t.Kind() == reflect.Float32 || t.Kind() == reflect.Float64)
}

// IsFloat32 If value is float32 return true, otherwise return false
func IsFloat32(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Float32
}

// IsFloat64 If value is float64 return true, otherwise return false
func IsFloat64(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Float64
}

// IsBool If value is bool return true, otherwise return false
func IsBool(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Bool
}

// IsTime If value is time return true, otherwise return false
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
