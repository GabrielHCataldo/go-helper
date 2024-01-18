package helper

import (
	"github.com/GabrielHCataldo/go-logger/logger"
	"testing"
)

func TestGetCallerInfo(t *testing.T) {
	name, line, funcName := GetCallerInfo(1)
	logger.Info(name+":"+line, "funcName:", funcName)
}
