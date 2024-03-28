package helper

import (
	"bufio"
	"bytes"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// IsJson check if value is json valid return true, otherwise return false.
func IsJson(a any) bool {
	bs := SimpleConvertToBytes(a)
	var js any
	return json.Unmarshal(bs, &js) == nil
}

// IsNotJson check if value is not json return true, otherwise return false.
func IsNotJson(a any) bool {
	return !IsJson(a)
}

// IsMap check if value is map return true, otherwise return false.
func IsMap(a any) bool {
	bs := SimpleConvertToBytes(a)
	var js map[string]any
	return json.Unmarshal(bs, &js) == nil
}

// IsNotMap check if value is not map return true, otherwise return false.
func IsNotMap(a any) bool {
	return !IsMap(a)
}

// IsSlice check if value is slice return true, otherwise return false.
func IsSlice(a any) bool {
	bs := SimpleConvertToBytes(a)
	var slice []any
	return json.Unmarshal(bs, &slice) == nil
}

// IsNotSlice check if value is not slice return true, otherwise return false.
func IsNotSlice(a any) bool {
	return !IsSlice(a)
}

// IsSliceOfMaps check if string value is slice of maps return true, otherwise return false.
func IsSliceOfMaps(a any) bool {
	bs := SimpleConvertToBytes(a)
	var slice []map[string]any
	return json.Unmarshal(bs, &slice) == nil
}

// IsNotSliceOfMaps check if value is not slice of maps return true, otherwise return false.
func IsNotSliceOfMaps(a any) bool {
	return !IsSliceOfMaps(a)
}

// IsInt check if value is int return true, otherwise return false.
func IsInt(a any) bool {
	s := SimpleConvertToString(a)
	_, err := strconv.Atoi(s)
	return err == nil
}

// IsNotInt check if value is not int return true, otherwise return false.
func IsNotInt(a any) bool {
	return !IsInt(a)
}

// IsBool check if value is bool return true, otherwise return false.
func IsBool(a any) bool {
	s := SimpleConvertToString(a)
	_, err := strconv.ParseBool(s)
	return err == nil
}

// IsNotBool check if value is not bool return true, otherwise return false.
func IsNotBool(a any) bool {
	return !IsBool(a)
}

// IsFloat check if value is float return true, otherwise return false.
func IsFloat(a any) bool {
	s := SimpleConvertToString(a)
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// IsNotFloat check if value is not float return true, otherwise return false.
func IsNotFloat(a any) bool {
	return !IsFloat(a)
}

// IsTime check if value is time return true, otherwise return false.
func IsTime(a any) bool {
	s := SimpleConvertToString(a)
	_, err := ConvertToTime(s)
	return err == nil
}

// IsNotTime check if value is not time return true, otherwise return false.
func IsNotTime(a any) bool {
	return !IsTime(a)
}

// IsPointerType If value is pointer return true, otherwise return false
func IsPointerType(a any) bool {
	t := reflect.TypeOf(a)
	return t != nil && t.Kind() == reflect.Pointer
}

// IsNotPointerType If value is not pointer return true, otherwise return false
func IsNotPointerType(a any) bool {
	return !IsPointerType(a)
}

// IsFuncType If value is func return true, otherwise return false
func IsFuncType(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointerType(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Func
}

// IsNotFuncType If value is not func return true, otherwise return false
func IsNotFuncType(a any) bool {
	return !IsFuncType(a)
}

// IsChanType If value is chan return true, otherwise return false
func IsChanType(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointerType(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Chan
}

// IsNotChanType If value is not chan return true, otherwise return false
func IsNotChanType(a any) bool {
	return !IsChanType(a)
}

// IsInterfaceType If value is interface return true, otherwise return false
func IsInterfaceType(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointerType(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Interface
}

// IsNotInterfaceType If value is interface return true, otherwise return false
func IsNotInterfaceType(a any) bool {
	return !IsInterfaceType(a)
}

// IsJsonType If value is struct, map, slice or array return true, otherwise return false
func IsJsonType(a any) bool {
	if IsNil(a) {
		return false
	}
	return IsStructType(a) || IsMapType(a) || IsSliceType(a)
}

// IsNotJsonType If value is not struct, map, slice or array return true, otherwise return false
func IsNotJsonType(a any) bool {
	return !IsJsonType(a)
}

// IsMapType If value is map return true, otherwise return false
func IsMapType(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Map
}

// IsNotMapType If value is not map return true, otherwise return false
func IsNotMapType(a any) bool {
	return !IsMapType(a)
}

// IsStructType If value is struct return true, otherwise return false
func IsStructType(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		t = t.Elem()
	}
	if IsErrorType(a) || IsTimeType(a) || IsFileType(a) || IsReaderType(a) || IsBufferType(a) || IsObjectIdType(a) {
		return false
	}
	return t != nil && t.Kind() == reflect.Struct
}

// IsNotStructType If value is not struct return true, otherwise return false
func IsNotStructType(a any) bool {
	return !IsStructType(a)
}

// IsSliceType If value is slice or array return true, otherwise return false
func IsSliceType(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		t = t.Elem()
	}
	if IsErrorType(a) || IsObjectIdType(a) {
		return false
	}
	return t != nil && (t.Kind() == reflect.Slice || t.Kind() == reflect.Array)
}

// IsNotSliceType If value is not slice or array return true, otherwise return false
func IsNotSliceType(a any) bool {
	return !IsSliceType(a)
}

// IsSliceOfMapsType If value is slice or array of maps return true, otherwise return false
func IsSliceOfMapsType(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		t = t.Elem()
	}
	if IsNotSliceType(a) {
		return false
	}
	return t != nil && (t.Elem().Kind() == reflect.Map)
}

// IsNotSliceOfMapsType If value is not slice or array of maps return true, otherwise return false
func IsNotSliceOfMapsType(a any) bool {
	return !IsSliceOfMapsType(a)
}

// IsStringType If value is string return true, otherwise return false
func IsStringType(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.String
}

// IsNotStringType If value is not string return true, otherwise return false
func IsNotStringType(a any) bool {
	return !IsStringType(a)
}

// IsIntType If value is int, int8, int16, int32 or int64 return true, otherwise return false
func IsIntType(a any) bool {
	if IsNil(a) || IsPrimitiveDateTimeType(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		t = t.Elem()
	}
	return t != nil && (t.Kind() == reflect.Int || t.Kind() == reflect.Int8 || t.Kind() == reflect.Int16 ||
		t.Kind() == reflect.Int32 || t.Kind() == reflect.Int64)
}

// IsNotIntType If value is not int, int8, int16, int32 or int64 return true, otherwise return false
func IsNotIntType(a any) bool {
	return !IsIntType(a)
}

// IsInt8Type If value is int8 return true, otherwise return false
func IsInt8Type(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Int8
}

// IsNotInt8Type If value is not int8 return true, otherwise return false
func IsNotInt8Type(a any) bool {
	return !IsInt8Type(a)
}

// IsInt16Type If value is int16 return true, otherwise return false
func IsInt16Type(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Int16
}

// IsNotInt16Type If value is not int16 return true, otherwise return false
func IsNotInt16Type(a any) bool {
	return !IsInt16Type(a)
}

// IsInt32Type If value is int32 return true, otherwise return false
func IsInt32Type(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Int32
}

// IsNotInt32Type If value is not int32 return true, otherwise return false
func IsNotInt32Type(a any) bool {
	return !IsInt32Type(a)
}

// IsInt64Type If value is int64 return true, otherwise return false
func IsInt64Type(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Int64
}

// IsNotInt64Type If value is not int64 return true, otherwise return false
func IsNotInt64Type(a any) bool {
	return !IsInt64Type(a)
}

// IsUintType If value is uint, uint8, uint16, uint32 or uint64 return true, otherwise return false
func IsUintType(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		t = t.Elem()
	}
	return t != nil && (t.Kind() == reflect.Uint || t.Kind() == reflect.Uint8 || t.Kind() == reflect.Uint16 ||
		t.Kind() == reflect.Uint32 || t.Kind() == reflect.Uint64)
}

// IsNotUintType If value is not uint, uint8, uint16, uint32 or uint64 return true, otherwise return false
func IsNotUintType(a any) bool {
	return !IsUintType(a)
}

// IsUint8Type If value is uint8 return true, otherwise return false
func IsUint8Type(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Uint8
}

// IsNotUint8Type If value is not uint8 return true, otherwise return false
func IsNotUint8Type(a any) bool {
	return !IsUint8Type(a)
}

// IsUint16Type If value is uint16 return true, otherwise return false
func IsUint16Type(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Uint16
}

// IsNotUint16Type If value is not uint16 return true, otherwise return false
func IsNotUint16Type(a any) bool {
	return !IsUint16Type(a)
}

// IsUint32Type If value is uint32 return true, otherwise return false
func IsUint32Type(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Uint32
}

// IsNotUint32Type If value is not uint32 return true, otherwise return false
func IsNotUint32Type(a any) bool {
	return !IsUint32Type(a)
}

// IsUint64Type If value is uint64 return true, otherwise return false
func IsUint64Type(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Uint64
}

// IsNotUint64Type If value is not uint64 return true, otherwise return false
func IsNotUint64Type(a any) bool {
	return !IsUint64Type(a)
}

// IsFloatType If value is float32 or float64 return true, otherwise return false
func IsFloatType(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		t = t.Elem()
	}
	return t != nil && (t.Kind() == reflect.Float32 || t.Kind() == reflect.Float64)
}

