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
	result := IsPointer(initTestStruct())
	log.Println("IsPointer:", result)
}

func TestIsNotPointer(t *testing.T) {
	result := IsNotPointer(initTestStruct())
	log.Println("IsNotPointer:", result)
}

func TestIsFunc(t *testing.T) {
	var v *func()
	result := IsFunc(v)
	log.Println("IsFunc:", result)
	v = ConvertToPointer(func() {})
	result = IsFunc(v)
	log.Println("IsFunc:", result)
}

func TestIsNotFunc(t *testing.T) {
	v := func() {}
	result := IsNotFunc(&v)
	log.Println("IsNotFunc:", result)
}

func TestIsChan(t *testing.T) {
	var v *chan struct{}
	result := IsChan(v)
	log.Println("IsChannel:", result)
	v = ConvertToPointer(make(chan struct{}))
	result = IsChan(v)
	log.Println("IsChannel:", result)
}

func TestIsNotChan(t *testing.T) {
	v := func() {}
	result := IsNotChan(&v)
	log.Println("IsNotChan:", result)
}

func TestIsInterface(t *testing.T) {
	var v *interface{}
	result := IsInterface(v)
	log.Println("IsInterface:", result)
	v2 := make(chan interface{}, 1)
	result = IsInterface(&v2)
	log.Println("IsInterface:", result)
}

func TestIsNotInterface(t *testing.T) {
	v := func() {}
	result := IsNotInterface(&v)
	log.Println("IsNotInterface:", result)
}

func TestIsJson(t *testing.T) {
	result := IsJson(initPointerNil())
	log.Println("IsJson:", result)
	result = IsJson(initTestStruct())
	log.Println("IsJson:", result)
	result = IsJson(time.Now())
	log.Println("IsJson:", result)
	result = IsJson(errors.New("test error value string"))
	log.Println("IsJson:", result)
}

func TestIsNotJson(t *testing.T) {
	v := func() {}
	result := IsNotJson(&v)
	log.Println("IsNotJson:", result)
}

func TestIsMap(t *testing.T) {
	result := IsMap(initPointerNil())
	log.Println("IsMap:", result)
	result = IsMap(initTestMap())
	log.Println("IsMap:", result)
}

func TestIsNotMap(t *testing.T) {
	v := func() {}
	result := IsNotMap(&v)
	log.Println("IsNotMap:", result)
}

func TestIsStruct(t *testing.T) {
	result := IsStruct(initPointerNil())
	log.Println("IsStruct:", result)
	result = IsStruct(initTestStruct())
	log.Println("IsStruct:", result)
	result = IsStruct(time.Now())
	log.Println("IsStruct:", result)
}

func TestIsNotStruct(t *testing.T) {
	v := func() {}
	result := IsNotStruct(&v)
	log.Println("IsNotStruct:", result)
}

func TestIsSlice(t *testing.T) {
	result := IsSlice(initPointerNil())
	log.Println("IsSlice:", result)
	result = IsSlice(ConvertToPointer([]any{"test", 123}))
	log.Println("IsSlice:", result)
}

func TestIsNotSlice(t *testing.T) {
	v := func() {}
	result := IsNotSlice(&v)
	log.Println("IsNotSlice:", result)
}

func TestIsString(t *testing.T) {
	result := IsString(initPointerNil())
	log.Println("IsString:", result)
	result = IsString(ConvertToPointer("string text"))
	log.Println("IsString:", result)
}

func TestIsNotString(t *testing.T) {
	v := func() {}
	result := IsNotString(&v)
	log.Println("IsNotString:", result)
}

func TestIsInt(t *testing.T) {
	result := IsInt(initPointerNil())
	log.Println("IsInt:", result)
	result = IsInt(ConvertToPointer(123))
	log.Println("IsInt:", result)
}

func TestIsNotInt(t *testing.T) {
	v := func() {}
	result := IsNotInt(&v)
	log.Println("IsNotInt:", result)
}

