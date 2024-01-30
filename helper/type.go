package helper

import (
	"bufio"
	"bytes"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"os"
	"reflect"
	"strings"
	"time"
)

// IsPointer If value is pointer return true, otherwise return false
func IsPointer(a any) bool {
	t := reflect.TypeOf(a)
	return t != nil && t.Kind() == reflect.Pointer
}

// IsNotPointer If value is not pointer return true, otherwise return false
func IsNotPointer(a any) bool {
	return !IsPointer(a)
}

// IsFunc If value is func return true, otherwise return false
func IsFunc(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointer(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Func
}

// IsNotFunc If value is not func return true, otherwise return false
func IsNotFunc(a any) bool {
	return !IsFunc(a)
}

// IsChan If value is chan return true, otherwise return false
func IsChan(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointer(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Chan
}

// IsNotChan If value is not chan return true, otherwise return false
func IsNotChan(a any) bool {
	return !IsChan(a)
}

// IsInterface If value is interface return true, otherwise return false
func IsInterface(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointer(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Interface
}

// IsNotInterface If value is interface return true, otherwise return false
func IsNotInterface(a any) bool {
	return !IsInterface(a)
}

// IsJson If value is struct, map, slice or array return true, otherwise return false
func IsJson(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointer(a) {
		t = t.Elem()
	}
	return IsStruct(a) || IsMap(a) || IsSlice(a)
}

// IsNotJson If value is not struct, map, slice or array return true, otherwise return false
func IsNotJson(a any) bool {
	return !IsJson(a)
}

// IsMap If value is map return true, otherwise return false
func IsMap(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointer(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Map
}

// IsNotMap If value is not map return true, otherwise return false
func IsNotMap(a any) bool {
	return !IsMap(a)
}

// IsStruct If value is struct return true, otherwise return false
func IsStruct(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointer(a) {
		t = t.Elem()
	}
	if IsError(a) || IsTime(a) || IsFile(a) || IsReader(a) || IsBuffer(a) || IsObjectId(a) {
		return false
	}
	return t != nil && t.Kind() == reflect.Struct
}

// IsNotStruct If value is not struct return true, otherwise return false
func IsNotStruct(a any) bool {
	return !IsStruct(a)
}

// IsSlice If value is slice or array return true, otherwise return false
func IsSlice(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointer(a) {
		t = t.Elem()
	}
	if IsError(a) || IsObjectId(a) {
		return false
	}
	return t != nil && (t.Kind() == reflect.Slice || t.Kind() == reflect.Array)
}

// IsNotSlice If value is not slice or array return true, otherwise return false
func IsNotSlice(a any) bool {
	return !IsSlice(a)
}

// IsString If value is string return true, otherwise return false
func IsString(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointer(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.String
}

// IsNotString If value is not string return true, otherwise return false
func IsNotString(a any) bool {
	return !IsString(a)
}

// IsInt If value is int, int8, int16, int32 or int64 return true, otherwise return false
func IsInt(a any) bool {
	if IsNil(a) || IsPrimitiveDateTime(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointer(a) {
		t = t.Elem()
	}
	return t != nil && (t.Kind() == reflect.Int || t.Kind() == reflect.Int8 || t.Kind() == reflect.Int16 ||
		t.Kind() == reflect.Int32 || t.Kind() == reflect.Int64)
}

// IsNotInt If value is not int, int8, int16, int32 or int64 return true, otherwise return false
func IsNotInt(a any) bool {
	return !IsInt(a)
}

// IsInt8 If value is int8 return true, otherwise return false
func IsInt8(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointer(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Int8
}

// IsNotInt8 If value is not int8 return true, otherwise return false
func IsNotInt8(a any) bool {
	return !IsInt8(a)
}

// IsInt16 If value is int16 return true, otherwise return false
func IsInt16(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointer(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Int16
}

// IsNotInt16 If value is not int16 return true, otherwise return false
func IsNotInt16(a any) bool {
	return !IsInt16(a)
}

// IsInt32 If value is int32 return true, otherwise return false
func IsInt32(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointer(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Int32
}

// IsNotInt32 If value is not int32 return true, otherwise return false
func IsNotInt32(a any) bool {
	return !IsInt32(a)
}

// IsInt64 If value is int64 return true, otherwise return false
func IsInt64(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointer(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Int64
}

// IsNotInt64 If value is not int64 return true, otherwise return false
func IsNotInt64(a any) bool {
	return !IsInt64(a)
}

// IsUint If value is uint, uint8, uint16, uint32 or uint64 return true, otherwise return false
func IsUint(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointer(a) {
		t = t.Elem()
	}
	return t != nil && (t.Kind() == reflect.Uint || t.Kind() == reflect.Uint8 || t.Kind() == reflect.Uint16 ||
		t.Kind() == reflect.Uint32 || t.Kind() == reflect.Uint64)
}

// IsNotUint If value is not uint, uint8, uint16, uint32 or uint64 return true, otherwise return false
func IsNotUint(a any) bool {
	return !IsUint(a)
}

// IsUint8 If value is uint8 return true, otherwise return false
func IsUint8(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointer(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Uint8
}

// IsNotUint8 If value is not uint8 return true, otherwise return false
func IsNotUint8(a any) bool {
	return !IsUint8(a)
}

// IsUint16 If value is uint16 return true, otherwise return false
func IsUint16(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointer(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Uint16
}

// IsNotUint16 If value is not uint16 return true, otherwise return false
func IsNotUint16(a any) bool {
	return !IsUint16(a)
}

// IsUint32 If value is uint32 return true, otherwise return false
func IsUint32(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointer(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Uint32
}

// IsNotUint32 If value is not uint32 return true, otherwise return false
func IsNotUint32(a any) bool {
	return !IsUint32(a)
}

// IsUint64 If value is uint64 return true, otherwise return false
func IsUint64(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointer(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Uint64
}

// IsNotUint64 If value is not uint64 return true, otherwise return false
func IsNotUint64(a any) bool {
	return !IsUint64(a)
}

// IsFloat If value is float32 or float64 return true, otherwise return false
func IsFloat(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointer(a) {
		t = t.Elem()
	}
	return t != nil && (t.Kind() == reflect.Float32 || t.Kind() == reflect.Float64)
}

// IsNotFloat If value is not float32 or float64 return true, otherwise return false
func IsNotFloat(a any) bool {
	return !IsFloat(a)
}

// IsFloat32 If value is float32 return true, otherwise return false
func IsFloat32(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointer(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Float32
}

// IsNotFloat32 If value is not float32 return true, otherwise return false
func IsNotFloat32(a any) bool {
	return !IsFloat32(a)
}

// IsFloat64 If value is float64 return true, otherwise return false
func IsFloat64(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointer(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Float64
}

// IsNotFloat64 If value is not float64 return true, otherwise return false
func IsNotFloat64(a any) bool {
	return !IsFloat64(a)
}

// IsBool If value is bool return true, otherwise return false
func IsBool(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointer(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Bool
}

// IsNotBool If value is not bool return true, otherwise return false
func IsNotBool(a any) bool {
	return !IsBool(a)
}

// IsTime If value is time return true, otherwise return false
func IsTime(a any) bool {
	if IsNil(a) {
		return false
	}
	vr := reflect.ValueOf(a)
	if IsPointer(a) {
		vr = vr.Elem()
	}
	return vr.CanConvert(reflect.TypeOf(time.Time{}))
}

// IsNotTime If value is not time return true, otherwise return false
func IsNotTime(a any) bool {
	return !IsTime(a)
}

// IsBytes If value is slice byte return true, otherwise return false
func IsBytes(a any) bool {
	if IsNil(a) {
		return false
	}
	vr := reflect.ValueOf(a)
	if IsPointer(a) {
		vr = vr.Elem()
	}
	_, ok := vr.Interface().([]byte)
	return ok
}

// IsNotBytes If value is not slice byte return true, otherwise return false
func IsNotBytes(a any) bool {
	return !IsBytes(a)
}

// IsError If value is error return true, otherwise return false
func IsError(a any) bool {
	if IsNil(a) {
		return false
	}
	vr := reflect.ValueOf(a)
	if IsPointer(a) {
		vr = vr.Elem()
	}
	return okCastError(a) || okCastError(vr.Interface())
}

// IsNotError If value is not error return true, otherwise return false
func IsNotError(a any) bool {
	return !IsError(a)
}

// IsFile If value is os.File return true, otherwise return false
func IsFile(a any) bool {
	if IsNil(a) {
		return false
	}
	vr := reflect.ValueOf(a)
	if IsPointer(a) {
		vr = vr.Elem()
	}
	_, ok := vr.Interface().(os.File)
	return ok || vr.CanConvert(reflect.TypeOf(os.File{}))
}

// IsNotFile If value is os.File return true, otherwise return false
func IsNotFile(a any) bool {
	return !IsFile(a)
}

// IsReader If value is io.Reader return true, otherwise return false
func IsReader(a any) bool {
	if IsNil(a) {
		return false
	}
	vr := reflect.ValueOf(a)
	if IsPointer(a) {
		vr = vr.Elem()
	}
	_, ok := vr.Interface().(bytes.Reader)
	_, ok2 := vr.Interface().(strings.Reader)
	_, ok3 := vr.Interface().(bufio.Reader)
	_, ok4 := vr.Interface().(io.Reader)
	return ok || ok2 || ok3 || ok4
}

// IsNotReader If value is io.Reader return true, otherwise return false
func IsNotReader(a any) bool {
	return !IsReader(a)
}

// IsBuffer If value is bytes.Buffer return true, otherwise return false
func IsBuffer(a any) bool {
	if IsNil(a) {
		return false
	}
	vr := reflect.ValueOf(a)
	if IsPointer(a) {
		vr = vr.Elem()
	}
	_, ok := vr.Interface().(bytes.Buffer)
	return ok || vr.CanConvert(reflect.TypeOf(bytes.Buffer{}))
}

// IsNotBuffer If value is bytes.Buffer return true, otherwise return false
func IsNotBuffer(a any) bool {
	return !IsBuffer(a)
}

// IsObjectId If value is primitive.ObjectID return true, otherwise return false
func IsObjectId(a any) bool {
	if IsNil(a) {
		return false
	}
	r := reflect.ValueOf(a)
	if IsPointer(a) {
		r = r.Elem()
	}
	return r.Type().String() == "primitive.ObjectID" && r.CanConvert(reflect.TypeOf(primitive.ObjectID{}))
}

// IsNotObjectId If value is not primitive.ObjectID return true, otherwise return false
func IsNotObjectId(a any) bool {
	return !IsObjectId(a)
}

// IsPrimitiveDateTime If value is primitive.DateTime return true, otherwise return false
func IsPrimitiveDateTime(a any) bool {
	if IsNil(a) {
		return false
	}
	r := reflect.ValueOf(a)
	if IsPointer(a) {
		r = r.Elem()
	}
	return r.Type().String() == "primitive.DateTime" &&
		r.CanConvert(reflect.TypeOf(primitive.NewDateTimeFromTime(time.Time{})))
}

// IsNotPrimitiveDateTime If value is not primitive.DateTime return true, otherwise return false
func IsNotPrimitiveDateTime(a any) bool {
	return !IsPrimitiveDateTime(a)
}

func okCastError(a any) bool {
	_, ok := a.(error)
	return ok
}