// IsNotFloatType If value is not float32 or float64 return true, otherwise return false
func IsNotFloatType(a any) bool {
	return !IsFloatType(a)
}

// IsFloat32Type If value is float32 return true, otherwise return false
func IsFloat32Type(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Float32
}

// IsNotFloat32Type If value is not float32 return true, otherwise return false
func IsNotFloat32Type(a any) bool {
	return !IsFloat32Type(a)
}

// IsFloat64Type If value is float64 return true, otherwise return false
func IsFloat64Type(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Float64
}

// IsNotFloat64Type If value is not float64 return true, otherwise return false
func IsNotFloat64Type(a any) bool {
	return !IsFloat64Type(a)
}

// IsBoolType If value is bool return true, otherwise return false
func IsBoolType(a any) bool {
	if IsNil(a) {
		return false
	}
	t := reflect.TypeOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		t = t.Elem()
	}
	return t != nil && t.Kind() == reflect.Bool
}

// IsNotBoolType If value is not bool return true, otherwise return false
func IsNotBoolType(a any) bool {
	return !IsBoolType(a)
}

// IsTimeType If value is time return true, otherwise return false
func IsTimeType(a any) bool {
	if IsNil(a) {
		return false
	}
	vr := reflect.ValueOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		vr = vr.Elem()
	}
	return vr.CanConvert(reflect.TypeOf(time.Time{}))
}