func TestIsInt8(t *testing.T) {
	result := IsInt8(initPointerNil())
	log.Println("IsInt8:", result)
	result = IsInt8(ConvertToPointer(123))
	log.Println("IsInt8:", result)
}

func TestIsNotInt8(t *testing.T) {
	v := func() {}
	result := IsNotInt8(&v)
	log.Println("IsNotInt8:", result)
}

func TestIsInt16(t *testing.T) {
	result := IsInt16(initPointerNil())
	log.Println("IsInt16:", result)
	result = IsInt16(ConvertToPointer(123))
	log.Println("IsInt16:", result)
}

func TestIsNotInt16(t *testing.T) {
	v := func() {}
	result := IsNotInt16(&v)
	log.Println("IsNotInt16:", result)
}

func TestIsInt32(t *testing.T) {
	result := IsInt32(initPointerNil())
	log.Println("IsInt32:", result)
	result = IsInt32(ConvertToPointer(123))
	log.Println("IsInt32:", result)
}

func TestIsNotInt32(t *testing.T) {
	v := func() {}
	result := IsNotInt32(&v)
	log.Println("IsNotInt32:", result)
}

func TestIsInt64(t *testing.T) {
	result := IsInt64(initPointerNil())
	log.Println("IsInt64:", result)
	result = IsInt64(ConvertToPointer(123))
	log.Println("IsInt64:", result)
}

func TestIsNotInt64(t *testing.T) {
	v := func() {}
	result := IsNotInt64(&v)
	log.Println("IsNotInt64:", result)
}

func TestIsUint(t *testing.T) {
	result := IsUint(initPointerNil())
	log.Println("IsUint:", result)
	result = IsUint(ConvertToPointer(123))
	log.Println("IsUint:", result)
}

func TestIsNotUint(t *testing.T) {
	v := 123
	result := IsNotUint(&v)
	log.Println("IsNotUint:", result)
}

func TestIsUint8(t *testing.T) {
	result := IsUint8(initPointerNil())
	log.Println("IsUint8:", result)
	result = IsUint8(ConvertToPointer(123))
	log.Println("IsUint8:", result)
}

func TestIsNotUint8(t *testing.T) {
	v := 123
	result := IsNotUint8(&v)
	log.Println("IsNotUint8:", result)
}

func TestIsUint16(t *testing.T) {
	result := IsUint16(initPointerNil())
	log.Println("IsUint16:", result)
	result = IsUint16(ConvertToPointer(123))
	log.Println("IsUint16:", result)
}

func TestIsNotUint16(t *testing.T) {
	v := 123
	result := IsNotUint16(&v)
	log.Println("IsNotUint16:", result)
}

func TestIsUint32(t *testing.T) {
	result := IsUint32(initPointerNil())
	log.Println("IsUint32:", result)
	result = IsUint32(ConvertToPointer(123))
	log.Println("IsUint32:", result)
}

func TestIsNotUint32(t *testing.T) {
	v := 123
	result := IsNotUint32(&v)
	log.Println("IsNotUint32:", result)
}

func TestIsUint64(t *testing.T) {
	result := IsUint64(initPointerNil())
	log.Println("IsUint64:", result)
	result = IsUint64(ConvertToPointer(123))
	log.Println("IsUint64:", result)
}

func TestIsNotUint64(t *testing.T) {
	v := 123
	result := IsNotUint64(&v)
	log.Println("IsNotUint64:", result)
}

func TestIsFloat(t *testing.T) {
	result := IsFloat(initPointerNil())
	log.Println("IsFloat:", result)
	result = IsFloat(ConvertToPointer(123.23))
	log.Println("IsFloat:", result)
}

func TestIsNotFloat(t *testing.T) {
	v := 123.23
	result := IsNotFloat(&v)
	log.Println("IsNotFloat:", result)
}

func TestIsFloat32(t *testing.T) {
	result := IsFloat32(initPointerNil())
	log.Println("IsFloat32:", result)
	result = IsFloat32(ConvertToPointer(123.23))
	log.Println("IsFloat32:", result)
}

func TestIsNotFloat32(t *testing.T) {
	v := 123.23
	result := IsNotFloat32(&v)
	log.Println("IsNotFloat32:", result)
}

