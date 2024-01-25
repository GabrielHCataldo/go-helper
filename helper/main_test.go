package helper

import (
	"bufio"
	"bytes"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

type testStruct struct {
	Name      string
	BirthDate time.Time
	Emails    []string
	Balance   float64
}

type testGenericValues struct {
	name    string
	value   []any
	wantErr bool
}

type testGenericValue struct {
	name    string
	value   any
	wantErr bool
}

type testGenericDestValue struct {
	name    string
	value   any
	dest    any
	wantErr bool
}

type testIntMinMax struct {
	name  string
	value int
	min   int
	max   int
}

type testFile struct {
	name    string
	uri     string
	wantErr bool
}

type testFileValue struct {
	name    string
	value   *os.File
	wantErr bool
}

func initListTestConvertByteUnit() []testGenericValue {
	return []testGenericValue{
		{
			name:    "success B",
			value:   "1B",
			wantErr: false,
		},
		{
			name:    "success KB",
			value:   "1KB",
			wantErr: false,
		},
		{
			name:    "success MB",
			value:   "1MB",
			wantErr: false,
		},
		{
			name:    "success GB",
			value:   "1GB",
			wantErr: false,
		},
		{
			name:    "success TB",
			value:   "1TB",
			wantErr: false,
		},
		{
			name:    "success PB",
			value:   "1PB",
			wantErr: false,
		},
		{
			name:    "success EB",
			value:   "1EB",
			wantErr: false,
		},
		{
			name:    "success ZB",
			value:   "1ZB",
			wantErr: false,
		},
		{
			name:    "success YB",
			value:   "1YB",
			wantErr: false,
		},
		{
			name:    "failed length",
			value:   "1",
			wantErr: true,
		},
		{
			name:    "failed type",
			value:   "1AB",
			wantErr: true,
		},
		{
			name:    "failed empty",
			value:   "",
			wantErr: true,
		},
		{
			name:    "failed format",
			value:   "B1",
			wantErr: true,
		},
	}
}

func initListTestConvertMegaByteUnit() []testGenericValue {
	return []testGenericValue{
		{
			name:    "success MB",
			value:   "1MB",
			wantErr: false,
		},
		{
			name:    "success GB",
			value:   "1GB",
			wantErr: false,
		},
		{
			name:    "success TB",
			value:   "1TB",
			wantErr: false,
		},
		{
			name:    "success PB",
			value:   "1PB",
			wantErr: false,
		},
		{
			name:    "success EB",
			value:   "1EB",
			wantErr: false,
		},
		{
			name:    "success ZB",
			value:   "1ZB",
			wantErr: false,
		},
		{
			name:    "success YB",
			value:   "1YB",
			wantErr: false,
		},
		{
			name:    "failed length",
			value:   "1",
			wantErr: true,
		},
		{
			name:    "failed type",
			value:   "1AB",
			wantErr: true,
		},
		{
			name:    "failed empty",
			value:   "",
			wantErr: true,
		},
		{
			name:    "failed format",
			value:   "B1",
			wantErr: true,
		},
	}
}

func initListTestConvertToString() []testGenericValue {
	osFile, _ := os.Open("../gopher-helper.png")
	osFile2, _ := os.Open("../gopher-helper.png")
	bs, _ := io.ReadAll(osFile2)
	reader := bytes.NewReader(bs)
	buffer := bytes.NewBuffer(bs)
	strReader := strings.NewReader("test reader")
	readCloser := io.NopCloser(bytes.NewReader(bs))
	bufioReader := bufio.NewReader(bytes.NewReader(bs))
	ioReader := io.Reader(bytes.NewReader(bs))

	return []testGenericValue{
		{
			name:  "success string",
			value: "test",
		},
		{
			name:  "success struct",
			value: *initTestStruct(),
		},
		{
			name:  "success pointer struct",
			value: initTestStruct(),
		},
		{
			name:  "success slice",
			value: []int{12, 23, 45, 456, 467578},
		},
		{
			name:  "success map",
			value: *initTestMap(),
		},
		{
			name:  "success int",
			value: 21,
		},
		{
			name:  "success int8",
			value: int8(21),
		},
		{
			name:  "success int16",
			value: int16(21232),
		},
		{
			name:  "success int32",
			value: int32(2112312312),
		},
		{
			name:  "success int64",
			value: int64(9123809),
		},
		{
			name:  "success uint",
			value: uint(21),
		},
		{
			name:  "success uint8",
			value: uint8(21),
		},
		{
			name:  "success uint16",
			value: uint16(21232),
		},
		{
			name:  "success uint32",
			value: uint32(2112312312),
		},
		{
			name:  "success uint64",
			value: uint64(9123809),
		},
		{
			name:  "success float32",
			value: float32(12.23),
		},
		{
			name:  "success float64",
			value: 12.23,
		},
		{
			name:  "success bool",
			value: true,
		},
		{
			name:  "success bytes",
			value: []byte{116, 114, 117, 101},
		},
		{
			name:  "success file",
			value: osFile,
		},
		{
			name:  "success reader",
			value: reader,
		},
		{
			name:  "success string reader",
			value: strReader,
		},
		{
			name:  "success reader nop closer",
			value: readCloser,
		},
		{
			name:  "success bufio reader",
			value: bufioReader,
		},
		{
			name:  "success io reader",
			value: ioReader,
		},
		{
			name:  "success buffer",
			value: buffer,
		},
		{
			name:  "success time",
			value: time.Now(),
		},
		{
			name:  "success enum",
			value: exampleEnumStr1,
		},
		{
			name:    "failed",
			value:   initChan(),
			wantErr: true,
		},
		{
			name:    "failed invalid",
			wantErr: true,
		},
	}
}

func initListTestConvertToInt() []testGenericValue {
	return []testGenericValue{
		{
			name:    "failed string",
			value:   "test",
			wantErr: true,
		},
		{
			name:    "failed struct",
			value:   *initTestStruct(),
			wantErr: true,
		},
		{
			name:    "failed pointer struct",
			value:   initTestStruct(),
			wantErr: true,
		},
		{
			name:    "failed slice",
			value:   []int{12, 23, 45, 456, 467578},
			wantErr: true,
		},
		{
			name:    "failed map",
			value:   *initTestMap(),
			wantErr: true,
		},
		{
			name:  "success int",
			value: 21,
		},
		{
			name:  "success int8",
			value: int8(21),
		},
		{
			name:  "success int16",
			value: int16(21232),
		},
		{
			name:  "success int32",
			value: int32(2112312312),
		},
		{
			name:  "success int64",
			value: int64(9123809),
		},
		{
			name:  "success uint",
			value: uint(21),
		},
		{
			name:  "success uint8",
			value: uint8(21),
		},
		{
			name:  "success uint16",
			value: uint16(21232),
		},
		{
			name:  "success uint32",
			value: uint32(2112312312),
		},
		{
			name:  "success uint64",
			value: uint64(9123809),
		},
		{
			name:  "success float32",
			value: float32(12.23),
		},
		{
			name:  "success float64",
			value: 12.23,
		},
		{
			name:    "failed bool",
			value:   true,
			wantErr: true,
		},
		{
			name:    "failed bytes",
			value:   []byte{},
			wantErr: true,
		},
		{
			name:    "failed time",
			value:   time.Now(),
			wantErr: true,
		},
		{
			name:  "success string int",
			value: "234",
		},
		{
			name:  "success string float",
			value: "234.23",
		},
		{
			name:  "success enum",
			value: exampleEnumInt1,
		},
		{
			name:    "failed string bool",
			value:   "true",
			wantErr: true,
		},
		{
			name:    "failed",
			value:   initChan(),
			wantErr: true,
		},
		{
			name:    "failed invalid",
			wantErr: true,
		},
	}
}

func initListTestConvertToFloat() []testGenericValue {
	return []testGenericValue{
		{
			name:    "failed string",
			value:   "test",
			wantErr: true,
		},
		{
			name:    "failed struct",
			value:   *initTestStruct(),
			wantErr: true,
		},
		{
			name:    "failed pointer struct",
			value:   initTestStruct(),
			wantErr: true,
		},
		{
			name:    "failed slice",
			value:   []int{12, 23, 45, 456, 467578},
			wantErr: true,
		},
		{
			name:    "failed map",
			value:   *initTestMap(),
			wantErr: true,
		},
		{
			name:  "success int",
			value: 21,
		},
		{
			name:  "success int8",
			value: int8(21),
		},
		{
			name:  "success int16",
			value: int16(21232),
		},
		{
			name:  "success int32",
			value: int32(2112312312),
		},
		{
			name:  "success int64",
			value: int64(9123809),
		},
		{
			name:  "success uint",
			value: uint(21),
		},
		{
			name:  "success uint8",
			value: uint8(21),
		},
		{
			name:  "success uint16",
			value: uint16(21232),
		},
		{
			name:  "success uint32",
			value: uint32(2112312312),
		},
		{
			name:  "success uint64",
			value: uint64(9123809),
		},
		{
			name:  "success float32",
			value: float32(12.23),
		},
		{
			name:  "success float64",
			value: 12.23,
		},
		{
			name:    "failed bool",
			value:   true,
			wantErr: true,
		},
		{
			name:    "failed bytes",
			value:   []byte{},
			wantErr: true,
		},
		{
			name:    "failed time",
			value:   time.Now(),
			wantErr: true,
		},
		{
			name:  "success string int",
			value: "234",
		},
		{
			name:  "success string float",
			value: "234.23",
		},
		{
			name:  "success enum",
			value: exampleEnumFloat1,
		},
		{
			name:  "success enum int",
			value: exampleEnumInt1,
		},
		{
			name:    "failed string bool",
			value:   "true",
			wantErr: true,
		},
		{
			name:    "failed",
			value:   initChan(),
			wantErr: true,
		},
		{
			name:    "failed invalid",
			wantErr: true,
		},
	}
}

func initListTestConvertToBool() []testGenericValue {
	return []testGenericValue{
		{
			name:    "failed string",
			value:   "test",
			wantErr: true,
		},
		{
			name:    "failed struct",
			value:   *initTestStruct(),
			wantErr: true,
		},
		{
			name:    "failed pointer struct",
			value:   initTestStruct(),
			wantErr: true,
		},
		{
			name:    "failed slice",
			value:   []int{12, 23, 45, 456, 467578},
			wantErr: true,
		},
		{
			name:    "failed map",
			value:   *initTestMap(),
			wantErr: true,
		},
		{
			name:    "failed int",
			value:   21,
			wantErr: true,
		},
		{
			name:    "failed int8",
			value:   int8(21),
			wantErr: true,
		},
		{
			name:    "failed int16",
			value:   int16(21232),
			wantErr: true,
		},
		{
			name:    "failed int32",
			value:   int32(2112312312),
			wantErr: true,
		},
		{
			name:    "failed int64",
			value:   int64(9123809),
			wantErr: true,
		},
		{
			name:    "failed uint",
			value:   uint(21),
			wantErr: true,
		},
		{
			name:    "failed uint8",
			value:   uint8(21),
			wantErr: true,
		},
		{
			name:    "failed uint16",
			value:   uint16(21232),
			wantErr: true,
		},
		{
			name:    "failed uint32",
			value:   uint32(2112312312),
			wantErr: true,
		},
		{
			name:    "failed uint64",
			value:   uint64(9123809),
			wantErr: true,
		},
		{
			name:    "failed float32",
			value:   float32(12.23),
			wantErr: true,
		},
		{
			name:    "failed float64",
			value:   12.23,
			wantErr: true,
		},
		{
			name:  "success bool",
			value: true,
		},
		{
			name:  "success enum",
			value: exampleEnumBool1,
		},
		{
			name:    "failed bytes",
			value:   []byte{},
			wantErr: true,
		},
		{
			name:    "failed time",
			value:   time.Now(),
			wantErr: true,
		},
		{
			name:    "failed string int",
			value:   "234",
			wantErr: true,
		},
		{
			name:    "failed string float",
			value:   "234.23",
			wantErr: true,
		},
		{
			name:  "success string bool",
			value: "true",
		},
		{
			name:    "failed",
			value:   initChan(),
			wantErr: true,
		},
		{
			name:    "failed invalid",
			wantErr: true,
		},
	}
}

func initListTestConvertToTime() []testGenericValue {
	t := time.Now()
	return []testGenericValue{
		{
			name:    "string",
			value:   "test",
			wantErr: true,
		},
		{
			name:    "struct",
			value:   *initTestStruct(),
			wantErr: true,
		},
		{
			name:    "pointer struct",
			value:   initTestStruct(),
			wantErr: true,
		},
		{
			name:    "slice",
			value:   []int{12, 23, 45, 456, 467578},
			wantErr: true,
		},
		{
			name:    "map",
			value:   *initTestMap(),
			wantErr: true,
		},
		{
			name:  "int",
			value: 21129380912381,
		},
		{
			name:  "int8",
			value: int8(21),
		},
		{
			name:  "int16",
			value: int16(21233),
		},
		{
			name:  "int32",
			value: int32(2112312312),
		},
		{
			name:  "int64",
			value: int64(912380912312),
		},
		{
			name:  "uint",
			value: uint(21),
		},
		{
			name:  "uint8",
			value: uint8(21),
		},
		{
			name:  "uint16",
			value: uint16(21232),
		},
		{
			name:  "uint32",
			value: uint32(2112312312),
		},
		{
			name:  "uint64",
			value: uint64(9123809),
		},
		{
			name:  "float32",
			value: float32(122132134324534523.63),
		},
		{
			name:  "float64",
			value: 12123129318230234.1293821,
		},
		{
			name:  "time layout",
			value: t.Format(time.Layout),
		},
		{
			name:  "time ansic",
			value: t.Format(time.ANSIC),
		},
		{
			name:  "time unixDate",
			value: t.Format(time.UnixDate),
		},
		{
			name:  "time rubyDate",
			value: t.Format(time.RubyDate),
		},
		{
			name:  "time rfc822",
			value: t.Format(time.RFC822),
		},
		{
			name:  "time rfc822z",
			value: t.Format(time.RFC822Z),
		},
		{
			name:  "time rfc850",
			value: t.Format(time.RFC850),
		},
		{
			name:  "time rfc1123",
			value: t.Format(time.RFC1123),
		},
		{
			name:  "time rfc1123z",
			value: t.Format(time.RFC1123Z),
		},
		{
			name:  "time rfc3339",
			value: t.Format(time.RFC3339),
		},
		{
			name:  "time rfc3339nano",
			value: t.Format(time.RFC3339Nano),
		},
		{
			name:  "time kitchen",
			value: t.Format(time.Kitchen),
		},
		{
			name:  "time stamp",
			value: t.Format(time.Stamp),
		},
		{
			name:  "time stampMilli",
			value: t.Format(time.StampMilli),
		},
		{
			name:  "time stampMicro",
			value: t.Format(time.StampMicro),
		},
		{
			name:  "time stampNano",
			value: t.Format(time.StampNano),
		},
		{
			name:  "time dateTime",
			value: t.Format(time.DateTime),
		},
		{
			name:  "time dateOnly",
			value: t.Format(time.DateOnly),
		},
		{
			name:  "time timeOnly",
			value: t.Format(time.TimeOnly),
		},
		{
			name:  "time custom type",
			value: exampleTime(time.Now()),
		},
		{
			name:    "bool",
			value:   true,
			wantErr: true,
		},
		{
			name:    "failed nil",
			wantErr: true,
		},
	}
}

func initListTestConvertToDest() []testGenericDestValue {
	var s string
	var i int
	var f float64
	var b bool
	var st testStruct
	var m map[string]any
	var sl []any
	var tm time.Time
	var fl os.File
	var rd bytes.Reader
	var bff bytes.Buffer
	var objID primitive.ObjectID
	var dt primitive.DateTime

	var valueStr = "test value string"
	return []testGenericDestValue{
		{
			name:  "success string",
			value: &valueStr,
			dest:  &s,
		},
		{
			name:  "success struct",
			value: "{\"Name\":\"Foo Bar\",\"BirthDate\":\"2024-01-04T18:57:02.438015-03:00\",\"Emails\":[\"foobar@gmail.com\",\"foobar2@hotmail.com\"],\"Balance\":231.123}",
			dest:  &st,
		},
		{
			name:  "success slice",
			value: "[12, 23, 45, 456, 467578]",
			dest:  &sl,
		},
		{
			name:  "success map",
			value: "{\"float\":231.123,\"slice\":[\"foobar@gmail.com\",\"foobar2@hotmail.com\"],\"string\":\"test\",\"struct\":{\"Name\":\"Foo Bar\",\"BirthDate\":\"2024-01-04T18:57:02.438039-03:00\",\"Emails\":[\"foobar@gmail.com\",\"foobar2@hotmail.com\"],\"Balance\":231.123},\"time\":\"2024-01-04T18:57:02.438037-03:00\"}",
			dest:  &m,
		},
		{
			name:  "success int",
			value: "213",
			dest:  &i,
		},
		{
			name:  "success float",
			value: "12.23",
			dest:  &f,
		},
		{
			name:  "success bool",
			value: "true",
			dest:  &b,
		},
		{
			name:  "success time",
			value: "2024-01-04T19:37:29.087202-03:00",
			dest:  &tm,
		},
		{
			name:  "success file",
			value: "file",
			dest:  &fl,
		},
		{
			name:  "success reader",
			value: "reader",
			dest:  &rd,
		},
		{
			name:  "success buffet",
			value: "buffet",
			dest:  &bff,
		},
		{
			name:  "success objectId",
			value: primitive.NewObjectID().Hex(),
			dest:  &objID,
		},
		{
			name:  "success dateTime",
			value: primitive.NewDateTimeFromTime(time.Now()).Time(),
			dest:  &dt,
		},
		{
			name:    "failed",
			value:   initChan(),
			wantErr: true,
		},
		{
			name:    "failed dest invalid",
			value:   valueStr,
			dest:    initChanPointer(),
			wantErr: true,
		},
		{
			name:    "failed value nil",
			wantErr: true,
		},
		{
			name:    "failed not pointer",
			value:   "true",
			dest:    b,
			wantErr: true,
		},
		{
			name:    "failed json format",
			value:   "true 1123 test",
			dest:    &st,
			wantErr: true,
		},
		{
			name:    "failed error json format",
			value:   errors.New("error string value"),
			dest:    &st,
			wantErr: true,
		},
	}
}

func initListTestConvertFileToBytes() []testFileValue {
	f, _ := os.Open("../gopher-helper.png")
	return []testFileValue{
		{
			name:  "success",
			value: f,
		},
		{
			name:    "failed",
			wantErr: true,
		},
		{
			name:    "failed convert",
			wantErr: true,
		},
	}
}

func initListTestConvertToFile() []testGenericValue {
	f, _ := os.Open("../gopher-helper.png")
	return []testGenericValue{
		{
			name:  "success",
			value: f,
		},
		{
			name:    "failed",
			wantErr: true,
		},
		{
			name:    "failed convert",
			value:   initChan(),
			wantErr: true,
		},
	}
}

func initListTestIsNil() []testGenericValue {
	return []testGenericValue{
		{
			name:  "pointer",
			value: initTestStruct(),
		},
		{
			name:  "map",
			value: initTestMap(),
		},
		{
			name:  "slice",
			value: []int{23, 123},
		},
		{
			name:  "chan",
			value: initChanNil(),
		},
		{
			name:  "any",
			value: initEmptyAny(),
		},
	}
}

func initListTestIsNotNil() []testGenericValue {
	return []testGenericValue{
		{
			name:  "pointer",
			value: initPointerNil(),
		},
		{
			name:  "map",
			value: initMapNil(),
		},
		{
			name:  "slice",
			value: initSliceNil(),
		},
		{
			name:  "chan",
			value: initChan(),
		},
		{
			name:  "any",
			value: initEmptyAny(),
		},
	}
}

func initListTestIsEmpty() []testGenericValue {
	return []testGenericValue{
		{
			name:  "string",
			value: "",
		},
		{
			name:  "struct",
			value: testStruct{},
		},
		{
			name:  "pointer struct",
			value: &testStruct{},
		},
		{
			name:  "pointer",
			value: initPointerNil(),
		},
		{
			name:  "slice",
			value: []int{},
		},
		{
			name:  "pointer slice",
			value: &[]int{},
		},
		{
			name:  "map",
			value: map[string]any{},
		},
		{
			name:  "pointer map",
			value: &map[string]any{},
		},
		{
			name:  "int",
			value: 0,
		},
		{
			name:  "uint",
			value: uint(0),
		},
		{
			name:  "float",
			value: 0,
		},
		{
			name:  "bool",
			value: false,
		},
		{
			name:  "byte",
			value: []byte{},
		},
		{
			name:  "objectId",
			value: primitive.NilObjectID,
		},
		{
			name:  "primitive datetime",
			value: primitive.NewDateTimeFromTime(time.Time{}),
		},
	}
}

func initListTestIsNotEmpty() []testGenericValue {
	return []testGenericValue{
		{
			name:  "string",
			value: "test",
		},
		{
			name:  "struct",
			value: *initTestStruct(),
		},
		{
			name:  "pointer struct",
			value: initTestStruct(),
		},
		{
			name:  "slice",
			value: []int{23, 123, 23134, 54345},
		},
		{
			name:  "pointer slice",
			value: &[]int{123, 23, 4343},
		},
		{
			name:  "map",
			value: *initTestMap(),
		},
		{
			name:  "pointer map",
			value: initTestMap(),
		},
		{
			name:  "int",
			value: 23,
		},
		{
			name:  "uint",
			value: uint(2432),
		},
		{
			name:  "float",
			value: 12.23,
		},
		{
			name:  "bool",
			value: true,
		},
		{
			name:  "byte",
			value: []byte{101, 133, 178, 255, 197, 19, 21, 77, 11, 38, 26, 152},
		},
		{
			name:  "objectId",
			value: primitive.NewObjectID(),
		},
		{
			name:  "primitive datetime",
			value: primitive.NewDateTimeFromTime(time.Now()),
		},
	}
}

func initListTestIsNumeric() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success",
			value: "123",
		},
		{
			name:  "failed",
			value: true,
		},
	}
}

