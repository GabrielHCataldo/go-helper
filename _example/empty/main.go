package main

import (
	"github.com/GabrielHCataldo/go-helper/helper"
	"github.com/GabrielHCataldo/go-logger/logger"
	"time"
)

func main() {
	nMap := map[string]any{}
	isEmpty := helper.IsEmpty(&nMap)
	logger.Info("pointer map is empty?", isEmpty)
	nMap["test"] = "value string"
	isNotEmpty := helper.IsNotEmpty(nMap)
	logger.Info("pointer map is not empty?", isNotEmpty)
}

func sliceEmptyAndNotEmpty() {
	nSlice := []any{}
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
