package helper

import (
	"github.com/GabrielHCataldo/go-logger/logger"
	"testing"
	"time"
)

func TestIsPointer(t *testing.T) {
	result := IsPointer(initTestStruct())
	logger.Info("IsPointer:", result)
}

func TestIsFunc(t *testing.T) {
	v := func() {}
	result := IsFunc(&v)
	logger.Info("IsFunc:", result)
}

func TestIsChan(t *testing.T) {
	v := make(chan struct{}, 1)
	result := IsChan(&v)
	logger.Info("IsChannel:", result)
}

func TestIsInterface(t *testing.T) {
	v := make(chan struct{}, 1)
	result := IsInterface(&v)
	logger.Info("IsInterface:", result)
}

func TestIsJson(t *testing.T) {
	result := IsJson(initTestStruct())
	logger.Info("IsJson:", result)
}

func TestIsMap(t *testing.T) {
	result := IsMap(initTestMap())
	logger.Info("IsMap:", result)
}

func TestIsStruct(t *testing.T) {
	result := IsStruct(initTestStruct())
	logger.Info("IsStruct:", result)
}

func TestIsSlice(t *testing.T) {
	v := []any{"test", 123}
	result := IsSlice(&v)
	logger.Info("IsSlice:", result)
}

func TestIsString(t *testing.T) {
	v := "string text"
	result := IsString(&v)
	logger.Info("IsString:", result)
}

func TestIsInt(t *testing.T) {
	v := 123
	result := IsInt(&v)
	logger.Info("IsInt:", result)
}

func TestIsInt8(t *testing.T) {
	v := 123
	result := IsInt8(&v)
	logger.Info("IsInt:", result)
}

func TestIsInt16(t *testing.T) {
	v := 123
	result := IsInt16(&v)
	logger.Info("IsInt16:", result)
}

func TestIsInt32(t *testing.T) {
	v := 123
	result := IsInt32(&v)
	logger.Info("IsInt32:", result)
}

func TestIsInt64(t *testing.T) {
	v := 123
	result := IsInt64(&v)
	logger.Info("IsInt64:", result)
}

func TestIsUint(t *testing.T) {
	v := 123
	result := IsUint(&v)
	logger.Info("IsUint:", result)
}

func TestIsUint8(t *testing.T) {
	v := 123
	result := IsUint8(&v)
	logger.Info("IsUint:", result)
}

func TestIsUint16(t *testing.T) {
	v := 123
	result := IsUint16(&v)
	logger.Info("IsUint16:", result)
}

func TestIsUint32(t *testing.T) {
	v := 123
	result := IsUint32(&v)
	logger.Info("IsUint32:", result)
}

func TestIsUint64(t *testing.T) {
	v := 123
	result := IsUint64(&v)
	logger.Info("IsUint64:", result)
}

func TestIsFloat(t *testing.T) {
	v := 123.23
	result := IsFloat(&v)
	logger.Info("IsFloat:", result)
}

func TestIsFloat32(t *testing.T) {
	v := 123.23
	result := IsFloat32(&v)
	logger.Info("IsFloat32:", result)
}

func TestIsFloat64(t *testing.T) {
	v := 123.23
	result := IsFloat64(&v)
	logger.Info("IsFloat64:", result)
}

func TestIsBool(t *testing.T) {
	v := true
	result := IsBool(&v)
	logger.Info("IsBool:", result)
}

func TestIsTime(t *testing.T) {
	v := time.Now()
	result := IsTime(&v)
	logger.Info("IsTime:", result)
}
