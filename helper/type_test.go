package helper

import (
	"bytes"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"log"
	"os"
	"testing"
	"time"
)

func TestIsPointer(t *testing.T) {
	result := IsPointerType(initTestStruct())
	log.Println("IsPointer:", result)
}

func TestIsNotPointer(t *testing.T) {
	result := IsNotPointerType(initTestStruct())
	log.Println("IsNotPointerType:", result)
}

func TestIsFunc(t *testing.T) {
	var v *func()
	result := IsFuncType(v)
	log.Println("IsFuncType:", result)
	v = ConvertToPointer(func() {})
	result = IsFuncType(v)
	log.Println("IsFuncType:", result)
}

func TestIsNotFunc(t *testing.T) {
	v := func() {}
	result := IsNotFuncType(&v)
	log.Println("IsNotFunc:", result)
}

func TestIsChan(t *testing.T) {
	var v *chan struct{}
	result := IsChanType(v)
	log.Println("IsChannel:", result)
	v = ConvertToPointer(make(chan struct{}))
	result = IsChanType(v)
	log.Println("IsChannel:", result)
}

func TestIsNotChan(t *testing.T) {
	v := func() {}
	result := IsNotChanType(&v)
	log.Println("IsNotChanType:", result)
}

func TestIsInterface(t *testing.T) {
	var v *interface{}
	result := IsInterfaceType(v)
	log.Println("IsInterfaceType:", result)
	v2 := make(chan interface{}, 1)
	result = IsInterfaceType(&v2)
	log.Println("IsInterfaceType:", result)
}

func TestIsNotInterface(t *testing.T) {
	v := func() {}
	result := IsNotInterfaceType(&v)
	log.Println("IsNotInterfaceType:", result)
}

func TestIsJson(t *testing.T) {
	result := IsJsonType(initPointerNil())
	log.Println("IsJsonType:", result)
	result = IsJsonType(initTestStruct())
	log.Println("IsJsonType:", result)
	result = IsJsonType(time.Now())
	log.Println("IsJsonType:", result)
	result = IsJsonType(errors.New("test error value string"))
	log.Println("IsJsonType:", result)
}

func TestIsNotJson(t *testing.T) {
	v := func() {}
	result := IsNotJsonType(&v)
	log.Println("IsNotJsonType:", result)
}

func TestIsMap(t *testing.T) {
	result := IsMapType(initPointerNil())
	log.Println("IsMapType:", result)
	result = IsMapType(initTestMap())
	log.Println("IsMapType:", result)
}

func TestIsNotMap(t *testing.T) {
	v := func() {}
	result := IsNotMapType(&v)
	log.Println("IsNotMap:", result)
}

func TestIsStruct(t *testing.T) {
	result := IsStructType(initPointerNil())
	log.Println("IsStructType:", result)
	result = IsStructType(initTestStruct())
	log.Println("IsStructType:", result)
	result = IsStructType(time.Now())
	log.Println("IsStructType:", result)
}

func TestIsNotStruct(t *testing.T) {
	v := func() {}
	result := IsNotStructType(&v)
	log.Println("IsNotStructType:", result)
}

func TestIsSlice(t *testing.T) {
	result := IsSliceType(initPointerNil())
	log.Println("IsSliceType:", result)
	result = IsSliceType(ConvertToPointer([]any{"test", 123}))
	log.Println("IsSliceType:", result)
}

func TestIsNotSlice(t *testing.T) {
	v := func() {}
	result := IsNotSliceType(&v)
	log.Println("IsNotSlice:", result)
}

func TestIsString(t *testing.T) {
	result := IsStringType(initPointerNil())
	log.Println("IsStringType:", result)
	result = IsStringType(ConvertToPointer("string text"))
	log.Println("IsStringType:", result)
}

func TestIsNotString(t *testing.T) {
	v := func() {}
	result := IsNotStringType(&v)
	log.Println("IsNotString:", result)
}

func TestIsInt(t *testing.T) {
	result := IsIntType(initPointerNil())
	log.Println("IsIntType:", result)
	result = IsIntType(ConvertToPointer(123))
	log.Println("IsIntType:", result)
}

func TestIsNotInt(t *testing.T) {
	v := func() {}
	result := IsNotIntType(&v)
	log.Println("IsNotIntType:", result)
}

