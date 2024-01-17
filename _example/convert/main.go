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
	playground()
}

func playground() {
	var destStruct exampleStruct
	// convert to struct to string base64, it's funny and sample
	s64 := helper.SimpleConvertToBase64(initExampleStruct())
	// convert string base64 to struct
	err := helper.ConvertToDest(s64, &destStruct)
	if helper.IsNotNil(err) {
		logger.Error("error convert base64 to struct:", err)
		return
	}
	logger.Info("converted base64 to struct:", destStruct)
	var destMap map[string]any
	// convert struct to simple map
	err = helper.ConvertToDest(destStruct, &destMap)
	if helper.IsNotNil(err) {
		logger.Error("error convert struct to map:", err)
		return
	}
	logger.Info("converted struct to map:", destMap)
}

func initExampleStruct() exampleStruct {
	return exampleStruct{
		Name:      "Foo Bar",
		BirthDate: time.Now(),
		Map:       map[string]any{"example": 1, "test": 2},
	}
}
