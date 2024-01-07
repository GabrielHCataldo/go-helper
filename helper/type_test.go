package helper

import (
	"errors"
	"github.com/GabrielHCataldo/go-logger/logger"
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
	v := func() {}
	result := IsFunc(&v)
	logger.Info("IsFunc:", result)
}

func TestIsNotFunc(t *testing.T) {
	v := func() {}
	result := IsNotFunc(&v)
	logger.Info("IsNotFunc:", result)
}

func TestIsChan(t *testing.T) {
	v := make(chan struct{}, 1)
	result := IsChan(&v)
	logger.Info("IsChannel:", result)
}

func TestIsNotChan(t *testing.T) {
	v := func() {}
	result := IsNotChan(&v)
	logger.Info("IsNotChan:", result)
}

func TestIsInterface(t *testing.T) {
	v := make(chan struct{}, 1)
	result := IsInterface(&v)
	logger.Info("IsInterface:", result)
}

func TestIsNotInterface(t *testing.T) {
	v := func() {}
	result := IsNotInterface(&v)
	logger.Info("IsNotInterface:", result)
}

func TestIsJson(t *testing.T) {
	result := IsJson(initTestStruct())
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
	result := IsMap(initTestMap())
	logger.Info("IsMap:", result)
}

func TestIsNotMap(t *testing.T) {
	v := func() {}
	result := IsNotMap(&v)
	logger.Info("IsNotMap:", result)
}

func TestIsStruct(t *testing.T) {
	result := IsStruct(initTestStruct())
	logger.Info("IsStruct:", result)
}

func TestIsNotStruct(t *testing.T) {
	v := func() {}
	result := IsNotStruct(&v)
	logger.Info("IsNotStruct:", result)
}

func TestIsSlice(t *testing.T) {
	v := []any{"test", 123}
	result := IsSlice(&v)
	logger.Info("IsSlice:", result)
}

func TestIsNotSlice(t *testing.T) {
	v := func() {}
	result := IsNotSlice(&v)
	logger.Info("IsNotSlice:", result)
}

func TestIsString(t *testing.T) {
	v := "string text"
	result := IsString(&v)
	logger.Info("IsString:", result)
}

func TestIsNotString(t *testing.T) {
	v := func() {}
	result := IsNotString(&v)
	logger.Info("IsNotString:", result)
}

func TestIsInt(t *testing.T) {
	v := 123
	result := IsInt(&v)
	logger.Info("IsInt:", result)
}

func TestIsNotInt(t *testing.T) {
	v := func() {}
	result := IsNotInt(&v)
	logger.Info("IsNotInt:", result)
}

func TestIsInt8(t *testing.T) {
	v := 123
	result := IsInt8(&v)
	logger.Info("IsInt:", result)
}

func TestIsNotInt8(t *testing.T) {
	v := func() {}
	result := IsNotInt8(&v)
	logger.Info("IsNotInt8:", result)
}

func TestIsInt16(t *testing.T) {
	v := 123
	result := IsInt16(&v)
	logger.Info("IsInt16:", result)
}

func TestIsNotInt16(t *testing.T) {
	v := func() {}
	result := IsNotInt16(&v)
	logger.Info("IsNotInt16:", result)
}

func TestIsInt32(t *testing.T) {
	v := 123
	result := IsInt32(&v)
	logger.Info("IsInt32:", result)
}

func TestIsNotInt32(t *testing.T) {
	v := func() {}
	result := IsNotInt32(&v)
	logger.Info("IsNotInt32:", result)
}

func TestIsInt64(t *testing.T) {
	v := 123
	result := IsInt64(&v)
	logger.Info("IsInt64:", result)
}

func TestIsNotInt64(t *testing.T) {
	v := func() {}
	result := IsNotInt64(&v)
	logger.Info("IsNotInt64:", result)
}

func TestIsUint(t *testing.T) {
	v := 123
	result := IsUint(&v)
	logger.Info("IsUint:", result)
}

func TestIsNotUint(t *testing.T) {
	v := 123
	result := IsNotUint(&v)
	logger.Info("IsNotUint:", result)
}

func TestIsUint8(t *testing.T) {
	v := 123
	result := IsUint8(&v)
	logger.Info("IsUint:", result)
}

func TestIsNotUint8(t *testing.T) {
	v := 123
	result := IsNotUint8(&v)
	logger.Info("IsNotUint8:", result)
}

func TestIsUint16(t *testing.T) {
	v := 123
	result := IsUint16(&v)
	logger.Info("IsUint16:", result)
}

func TestIsNotUint16(t *testing.T) {
	v := 123
	result := IsNotUint16(&v)
	logger.Info("IsNotUint16:", result)
}

func TestIsUint32(t *testing.T) {
	v := 123
	result := IsUint32(&v)
	logger.Info("IsUint32:", result)
}

func TestIsNotUint32(t *testing.T) {
	v := 123
	result := IsNotUint32(&v)
	logger.Info("IsNotUint32:", result)
}

func TestIsUint64(t *testing.T) {
	v := 123
	result := IsUint64(&v)
	logger.Info("IsUint64:", result)
}

func TestIsNotUint64(t *testing.T) {
	v := 123
	result := IsNotUint64(&v)
	logger.Info("IsNotUint64:", result)
}

func TestIsFloat(t *testing.T) {
	v := 123.23
	result := IsFloat(&v)
	logger.Info("IsFloat:", result)
}

func TestIsNotFloat(t *testing.T) {
	v := 123.23
	result := IsNotFloat(&v)
	logger.Info("IsNotFloat:", result)
}

func TestIsFloat32(t *testing.T) {
	v := 123.23
	result := IsFloat32(&v)
	logger.Info("IsFloat32:", result)
}

func TestIsNotFloat32(t *testing.T) {
	v := 123.23
	result := IsNotFloat32(&v)
	logger.Info("IsNotFloat32:", result)
}

func TestIsFloat64(t *testing.T) {
	v := 123.23
	result := IsFloat64(&v)
	logger.Info("IsFloat64:", result)
}

func TestIsNotFloat64(t *testing.T) {
	v := 123.23
	result := IsNotFloat64(&v)
	logger.Info("IsNotFloat64:", result)
}

func TestIsBool(t *testing.T) {
	v := true
	result := IsBool(&v)
	logger.Info("IsBool:", result)
}

func TestIsNotBool(t *testing.T) {
	v := true
	result := IsNotBool(&v)
	logger.Info("IsNotBool:", result)
}

func TestIsTime(t *testing.T) {
	v := time.Now()
	result := IsTime(&v)
	logger.Info("IsTime:", result)
}

func TestIsNotTime(t *testing.T) {
	v := time.Now()
	result := IsNotTime(&v)
	logger.Info("IsNotTime:", result)
}

func TestIsBytes(t *testing.T) {
	var v []byte
	result := IsBytes(&v)
	logger.Info("IsBytes:", result)
}

func TestIsNotBytes(t *testing.T) {
	var v []byte
	result := IsNotBytes(&v)
	logger.Info("IsNotBytes:", result)
}

func TestIsError(t *testing.T) {
	var v any
	v = errors.New("test error value")
	result := IsError(v)
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
	result := IsFile(v)
	logger.Info("IsFile:", result)
}

func TestIsNotFile(t *testing.T) {
	var v any
	v, _ = os.Open("../gopher-helper.png")
	result := IsNotFile(&v)
	logger.Info("IsNotFile:", result)
}
