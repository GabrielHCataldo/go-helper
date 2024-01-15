package main

import (
	"github.com/GabrielHCataldo/go-helper/helper"
	"github.com/GabrielHCataldo/go-logger/logger"
)

func main() {
	s1 := "value"
	s2 := "value"
	s3 := "value"
	s4 := "value"

	equals := helper.IsEqual(s1, s2, s3, s4)
	logger.Info("equals?", equals)

	s1 = "value1"

	notEquals := helper.IsNotEqual(s1, s2, s3, s4)
	logger.Info("not equals?", notEquals)
}
