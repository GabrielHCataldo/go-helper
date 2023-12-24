package helper

import (
	"github.com/GabrielHCataldo/go-logger/logger"
	"testing"
)

func TestRound(t *testing.T) {
	v := 1.6
	result := Round(v, 2)
	logger.Info("Round:", result)
}

func TestRoundUp(t *testing.T) {
	v := 1.5
	result := RoundUp(v, 2)
	logger.Info("RoundUp:", result)
}

func TestRoundDown(t *testing.T) {
	v := 1.5
	result := RoundDown(v, 2)
	logger.Info("RoundDown:", result)
}