func initListTestIsLetter() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success",
			value: "letter",
		},
		{
			name:  "failed",
			value: 123,
		},
	}
}

func initListTestIsStringJson() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success",
			value: "{\"Name\":\"Foo Bar\",\"BirthDate\":\"2024-01-04T20:00:31.570313-03:00\",\"Emails\":[\"foobar@gmail.com\",\"foobar2@hotmail.com\"],\"Balance\":231.123}",
		},
		{
			name:  "failed",
			value: 123,
		},
	}
}

func initListTestIsStringInt() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success",
			value: "123",
		},
		{
			name:  "failed",
			value: 123,
		},
	}
}

func initListTestIsStringBool() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success",
			value: "false",
		},
		{
			name:  "failed",
			value: 123,
		},
	}
}

func initListTestIsStringFloat() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success",
			value: "12.32",
		},
		{
			name:  "failed",
			value: 123,
		},
	}
}

func initListTestIsStringTime() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success",
			value: time.Now().Format(time.RFC1123Z),
		},
		{
			name:  "failed",
			value: 123,
		},
	}
}

func initListTestEquals() []testGenericValues {
	mapTest := initTestMap()
	structTest := initTestStruct()
	return []testGenericValues{
		{
			name:  "string",
			value: []any{"test", "test"},
		},
		{
			name:  "string diff cases",
			value: []any{"testaT", "testAt"},
		},
		{
			name:  "struct",
			value: []any{structTest, structTest},
		},
		{
			name:  "map",
			value: []any{mapTest, mapTest},
		},
		{
			name:  "int",
			value: []any{12, 12},
		},
		{
			name:  "not equals",
			value: []any{"test", 12},
		},
		{
			name:  "bool",
			value: []any{false, false},
		},
		{
			name:  "bool",
			value: []any{true, true},
		},
	}
}