func TestIsFloat64(t *testing.T) {
	result := IsFloat64(initPointerNil())
	log.Println("IsFloat64:", result)
	result = IsFloat32(ConvertToPointer(123.23))
	log.Println("IsFloat64:", result)
}

func TestIsNotFloat64(t *testing.T) {
	v := 123.23
	result := IsNotFloat64(&v)
	log.Println("IsNotFloat64:", result)
}

func TestIsBool(t *testing.T) {
	result := IsBool(initPointerNil())
	log.Println("IsBool:", result)
	result = IsBool(true)
	log.Println("IsBool:", result)
}

func TestIsNotBool(t *testing.T) {
	v := true
	result := IsNotBool(&v)
	log.Println("IsNotBool:", result)
}

func TestIsTime(t *testing.T) {
	result := IsTime(initPointerNil())
	log.Println("IsTime:", result)
	result = IsTime(ConvertToPointer(time.Now()))
	log.Println("IsTime:", result)
	result = IsTime(primitive.NewDateTimeFromTime(time.Now()))
	log.Println("IsTime:", result)
	result = IsTime(primitive.Timestamp{
		T: uint32(time.Now().Unix()),
		I: 0,
	})
	log.Println("IsTime:", result)
}

func TestIsNotTime(t *testing.T) {
	v := time.Now()
	result := IsNotTime(&v)
	log.Println("IsNotTime:", result)
}

func TestIsBytes(t *testing.T) {
	result := IsBytes(initPointerNil())
	log.Println("IsBytes:", result)
	result = IsBytes(ConvertToPointer([]byte{}))
	log.Println("IsBytes:", result)
}

func TestIsNotBytes(t *testing.T) {
	var v []byte
	result := IsNotBytes(&v)
	log.Println("IsNotBytes:", result)
}

func TestIsError(t *testing.T) {
	result := IsError(initPointerNil())
	log.Println("IsError:", result)
	result = IsError(ConvertToPointer(errors.New("test error value")))
	log.Println("IsError:", result)
}

func TestIsNotError(t *testing.T) {
	v := "test error value"
	result := IsNotError(&v)
	log.Println("IsNotError:", result)
}

func TestIsFile(t *testing.T) {
	var v any
	v, _ = os.Open("../gopher-helper.png")
	result := IsFile(initPointerNil())
	log.Println("IsFile:", result)
	result = IsFile(v)
	log.Println("IsFile:", result)
}

func TestIsNotFile(t *testing.T) {
	var v any
	v, _ = os.Open("../gopher-helper.png")
	result := IsNotFile(&v)
	log.Println("IsNotFile:", result)
}

func TestIsReader(t *testing.T) {
	bs, _ := os.ReadFile("../gopher-helper.png")
	r := bytes.NewReader(bs)
	var v any
	v = r
	result := IsReader(initPointerNil())
	log.Println("IsReader:", result)
	result = IsReader(v)
	log.Println("IsReader:", result)
	result = IsReader(io.Reader(r))
	log.Println("IsReader:", result)
	result = IsReader(io.NopCloser(r))
	log.Println("IsReader:", result)
}

func TestIsNotReader(t *testing.T) {
	bs, _ := os.ReadFile("../gopher-helper.png")
	var v any
	v = bytes.NewBuffer(bs)
	result := IsNotReader(v)
	log.Println("IsNotReader:", result)
}

func TestIsBuffer(t *testing.T) {
	bs, _ := os.ReadFile("../gopher-helper.png")
	var v any
	v = bytes.NewBuffer(bs)
	result := IsBuffer(initPointerNil())
	log.Println("IsBuffer:", result)
	result = IsBuffer(v)
	log.Println("IsBuffer:", result)
}

func TestIsNotBuffer(t *testing.T) {
	bs, _ := os.ReadFile("../gopher-helper.png")
	var v any
	v = bytes.NewBuffer(bs)
	result := IsNotBuffer(v)
	log.Println("IsNotBuffer:", result)
}