func TestIsInt8(t *testing.T) {
	result := IsInt8Type(initPointerNil())
	log.Println("IsInt8Type:", result)
	result = IsInt8Type(ConvertToPointer(123))
	log.Println("IsInt8Type:", result)
}

func TestIsNotInt8(t *testing.T) {
	v := func() {}
	result := IsNotInt8Type(&v)
	log.Println("IsNotInt8Type:", result)
}

func TestIsInt16(t *testing.T) {
	result := IsInt16Type(initPointerNil())
	log.Println("IsInt16Type:", result)
	result = IsInt16Type(ConvertToPointer(123))
	log.Println("IsInt16Type:", result)
}

func TestIsNotInt16(t *testing.T) {
	v := func() {}
	result := IsNotInt16Type(&v)
	log.Println("IsNotInt16Type:", result)
}

func TestIsInt32(t *testing.T) {
	result := IsInt32Type(initPointerNil())
	log.Println("IsInt32Type:", result)
	result = IsInt32Type(ConvertToPointer(123))
	log.Println("IsInt32Type:", result)
}

func TestIsNotInt32(t *testing.T) {
	v := func() {}
	result := IsNotInt32Type(&v)
	log.Println("IsNotInt32Type:", result)
}

func TestIsInt64(t *testing.T) {
	result := IsInt64Type(initPointerNil())
	log.Println("IsInt64Type:", result)
	result = IsInt64Type(ConvertToPointer(123))
	log.Println("IsInt64Type:", result)
}

func TestIsNotInt64(t *testing.T) {
	v := func() {}
	result := IsNotInt64Type(&v)
	log.Println("IsNotInt64Type:", result)
}

func TestIsUint(t *testing.T) {
	result := IsUintType(initPointerNil())
	log.Println("IsUintType:", result)
	result = IsUintType(ConvertToPointer(123))
	log.Println("IsUintType:", result)
}

func TestIsNotUint(t *testing.T) {
	v := 123
	result := IsNotUintType(&v)
	log.Println("IsNotUint:", result)
}

func TestIsUint8(t *testing.T) {
	result := IsUint8Type(initPointerNil())
	log.Println("IsUint8Type:", result)
	result = IsUint8Type(ConvertToPointer(123))
	log.Println("IsUint8Type:", result)
}

func TestIsNotUint8(t *testing.T) {
	v := 123
	result := IsNotUint8Type(&v)
	log.Println("IsNotUint8Type:", result)
}

func TestIsUint16(t *testing.T) {
	result := IsUint16Type(initPointerNil())
	log.Println("IsUint16Type:", result)
	result = IsUint16Type(ConvertToPointer(123))
	log.Println("IsUint16Type:", result)
}

func TestIsNotUint16(t *testing.T) {
	v := 123
	result := IsNotUint16Type(&v)
	log.Println("IsNotUint16Type:", result)
}

func TestIsUint32(t *testing.T) {
	result := IsUint32Type(initPointerNil())
	log.Println("IsUint32Type:", result)
	result = IsUint32Type(ConvertToPointer(123))
	log.Println("IsUint32Type:", result)
}

func TestIsNotUint32(t *testing.T) {
	v := 123
	result := IsNotUint32Type(&v)
	log.Println("IsNotUint32Type:", result)
}

func TestIsUint64(t *testing.T) {
	result := IsUint64Type(initPointerNil())
	log.Println("IsUint64Type:", result)
	result = IsUint64Type(ConvertToPointer(123))
	log.Println("IsUint64Type:", result)
}

func TestIsNotUint64(t *testing.T) {
	v := 123
	result := IsNotUint64Type(&v)
	log.Println("IsNotUint64Type:", result)
}

func TestIsFloat(t *testing.T) {
	result := IsFloatType(initPointerNil())
	log.Println("IsFloatType:", result)
	result = IsFloatType(ConvertToPointer(123.23))
	log.Println("IsFloatType:", result)
}

func TestIsNotFloat(t *testing.T) {
	v := 123.23
	result := IsNotFloatType(&v)
	log.Println("IsNotFloatType:", result)
}

func TestIsFloat32(t *testing.T) {
	result := IsFloat32Type(initPointerNil())
	log.Println("IsFloat32Type:", result)
	result = IsFloat32Type(ConvertToPointer(123.23))
	log.Println("IsFloat32Type:", result)
}