func initListTestGreaterThan() []testGenericValues {
	mapTest := initTestMap()
	structTest := initTestStruct()
	return []testGenericValues{
		{
			name:  "string",
			value: []any{"tests", "test"},
		},
		{
			name:  "struct",
			value: []any{structTest, structTest},
		},
		{
			name:  "map",
			value: []any{mapTest, mapTest},
		},
		{
			name:  "int",
			value: []any{13, 12},
		},
		{
			name:  "float",
			value: []any{14.2, 13.2},
		},
		{
			name:  "equal",
			value: []any{14.2, 14.2},
		},
		{
			name:  "false",
			value: []any{14.2, 15.2},
		},
	}
}

func initListTestLessThan() []testGenericValues {
	mapTest := initTestMap()
	structTest := initTestStruct()
	return []testGenericValues{
		{
			name:  "string",
			value: []any{"test", "tests"},
		},
		{
			name:  "struct",
			value: []any{structTest, structTest},
		},
		{
			name:  "map",
			value: []any{mapTest, mapTest},
		},
		{
			name:  "int",
			value: []any{12, 13},
		},
		{
			name:  "float",
			value: []any{13.2, 14.2},
		},
		{
			name:  "equal",
			value: []any{14.2, 14.2},
		},
		{
			name:  "false",
			value: []any{15.2, 14.2},
		},
	}
}

