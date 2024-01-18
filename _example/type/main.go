package main

import (
	"github.com/GabrielHCataldo/go-helper/helper"
	"log"
	"time"
)

func main() {
	var a any = map[string]any{}

	isMap := helper.IsMap(a)
	log.Println("is map?", isMap)

	a = "value1"

	isString := helper.IsString(a)
	log.Println("is string?", isString)

	a = 12

	isInt := helper.IsInt(a)
	log.Println("is int?", isInt)

	a = true

	isBool := helper.IsBool(a)
	log.Println("is bool?", isBool)

	a = time.Now()

	isTime := helper.IsTime(a)
	log.Println("is time?", isTime)
}
