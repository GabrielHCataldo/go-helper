package helper

import (
	"github.com/GabrielHCataldo/go-logger/logger"
	"testing"
)

func TestEquals(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := Equals(tt.value...)
			logger.Info("Equals:", empty)
		})
	}
}

func TestNotEquals(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := NotEquals(tt.value...)
			logger.Info("NotEquals:", empty)
		})
	}
}