func initListTestIntMin() []testIntMinMax {
	return []testIntMinMax{
		{
			name:  "value",
			min:   0,
			value: 1,
		},
		{
			name:  "min",
			min:   0,
			value: -1,
		},
	}
}

func initListTestIntMax() []testIntMinMax {
	return []testIntMinMax{
		{
			name:  "value",
			max:   3,
			value: 2,
		},
		{
			name:  "max",
			max:   3,
			value: 4,
		},
	}
}

func initListTestGetFirstLastName() []testGenericValue {
	return []testGenericValue{
		{
			name:  "full",
			value: "Gabriel Cataldo",
		},
		{
			name:  "partial",
			value: "Gabriel",
		},
		{
			name:  "empty",
			value: "",
		},
	}
}

func initListTestTime() []testGenericValues {
	return []testGenericValues{
		{
			name: "after",
			value: []any{
				time.Date(2030, 12, 1, 0, 0, 0, 0, time.Local),
				time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local),
			},
		},
		{
			name: "before",
			value: []any{
				time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local),
				time.Date(2030, 12, 1, 0, 0, 0, 0, time.Local),
			},
		},
		{
			name: "now",
			value: []any{
				time.Now(),
				time.Now(),
			},
		},
		{
			name: "today",
			value: []any{
				time.Now(),
				time.Now(),
			},
		},
		{
			name: "failed value type a",
			value: []any{
				"string value",
				time.Now(),
			},
			wantErr: true,
		},
		{
			name: "failed value type b",
			value: []any{
				time.Now(),
				"string value",
			},
			wantErr: true,
		},
	}
}

