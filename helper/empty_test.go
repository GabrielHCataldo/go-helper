package helper

import (
	"github.com/GabrielHCataldo/go-logger/logger"
	"testing"
)

func TestIsNil(t *testing.T) {
	for _, tt := range initListTestIsNil() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsNil(tt.value)
			logger.Info("IsNil:", empty)
		})
	}
}

func TestIsNotNil(t *testing.T) {
	for _, tt := range initListTestIsNotNil() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsNotNil(tt.value)
			logger.Info("IsNotNil:", empty)
		})
	}
}

func TestIsEmpty(t *testing.T) {
	for _, tt := range initListTestIsEmpty() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsEmpty(tt.value)
			logger.Info("IsEmpty:", empty)
		})
	}
}

func TestIsNotEmpty(t *testing.T) {
	for _, tt := range initListTestIsNotEmpty() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsNotEmpty(tt.value)
			logger.Info("IsNotEmpty:", empty)
		})
	}
}
