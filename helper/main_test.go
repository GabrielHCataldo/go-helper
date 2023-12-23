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

type testConvertByteUnit struct {
	name    string
	value   string
	wantErr bool
}

type testConvertMegaByteUnit struct {
	name    string
	value   string
	wantErr bool
}

type testConvertToString struct {
	name      string
	value     any
	wantEmpty bool
}

type testValueGeneric struct {
	name    string
	value   any
	wantErr bool
}

func initListTestConvertByteUnit() []testConvertByteUnit {
	return []testConvertByteUnit{
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
	}
}

func initListTestConvertMegaByteUnit() []testConvertMegaByteUnit {
	return []testConvertMegaByteUnit{
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
	}
}

func initListTestConvertToString() []testConvertToString {
	return []testConvertToString{
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
			name:  "success byte",
			value: []byte{101, 133, 178, 255, 197, 19, 21, 77, 11, 38, 26, 152},
		},
	}
}

func initListTestIsNil() []testValueGeneric {
	return []testValueGeneric{
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
	}
}

func initListTestIsNonNil() []testValueGeneric {
	return []testValueGeneric{
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
	}
}

func initListTestIsEmpty() []testValueGeneric {
	return []testValueGeneric{
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

func initListTestIsNotEmpty() []testValueGeneric {
	return []testValueGeneric{
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

func initListTestIsPointerNil() []testValueGeneric {
	return []testValueGeneric{
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

func initListTestIsPointerNonNil() []testValueGeneric {
	return []testValueGeneric{
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

func initListTestIsJsonEmpty() []testValueGeneric {
	return []testValueGeneric{
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

func initListTestIsJsonNotEmpty() []testValueGeneric {
	return []testValueGeneric{
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

func initListTestIsMapEmpty() []testValueGeneric {
	return []testValueGeneric{
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

func initListTestIsMapNotEmpty() []testValueGeneric {
	return []testValueGeneric{
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

func initListTestIsStructEmpty() []testValueGeneric {
	return []testValueGeneric{
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

func initListTestIsStructNotEmpty() []testValueGeneric {
	return []testValueGeneric{
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

func initTestStruct() *testStruct {
	return &testStruct{
		Name:      "Foo Bar",
		BirthDate: time.Now(),
		Emails:    []string{"foobar@gmail.com", "foobar2@hotmail.com"},
		Balance:   231.123,
	}
}

func initTestMap() *map[string]any {
	return &map[string]any{
		"string": "test",
		"time":   time.Now(),
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

func panicRecovery(t *testing.T, wantErr bool) {
	if r := recover(); (r != nil) != wantErr {
		t.Errorf("panic() wantErr=%v", wantErr)
	}
}
