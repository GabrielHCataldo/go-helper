package main

import (
	"github.com/GabrielHCataldo/go-helper/helper"
	"github.com/GabrielHCataldo/go-logger/logger"
	"time"
)

type exampleStruct struct {
	Name      string
	BirthDate time.Time
	Map       map[string]any
}

func main() {
	emptyAndNotEmpty()
	nilAndNotNil()
}

func nilAndNotNil() {
	var anyPointer *any
	var anotherValue []any
	isNil := helper.IsNil(anyPointer, anotherValue)
	logger.Info("is nil?", isNil)
	anyPointer = helper.ConvertToPointer(any("value string"))
	anotherValue = []any{12, "test"}
	isNotNil := helper.IsNotNil(anyPointer, anotherValue)
	logger.Info("is not nil?", isNotNil)
}

func emptyAndNotEmpty() {
	var anyPointer *any
	nStruct := exampleStruct{}
	nBlankString := "  "
	nMap := map[string]any{}
	var nSlice []any
	isEmpty := helper.IsEmpty(anyPointer, nStruct, nBlankString, nMap, nSlice)
	logger.Info("is empty?", isEmpty)
	anyPointer = helper.ConvertToPointer(any("value string"))
	nBlankString = "test"
	nStruct.Name = "test name"
	nMap["test"] = "value string"
	nSlice = append(nSlice, "test string not empty")
	isNotEmpty := helper.IsNotEmpty(anyPointer, nStruct, nBlankString, nMap, nSlice)
	logger.Info("is not empty?", isNotEmpty)
}
