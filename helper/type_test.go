package helper

import (
	"bytes"
	"errors"
	"github.com/GabrielHCataldo/go-logger/logger"
	"io"
	"os"
	"testing"
	"time"
)

func TestIsPointer(t *testing.T) {
	result := IsPointer(initTestStruct())
	logger.Info("IsPointer:", result)
}

func TestIsNotPointer(t *testing.T) {
	result := IsNotPointer(initTestStruct())
	logger.Info("IsNotPointer:", result)
}

func TestIsFunc(t *testing.T) {
	var v *func()
	result := IsFunc(v)
	logger.Info("IsFunc:", result)
	v = ConvertToPointer(func() {})
	result = IsFunc(v)
	logger.Info("IsFunc:", result)
}

func TestIsNotFunc(t *testing.T) {
	v := func() {}
	result := IsNotFunc(&v)
	logger.Info("IsNotFunc:", result)
}

func TestIsChan(t *testing.T) {
	var v *chan struct{}
	result := IsChan(v)
	logger.Info("IsChannel:", result)
	v = ConvertToPointer(make(chan struct{}))
	result = IsChan(v)
	logger.Info("IsChannel:", result)
}

func TestIsNotChan(t *testing.T) {
	v := func() {}
	result := IsNotChan(&v)
	logger.Info("IsNotChan:", result)
}

func TestIsInterface(t *testing.T) {
	var v *interface{}
	result := IsInterface(v)
	logger.Info("IsInterface:", result)
	v2 := make(chan interface{}, 1)
	result = IsInterface(&v2)
	logger.Info("IsInterface:", result)
}

func TestIsNotInterface(t *testing.T) {
	v := func() {}
	result := IsNotInterface(&v)
	logger.Info("IsNotInterface:", result)
}

func TestIsJson(t *testing.T) {
	result := IsJson(initPointerNil())
	logger.Info("IsJson:", result)
	result = IsJson(initTestStruct())
	logger.Info("IsJson:", result)
	result = IsJson(time.Now())
	logger.Info("IsJson:", result)
	result = IsJson(errors.New("test error value string"))
	logger.Info("IsJson:", result)
}

func TestIsNotJson(t *testing.T) {
	v := func() {}
	result := IsNotJson(&v)
	logger.Info("IsNotJson:", result)
}

func TestIsMap(t *testing.T) {
	result := IsMap(initPointerNil())
	logger.Info("IsMap:", result)
	result = IsMap(initTestMap())
	logger.Info("IsMap:", result)
}

func TestIsNotMap(t *testing.T) {
	v := func() {}
	result := IsNotMap(&v)
	logger.Info("IsNotMap:", result)
}

func TestIsStruct(t *testing.T) {
	result := IsStruct(initPointerNil())
	logger.Info("IsStruct:", result)
	result = IsStruct(initTestStruct())
	logger.Info("IsStruct:", result)
}

func TestIsNotStruct(t *testing.T) {
	v := func() {}
	result := IsNotStruct(&v)
	logger.Info("IsNotStruct:", result)
}

func TestIsSlice(t *testing.T) {
	result := IsSlice(initPointerNil())
	logger.Info("IsSlice:", result)
	result = IsSlice(ConvertToPointer([]any{"test", 123}))
	logger.Info("IsSlice:", result)
}

func TestIsNotSlice(t *testing.T) {
	v := func() {}
	result := IsNotSlice(&v)
	logger.Info("IsNotSlice:", result)
}

func TestIsString(t *testing.T) {
	result := IsString(initPointerNil())
	logger.Info("IsString:", result)
	result = IsString(ConvertToPointer("string text"))
	logger.Info("IsString:", result)
}

func TestIsNotString(t *testing.T) {
	v := func() {}
	result := IsNotString(&v)
	logger.Info("IsNotString:", result)
}

func TestIsInt(t *testing.T) {
	result := IsInt(initPointerNil())
	logger.Info("IsInt:", result)
	result = IsInt(ConvertToPointer(123))
	logger.Info("IsInt:", result)
}

