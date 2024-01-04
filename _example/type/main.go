package main

import (
	"github.com/GabrielHCataldo/go-helper/helper"
	"github.com/GabrielHCataldo/go-logger/logger"
	"time"
)

func main() {
	var a any = map[string]any{}

	isMap := helper.IsMap(a)
	logger.Info("is map?", isMap)

	a = "value1"

	isString := helper.IsString(a)
	logger.Info("is string?", isString)

	a = 12

	isInt := helper.IsInt(a)
	logger.Info("is int?", isInt)

	a = true

	isBool := helper.IsBool(a)
	logger.Info("is bool?", isBool)

	a = time.Now()

	isTime := helper.IsTime(a)
	logger.Info("is time?", isTime)
}
