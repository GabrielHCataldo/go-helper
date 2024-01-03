package main

import (
	"github.com/GabrielHCataldo/go-helper/helper"
	"github.com/GabrielHCataldo/go-logger/logger"
	"time"
)

func main() {
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