func TestIsNotInt(t *testing.T) {
	v := func() {}
	result := IsNotInt(&v)
	logger.Info("IsNotInt:", result)
}

func TestIsInt8(t *testing.T) {
	result := IsInt8(initPointerNil())
	logger.Info("IsInt8:", result)
	result = IsInt8(ConvertToPointer(123))
	logger.Info("IsInt8:", result)
}

func TestIsNotInt8(t *testing.T) {
	v := func() {}
	result := IsNotInt8(&v)
	logger.Info("IsNotInt8:", result)
}

func TestIsInt16(t *testing.T) {
	result := IsInt16(initPointerNil())
	logger.Info("IsInt16:", result)
	result = IsInt16(ConvertToPointer(123))
	logger.Info("IsInt16:", result)
}

func TestIsNotInt16(t *testing.T) {
	v := func() {}
	result := IsNotInt16(&v)
	logger.Info("IsNotInt16:", result)
}

func TestIsInt32(t *testing.T) {
	result := IsInt32(initPointerNil())
	logger.Info("IsInt32:", result)
	result = IsInt32(ConvertToPointer(123))
	logger.Info("IsInt32:", result)
}

func TestIsNotInt32(t *testing.T) {
	v := func() {}
	result := IsNotInt32(&v)
	logger.Info("IsNotInt32:", result)
}

func TestIsInt64(t *testing.T) {
	result := IsInt64(initPointerNil())
	logger.Info("IsInt64:", result)
	result = IsInt64(ConvertToPointer(123))
	logger.Info("IsInt64:", result)
}

func TestIsNotInt64(t *testing.T) {
	v := func() {}
	result := IsNotInt64(&v)
	logger.Info("IsNotInt64:", result)
}

func TestIsUint(t *testing.T) {
	result := IsUint(initPointerNil())
	logger.Info("IsUint:", result)
	result = IsUint(ConvertToPointer(123))
	logger.Info("IsUint:", result)
}

func TestIsNotUint(t *testing.T) {
	v := 123
	result := IsNotUint(&v)
	logger.Info("IsNotUint:", result)
}

func TestIsUint8(t *testing.T) {
	result := IsUint8(initPointerNil())
	logger.Info("IsUint8:", result)
	result = IsUint8(ConvertToPointer(123))
	logger.Info("IsUint8:", result)
}

func TestIsNotUint8(t *testing.T) {
	v := 123
	result := IsNotUint8(&v)
	logger.Info("IsNotUint8:", result)
}

func TestIsUint16(t *testing.T) {
	result := IsUint16(initPointerNil())
	logger.Info("IsUint16:", result)
	result = IsUint16(ConvertToPointer(123))
	logger.Info("IsUint16:", result)
}

func TestIsNotUint16(t *testing.T) {
	v := 123
	result := IsNotUint16(&v)
	logger.Info("IsNotUint16:", result)
}

func TestIsUint32(t *testing.T) {
	result := IsUint32(initPointerNil())
	logger.Info("IsUint32:", result)
	result = IsUint32(ConvertToPointer(123))
	logger.Info("IsUint32:", result)
}

func TestIsNotUint32(t *testing.T) {
	v := 123
	result := IsNotUint32(&v)
	logger.Info("IsNotUint32:", result)
}

func TestIsUint64(t *testing.T) {
	result := IsUint64(initPointerNil())
	logger.Info("IsUint64:", result)
	result = IsUint64(ConvertToPointer(123))
	logger.Info("IsUint64:", result)
}

func TestIsNotUint64(t *testing.T) {
	v := 123
	result := IsNotUint64(&v)
	logger.Info("IsNotUint64:", result)
}

func TestIsFloat(t *testing.T) {
	result := IsFloat(initPointerNil())
	logger.Info("IsFloat:", result)
	result = IsFloat(ConvertToPointer(123.23))
	logger.Info("IsFloat:", result)
}

func TestIsNotFloat(t *testing.T) {
	v := 123.23
	result := IsNotFloat(&v)
	logger.Info("IsNotFloat:", result)
}