func TestIsNotFloat32(t *testing.T) {
	v := 123.23
	result := IsNotFloat32Type(&v)
	log.Println("IsNotFloat32Type:", result)
}

func TestIsFloat64(t *testing.T) {
	result := IsFloat64Type(initPointerNil())
	log.Println("IsFloat64Type:", result)
	result = IsFloat32Type(ConvertToPointer(123.23))
	log.Println("IsFloat64Type:", result)
}

func TestIsNotFloat64(t *testing.T) {
	v := 123.23
	result := IsNotFloat64Type(&v)
	log.Println("IsNotFloat64Type:", result)
}

func TestIsBool(t *testing.T) {
	result := IsBoolType(initPointerNil())
	log.Println("IsBoolType:", result)
	result = IsBoolType(true)
	log.Println("IsBoolType:", result)
}

func TestIsNotBool(t *testing.T) {
	v := true
	result := IsNotBoolType(&v)
	log.Println("IsNotBoolType:", result)
}

func TestIsTime(t *testing.T) {
	result := IsTimeType(initPointerNil())
	log.Println("IsTimeType:", result)
	result = IsTimeType(ConvertToPointer(time.Now()))
	log.Println("IsTimeType:", result)
	result = IsTimeType(primitive.NewDateTimeFromTime(time.Now()))
	log.Println("IsTimeType:", result)
	result = IsTimeType(primitive.Timestamp{
		T: uint32(time.Now().Unix()),
		I: 0,
	})
	log.Println("IsTimeType:", result)
}

func TestIsNotTime(t *testing.T) {
	v := time.Now()
	result := IsNotTimeType(&v)
	log.Println("IsNotTimeType:", result)
}

func TestIsBytes(t *testing.T) {
	result := IsBytesType(initPointerNil())
	log.Println("IsBytesType:", result)
	result = IsBytesType(ConvertToPointer([]byte{}))
	log.Println("IsBytesType:", result)
}

func TestIsNotBytes(t *testing.T) {
	var v []byte
	result := IsNotBytesType(&v)
	log.Println("IsNotBytesType:", result)
}

func TestIsError(t *testing.T) {
	result := IsErrorType(initPointerNil())
	log.Println("IsErrorType:", result)
	result = IsErrorType(ConvertToPointer(errors.New("test error value")))
	log.Println("IsErrorType:", result)
}

func TestIsNotError(t *testing.T) {
	v := "test error value"
	result := IsNotErrorType(&v)
	log.Println("IsNotErrorType:", result)
}

func TestIsFile(t *testing.T) {
	var v any
	v, _ = os.Open("../gopher-helper.png")
	result := IsFileType(initPointerNil())
	log.Println("IsFileType:", result)
	result = IsFileType(v)
	log.Println("IsFileType:", result)
}

func TestIsNotFile(t *testing.T) {
	var v any
	v, _ = os.Open("../gopher-helper.png")
	result := IsNotFileType(&v)
	log.Println("IsNotFileType:", result)
}

func TestIsReader(t *testing.T) {
	bs, _ := os.ReadFile("../gopher-helper.png")
	r := bytes.NewReader(bs)
	var v any
	v = r
	result := IsReaderType(initPointerNil())
	log.Println("IsReaderType:", result)
	result = IsReaderType(v)
	log.Println("IsReaderType:", result)
	result = IsReaderType(io.Reader(r))
	log.Println("IsReaderType:", result)
	result = IsReaderType(io.NopCloser(r))
	log.Println("IsReaderType:", result)
}

func TestIsNotReader(t *testing.T) {
	bs, _ := os.ReadFile("../gopher-helper.png")
	var v any
	v = bytes.NewBuffer(bs)
	result := IsNotReaderType(v)
	log.Println("IsNotReaderType:", result)
}

func TestIsBuffer(t *testing.T) {
	bs, _ := os.ReadFile("../gopher-helper.png")
	var v any
	v = bytes.NewBuffer(bs)
	result := IsBufferType(initPointerNil())
	log.Println("IsBufferType:", result)
	result = IsBufferType(v)
	log.Println("IsBufferType:", result)
}

func TestIsNotBuffer(t *testing.T) {
	bs, _ := os.ReadFile("../gopher-helper.png")
	var v any
	v = bytes.NewBuffer(bs)
	result := IsNotBufferType(v)
	log.Println("IsNotBufferType:", result)
}