func initListTestGetFileString() []testFile {
	return []testFile{
		{
			name: "success",
			uri:  "../postal-codes.json",
		},
		{
			name:    "failed",
			uri:     "",
			wantErr: true,
		},
	}
}

func initListTestGetFileJson() []testFile {
	return []testFile{
		{
			name: "success",
			uri:  "../postal-codes.json",
		},
		{
			name:    "failed",
			uri:     "",
			wantErr: true,
		},
	}
}

func initTestStruct() *testStruct {
	return &testStruct{
		Name:      "Foo Bar",
		BirthDate: time.Now(),
		Emails:    []string{"foobar@gmail.com", "foobar2@hotmail.com"},
		Balance:   231.123,
	}
}

func initTestMap() *map[string]any {
	now := time.Now()
	return &map[string]any{
		"string": "test",
		"time":   now,
		"slice":  []string{"foobar@gmail.com", "foobar2@hotmail.com"},
		"float":  231.123,
		"struct": initTestStruct(),
	}
}

func initPointerNil() *any {
	return nil
}

func initMapNil() map[string]any {
	return nil
}

func initSliceNil() []any {
	return nil
}

func initChanNil() chan string {
	return nil
}

func initChan() chan string {
	return make(chan string, 1)
}

func initChanPointer() *chan string {
	c := make(chan string, 1)
	return &c
}

func initEmptyAny() any {
	var a any
	return a
}

func panicRecovery(t *testing.T, wantErr bool) {
	if r := recover(); (r != nil) != wantErr {
		t.Errorf("panic() wantErr=%v", wantErr)
	}
}
