package main

import (
	"github.com/GabrielHCataldo/go-helper/helper"
	"log"
	"time"
)

func main() {
	var a any = map[string]any{}

	isMap := helper.IsMapType(a)
	log.Println("is map?", isMap)

	a = "value1"

	isString := helper.IsStringType(a)
	log.Println("is string?", isString)

	a = 12

	isInt := helper.IsIntType(a)
	log.Println("is int?", isInt)

	a = true

	isBool := helper.IsBoolType(a)
	log.Println("is bool?", isBool)

	a = time.Now()

	isTime := helper.IsTimeType(a)
	log.Println("is time?", isTime)
}
