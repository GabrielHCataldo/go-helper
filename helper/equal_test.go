package helper

import (
	"github.com/GabrielHCataldo/go-logger/logger"
	"testing"
)

func TestEqual(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := Equal(tt.value...)
			logger.Info("Equal:", empty)
		})
	}
}

func TestNotEqual(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := NotEqual(tt.value...)
			logger.Info("NotEqual:", empty)
		})
	}
}
