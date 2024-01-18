package main

import (
	"github.com/GabrielHCataldo/go-helper/helper"
	"log"
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
		log.Println("error convert base64 to struct:", err)
		return
	}
	log.Println("converted base64 to struct:", destStruct)
	var destMap map[string]any
	// convert struct to simple map
	err = helper.ConvertToDest(destStruct, &destMap)
	if helper.IsNotNil(err) {
		log.Println("error convert struct to map:", err)
		return
	}
	log.Println("converted struct to map:", destMap)
}

func initExampleStruct() exampleStruct {
	return exampleStruct{
		Name:      "Foo Bar",
		BirthDate: time.Now(),
		Map:       map[string]any{"example": 1, "test": 2},
	}
}