func TestIsFloat32(t *testing.T) {
	result := IsFloat32(initPointerNil())
	logger.Info("IsFloat32:", result)
	result = IsFloat32(ConvertToPointer(123.23))
	logger.Info("IsFloat32:", result)
}

func TestIsNotFloat32(t *testing.T) {
	v := 123.23
	result := IsNotFloat32(&v)
	logger.Info("IsNotFloat32:", result)
}

func TestIsFloat64(t *testing.T) {
	result := IsFloat64(initPointerNil())
	logger.Info("IsFloat64:", result)
	result = IsFloat32(ConvertToPointer(123.23))
	logger.Info("IsFloat64:", result)
}

func TestIsNotFloat64(t *testing.T) {
	v := 123.23
	result := IsNotFloat64(&v)
	logger.Info("IsNotFloat64:", result)
}

func TestIsBool(t *testing.T) {
	result := IsBool(initPointerNil())
	logger.Info("IsBool:", result)
	result = IsBool(true)
	logger.Info("IsBool:", result)
}

func TestIsNotBool(t *testing.T) {
	v := true
	result := IsNotBool(&v)
	logger.Info("IsNotBool:", result)
}

func TestIsTime(t *testing.T) {
	result := IsTime(initPointerNil())
	logger.Info("IsTime:", result)
	result = IsTime(ConvertToPointer(time.Now()))
	logger.Info("IsTime:", result)
}

func TestIsNotTime(t *testing.T) {
	v := time.Now()
	result := IsNotTime(&v)
	logger.Info("IsNotTime:", result)
}

func TestIsBytes(t *testing.T) {
	result := IsBytes(initPointerNil())
	logger.Info("IsBytes:", result)
	result = IsBytes(ConvertToPointer([]byte{}))
	logger.Info("IsBytes:", result)
}

func TestIsNotBytes(t *testing.T) {
	var v []byte
	result := IsNotBytes(&v)
	logger.Info("IsNotBytes:", result)
}

func TestIsError(t *testing.T) {
	result := IsError(initPointerNil())
	logger.Info("IsError:", result)
	result = IsError(ConvertToPointer(errors.New("test error value")))
	logger.Info("IsError:", result)
}

func TestIsNotError(t *testing.T) {
	v := "test error value"
	result := IsNotError(&v)
	logger.Info("IsNotError:", result)
}

func TestIsFile(t *testing.T) {
	var v any
	v, _ = os.Open("../gopher-helper.png")
	result := IsFile(initPointerNil())
	logger.Info("IsFile:", result)
	result = IsFile(v)
	logger.Info("IsFile:", result)
}

func TestIsNotFile(t *testing.T) {
	var v any
	v, _ = os.Open("../gopher-helper.png")
	result := IsNotFile(&v)
	logger.Info("IsNotFile:", result)
}

func TestIsReader(t *testing.T) {
	bs, _ := os.ReadFile("../gopher-helper.png")
	r := bytes.NewReader(bs)
	var v any
	v = r
	result := IsReader(initPointerNil())
	logger.Info("IsReader:", result)
	result = IsReader(v)
	logger.Info("IsReader:", result)
	result = IsReader(io.Reader(r))
	logger.Info("IsReader:", result)
	result = IsReader(io.NopCloser(r))
	logger.Info("IsReader:", result)
}

func TestIsNotReader(t *testing.T) {
	bs, _ := os.ReadFile("../gopher-helper.png")
	var v any
	v = bytes.NewBuffer(bs)
	result := IsNotReader(v)
	logger.Info("IsNotReader:", result)
}

func TestIsBuffer(t *testing.T) {
	bs, _ := os.ReadFile("../gopher-helper.png")
	var v any
	v = bytes.NewBuffer(bs)
	result := IsBuffer(initPointerNil())
	logger.Info("IsBuffer:", result)
	result = IsBuffer(v)
	logger.Info("IsBuffer:", result)
}

func TestIsNotBuffer(t *testing.T) {
	bs, _ := os.ReadFile("../gopher-helper.png")
	var v any
	v = bytes.NewBuffer(bs)
	result := IsNotBuffer(v)
	logger.Info("IsNotBuffer:", result)
}
