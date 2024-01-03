package main

import (
	"github.com/GabrielHCataldo/go-helper/helper"
	"github.com/GabrielHCataldo/go-logger/logger"
)

func main() {
	randomResult := helper.RandomBool()
	logger.Info("result randomBool:", randomResult)
}
