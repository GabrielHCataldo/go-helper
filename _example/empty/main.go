package main

import (
	"github.com/GabrielHCataldo/go-helper/helper"
	"github.com/GabrielHCataldo/go-logger/logger"
	"time"
)

func main() {
	allNilAndAllNotNil()
	allEmptyAndAllNotEmpty()
	pointerEmptyAndNotEmpty()
	mapEmptyAndNotEmpty()
	sliceEmptyAndNotEmpty()
	stringEmptyAndNotEmpty()
	structEmptyAndNotEmpty()
	pointerNilAndNotNil()
}

func allNilAndAllNotNil() {
	var anyPointer *any
	var nMap map[string]any
	var nSlice []any
	var nChan chan struct{}
	var nInterface interface{}

	allNil := helper.AllNil(anyPointer, nMap, nSlice, nChan, nInterface)
	logger.Info("all are nil?", allNil)

	anyPointer = helper.ConvertToPointer(any("value string"))
	nMap = map[string]any{}
	nSlice = append(nSlice, any("value"))
	nChan = make(chan struct{}, 1)
	nInterface = "value"

	allNotNil := helper.AllNotNil(anyPointer, nMap, nSlice, nChan, nInterface)
	logger.Info("all are not nil?", allNotNil)
}

func allEmptyAndAllNotEmpty() {
	var anyPointer *any
	var nMap map[string]any
	var nSlice []any
	var nString string
	var nInt int
	var nFloat float64
	var nBool bool
	allEmpty := helper.AllEmpty(anyPointer, nMap, nSlice, nString, nInt, nFloat, nBool)
	logger.Info("all are empty?", allEmpty)
	anyPointer = helper.ConvertToPointer(any("value string"))
	nMap = map[string]any{
		"test": "value",
	}
	nString = "value"
	nInt = 1
	nFloat = 1.0
	nBool = true
	nSlice = append(nSlice, any("value"))
	allNotEmpty := helper.AllNotEmpty(anyPointer, nMap, nString, nInt, nFloat, nBool)
	logger.Info("all are not empty?", allNotEmpty)
}

func pointerNilAndNotNil() {
	var anyPointer *any
	isNil := helper.IsNil(anyPointer)
	logger.Info("pointer is nil?", isNil)
	anyPointer = helper.ConvertToPointer(any("value string"))
	isNotNil := helper.IsNotNil(anyPointer)
	logger.Info("pointer is not nil?", isNotNil)
}

func pointerEmptyAndNotEmpty() {
	var anyPointer *any
	isEmpty := helper.IsEmpty(anyPointer)
	logger.Info("pointer is empty?", isEmpty)
	anyPointer = helper.ConvertToPointer(any("value string"))
	isNotEmpty := helper.IsNotEmpty(anyPointer)
	logger.Info("pointer is not empty?", isNotEmpty)
}

func mapEmptyAndNotEmpty() {
	nMap := map[string]any{}
	isEmpty := helper.IsEmpty(&nMap)
	logger.Info("pointer map is empty?", isEmpty)
	nMap["test"] = "value string"
	isNotEmpty := helper.IsNotEmpty(nMap)
	logger.Info("pointer map is not empty?", isNotEmpty)
}

func sliceEmptyAndNotEmpty() {
	var nSlice []any
	isEmpty := helper.IsEmpty(nSlice)
	logger.Info("slice is empty?", isEmpty)
	nSlice = append(nSlice, "test string not empty")
	isNotEmpty := helper.IsNotEmpty(nSlice)
	logger.Info("slice is not empty?", isNotEmpty)
}

func stringEmptyAndNotEmpty() {
	nBlankString := "  "
	isEmpty := helper.IsEmpty(nBlankString)
	logger.Info("string is empty?", isEmpty)
	nBlankString = "test name"
	isNotEmpty := helper.IsNotEmpty(nBlankString)
	logger.Info("string is not empty?", isNotEmpty)
}

type exampleStruct struct {
	Name      string
	BirthDate time.Time
	Map       map[string]any
}

func structEmptyAndNotEmpty() {
	nStruct := exampleStruct{}
	isEmpty := helper.IsEmpty(nStruct)
	logger.Info("struct is empty?", isEmpty)
	nStruct.Name = "test name"
	isNotEmpty := helper.IsNotEmpty(nStruct)
	logger.Info("struct is not empty?", isNotEmpty)
}
