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

// IsNotPointer If value is not pointer return true, otherwise return false
func IsNotPointer(v any) bool {
	return !IsPointer(v)
}

// IsFunc If value is func return true, otherwise return false
func IsFunc(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Func
}

// IsNotFunc If value is not func return true, otherwise return false
func IsNotFunc(v any) bool {
	return !IsFunc(v)
}

// IsChan If value is chan return true, otherwise return false
func IsChan(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Chan
}

// IsNotChan If value is not chan return true, otherwise return false
func IsNotChan(v any) bool {
	return !IsChan(v)
}

// IsInterface If value is interface return true, otherwise return false
func IsInterface(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Interface
}

// IsNotInterface If value is interface return true, otherwise return false
func IsNotInterface(v any) bool {
	return !IsInterface(v)
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

// IsNotJson If value is not struct, map, slice or array return true, otherwise return false
func IsNotJson(v any) bool {
	return !IsJson(v)
}

// IsMap If value is map return true, otherwise return false
func IsMap(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Map
}

// IsNotMap If value is not map return true, otherwise return false
func IsNotMap(v any) bool {
	return !IsMap(v)
}

// IsStruct If value is struct return true, otherwise return false
func IsStruct(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Struct
}

// IsNotStruct If value is not struct return true, otherwise return false
func IsNotStruct(v any) bool {
	return !IsStruct(v)
}

// IsSlice If value is slice or array return true, otherwise return false
func IsSlice(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && (t.Kind() == reflect.Slice || t.Kind() == reflect.Array)
}

// IsNotSlice If value is not slice or array return true, otherwise return false
func IsNotSlice(v any) bool {
	return !IsSlice(v)
}

// IsString If value is string return true, otherwise return false
func IsString(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.String
}

// IsNotString If value is not string return true, otherwise return false
func IsNotString(v any) bool {
	return !IsString(v)
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

// IsNotInt If value is not int, int8, int16, int32 or int64 return true, otherwise return false
func IsNotInt(v any) bool {
	return !IsInt(v)
}

// IsInt8 If value is int8 return true, otherwise return false
func IsInt8(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Int8
}

// IsNotInt8 If value is not int8 return true, otherwise return false
func IsNotInt8(v any) bool {
	return !IsInt8(v)
}

// IsInt16 If value is int16 return true, otherwise return false
func IsInt16(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Int16
}

// IsNotInt16 If value is not int16 return true, otherwise return false
func IsNotInt16(v any) bool {
	return !IsInt16(v)
}

// IsInt32 If value is int32 return true, otherwise return false
func IsInt32(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Int32
}

// IsNotInt32 If value is not int32 return true, otherwise return false
func IsNotInt32(v any) bool {
	return !IsInt32(v)
}

// IsInt64 If value is int64 return true, otherwise return false
func IsInt64(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Int64
}

// IsNotInt64 If value is not int64 return true, otherwise return false
func IsNotInt64(v any) bool {
	return !IsInt64(v)
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

// IsNotUint If value is not uint, uint8, uint16, uint32 or uint64 return true, otherwise return false
func IsNotUint(v any) bool {
	return !IsUint(v)
}

// IsUint8 If value is uint8 return true, otherwise return false
func IsUint8(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Uint8
}

// IsNotUint8 If value is not uint8 return true, otherwise return false
func IsNotUint8(v any) bool {
	return !IsUint8(v)
}

// IsUint16 If value is uint16 return true, otherwise return false
func IsUint16(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Uint16
}

// IsNotUint16 If value is not uint16 return true, otherwise return false
func IsNotUint16(v any) bool {
	return !IsUint16(v)
}

// IsUint32 If value is uint32 return true, otherwise return false
func IsUint32(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Uint32
}

// IsNotUint32 If value is not uint32 return true, otherwise return false
func IsNotUint32(v any) bool {
	return !IsUint32(v)
}

// IsUint64 If value is uint64 return true, otherwise return false
func IsUint64(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Uint64
}

// IsNotUint64 If value is not uint64 return true, otherwise return false
func IsNotUint64(v any) bool {
	return !IsUint64(v)
}

// IsFloat If value is float32 or float64 return true, otherwise return false
func IsFloat(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && (t.Kind() == reflect.Float32 || t.Kind() == reflect.Float64)
}

// IsNotFloat If value is not float32 or float64 return true, otherwise return false
func IsNotFloat(v any) bool {
	return !IsFloat(v)
}

// IsFloat32 If value is float32 return true, otherwise return false
func IsFloat32(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Float32
}

// IsNotFloat32 If value is not float32 return true, otherwise return false
func IsNotFloat32(v any) bool {
	return !IsFloat32(v)
}

// IsFloat64 If value is float64 return true, otherwise return false
func IsFloat64(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Float64
}

// IsNotFloat64 If value is not float64 return true, otherwise return false
func IsNotFloat64(v any) bool {
	return !IsFloat64(v)
}

// IsBool If value is bool return true, otherwise return false
func IsBool(v any) bool {
	t := reflect.TypeOf(v)
	if IsPointer(v) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Bool
}

// IsNotBool If value is not bool return true, otherwise return false
func IsNotBool(v any) bool {
	return !IsBool(v)
}

// IsTime If value is time return true, otherwise return false
func IsTime(v any) bool {
	vr := reflect.ValueOf(v)
	if IsPointer(v) {
		vr = vr.Elem()
	}
	_, ok := vr.Interface().(time.Time)
	return ok
}

// IsNotTime If value is not time return true, otherwise return false
func IsNotTime(v any) bool {
	return !IsTime(v)
}

// IsBytes If value is slice byte return true, otherwise return false
func IsBytes(v any) bool {
	vr := reflect.ValueOf(v)
	if IsPointer(v) {
		vr = vr.Elem()
	}
	_, ok := vr.Interface().([]byte)
	return ok
}

// IsNotBytes If value is not slice byte return true, otherwise return false
func IsNotBytes(v any) bool {
	return !IsBytes(v)
}
