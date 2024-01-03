package helper

import (
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
			name:    "success chan",
			value:   initChan(),
			wantErr: true,
		},
	}
}

func initListTestConvertToTime() []testGenericValue {
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
			name:    "bool",
			value:   true,
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

func initListTestIsNonNil() []testGenericValue {
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
	}
}

func initListTestIsPointerNil() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success",
			value: initPointerNil(),
		},
		{
			name:    "failed",
			value:   "",
			wantErr: true,
		},
	}
}

func initListTestIsPointerNonNil() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success",
			value: initTestStruct(),
		},
		{
			name:    "failed",
			value:   "",
			wantErr: true,
		},
	}
}

func initListTestIsJsonEmpty() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success struct",
			value: testStruct{},
		},
		{
			name:  "success map",
			value: map[string]any{},
		},
		{
			name:    "failed",
			value:   "",
			wantErr: true,
		},
	}
}

func initListTestIsJsonNotEmpty() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success struct",
			value: *initTestStruct(),
		},
		{
			name:  "success map",
			value: *initTestMap(),
		},
		{
			name:    "failed",
			value:   "",
			wantErr: true,
		},
	}
}

func initListTestIsMapEmpty() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success",
			value: map[string]any{},
		},
		{
			name:    "failed",
			value:   "",
			wantErr: true,
		},
	}
}

func initListTestIsMapNotEmpty() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success",
			value: *initTestMap(),
		},
		{
			name:    "failed",
			value:   "",
			wantErr: true,
		},
	}
}

func initListTestIsStructEmpty() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success",
			value: testStruct{},
		},
		{
			name:    "failed",
			value:   "",
			wantErr: true,
		},
	}
}

func initListTestIsStructNotEmpty() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success",
			value: *initTestStruct(),
		},
		{
			name:    "failed",
			value:   "",
			wantErr: true,
		},
	}
}

func initListTestIsSliceEmpty() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success",
			value: []any{},
		},
		{
			name:    "failed",
			value:   "",
			wantErr: true,
		},
	}
}

func initListTestIsSliceNotEmpty() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success",
			value: []any{"test", 23, 123},
		},
		{
			name:    "failed",
			value:   "",
			wantErr: true,
		},
	}
}

func initListTestIsStringEmpty() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success",
			value: *initEmptyPointerString(),
		},
		{
			name:  "success pointer empty",
			value: initEmptyPointerString(),
		},
		{
			name:    "failed",
			value:   23,
			wantErr: true,
		},
	}
}

func initListTestIsStringNotEmpty() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success",
			value: "not empty",
		},
		{
			name:    "failed",
			value:   23,
			wantErr: true,
		},
	}
}

func initListTestIsIntEmpty() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success",
			value: 0,
		},
		{
			name:    "failed",
			value:   "",
			wantErr: true,
		},
	}
}

func initListTestIsIntNotEmpty() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success",
			value: 23,
		},
		{
			name:    "failed",
			value:   "",
			wantErr: true,
		},
	}
}

func initListTestIsFloatEmpty() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success",
			value: 0.0,
		},
		{
			name:    "failed",
			value:   "",
			wantErr: true,
		},
	}
}

func initListTestIsFloatNotEmpty() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success",
			value: 23.23,
		},
		{
			name:    "failed",
			value:   "",
			wantErr: true,
		},
	}
}

func initListTestIsBoolEmpty() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success",
			value: false,
		},
		{
			name:    "failed",
			value:   "",
			wantErr: true,
		},
	}
}

func initListTestIsBoolNotEmpty() []testGenericValue {
	return []testGenericValue{
		{
			name:  "success",
			value: true,
		},
		{
			name:    "failed",
			value:   "",
			wantErr: true,
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

func initListTestEquals() []testGenericValues {
	mapTest := initTestMap()
	structTest := initTestStruct()
	return []testGenericValues{
		{
			name:  "string",
			value: []any{"test", "test"},
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
			value: []any{"test", "test", 12},
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
		{
			name:    "failed",
			value:   true,
			wantErr: true,
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

func initEmptyAny() any {
	var a any
	return a
}
func initEmptyPointerString() *string {
	var a string
	return &a
}

func panicRecovery(t *testing.T, wantErr bool) {
	if r := recover(); (r != nil) != wantErr {
		t.Errorf("panic() wantErr=%v", wantErr)
	}
}