// IsNotTimeType If value is not time return true, otherwise return false
func IsNotTimeType(a any) bool {
	return !IsTimeType(a)
}

// IsBytesType If value is slice byte return true, otherwise return false
func IsBytesType(a any) bool {
	if IsNil(a) {
		return false
	}
	vr := reflect.ValueOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		vr = vr.Elem()
	}
	_, ok := vr.Interface().([]byte)
	return ok
}

// IsNotBytesType If value is not slice byte return true, otherwise return false
func IsNotBytesType(a any) bool {
	return !IsBytesType(a)
}

// IsErrorType If value is error return true, otherwise return false
func IsErrorType(a any) bool {
	if IsNil(a) {
		return false
	}
	vr := reflect.ValueOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		vr = vr.Elem()
	}
	return okCastError(a) || okCastError(vr.Interface())
}

// IsNotErrorType If value is not error return true, otherwise return false
func IsNotErrorType(a any) bool {
	return !IsErrorType(a)
}

// IsFileType If value is os.File return true, otherwise return false
func IsFileType(a any) bool {
	if IsNil(a) {
		return false
	}
	vr := reflect.ValueOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		vr = vr.Elem()
	}
	_, ok := vr.Interface().(os.File)
	return ok || vr.CanConvert(reflect.TypeOf(os.File{}))
}

// IsNotFileType If value is os.File return true, otherwise return false
func IsNotFileType(a any) bool {
	return !IsFileType(a)
}

// IsReaderType If value is io.Reader return true, otherwise return false
func IsReaderType(a any) bool {
	if IsNil(a) {
		return false
	}
	vr := reflect.ValueOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		vr = vr.Elem()
	}
	_, ok := vr.Interface().(bytes.Reader)
	_, ok2 := vr.Interface().(strings.Reader)
	_, ok3 := vr.Interface().(bufio.Reader)
	_, ok4 := vr.Interface().(io.Reader)
	return ok || ok2 || ok3 || ok4
}

// IsNotReaderType If value is io.Reader return true, otherwise return false
func IsNotReaderType(a any) bool {
	return !IsReaderType(a)
}

// IsBufferType If value is bytes.Buffer return true, otherwise return false
func IsBufferType(a any) bool {
	if IsNil(a) {
		return false
	}
	vr := reflect.ValueOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		vr = vr.Elem()
	}
	_, ok := vr.Interface().(bytes.Buffer)
	return ok || vr.CanConvert(reflect.TypeOf(bytes.Buffer{}))
}

// IsNotBufferType If value is bytes.Buffer return true, otherwise return false
func IsNotBufferType(a any) bool {
	return !IsBufferType(a)
}

// IsObjectIdType If value is primitive.ObjectID return true, otherwise return false
func IsObjectIdType(a any) bool {
	if IsNil(a) {
		return false
	}
	r := reflect.ValueOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		r = r.Elem()
	}
	s, ok := r.Interface().(string)
	if ok {
		_, err := primitive.ObjectIDFromHex(s)
		if IsNil(err) {
			return true
		}
	}
	return r.Type().String() == "primitive.ObjectID" && r.CanConvert(reflect.TypeOf(primitive.ObjectID{}))
}

// IsNotObjectIdType If value is not primitive.ObjectID return true, otherwise return false
func IsNotObjectIdType(a any) bool {
	return !IsObjectIdType(a)
}

// IsPrimitiveDateTimeType If value is primitive.DateTime return true, otherwise return false
func IsPrimitiveDateTimeType(a any) bool {
	if IsNil(a) {
		return false
	}
	r := reflect.ValueOf(a)
	if IsPointerType(a) || IsInterfaceType(a) {
		r = r.Elem()
	}
	return r.Type().String() == "primitive.DateTime" &&
		r.CanConvert(reflect.TypeOf(primitive.NewDateTimeFromTime(time.Time{})))
}

// IsNotPrimitiveDateTimeType If value is not primitive.DateTime return true, otherwise return false
func IsNotPrimitiveDateTimeType(a any) bool {
	return !IsPrimitiveDateTimeType(a)
}

func okCastError(a any) bool {
	_, ok := a.(error)
	return ok
}
