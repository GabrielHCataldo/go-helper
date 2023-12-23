package helper

import (
	"github.com/GabrielHCataldo/go-logger/logger"
	"testing"
)

func TestRandomBool(t *testing.T) {
	b := RandomBool()
	logger.Info("result:", b)
}
