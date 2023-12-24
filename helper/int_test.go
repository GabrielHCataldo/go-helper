package helper

import (
	"github.com/GabrielHCataldo/go-logger/logger"
	"testing"
)

func TestMinInt(t *testing.T) {
	for _, tt := range initListTestIntMin() {
		t.Run(tt.name, func(t *testing.T) {
			result := MinInt(tt.value, tt.min)
			logger.Info("MinInt:", result)
		})
	}
}

func TestMinInt32(t *testing.T) {
	for _, tt := range initListTestIntMin() {
		t.Run(tt.name, func(t *testing.T) {
			result := MinInt32(int32(tt.value), int32(tt.min))
			logger.Info("MinInt32:", result)
		})
	}
}

func TestMinInt64(t *testing.T) {
	for _, tt := range initListTestIntMin() {
		t.Run(tt.name, func(t *testing.T) {
			result := MinInt64(int64(tt.value), int64(tt.min))
			logger.Info("MinInt64:", result)
		})
	}
}

func TestMaxInt(t *testing.T) {
	for _, tt := range initListTestIntMax() {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxInt(tt.value, tt.max)
			logger.Info("MaxInt:", result)
		})
	}
}

func TestMaxInt32(t *testing.T) {
	for _, tt := range initListTestIntMax() {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxInt32(int32(tt.value), int32(tt.max))
			logger.Info("MaxInt32:", result)
		})
	}
}

func TestMaxInt64(t *testing.T) {
	for _, tt := range initListTestIntMax() {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxInt64(int64(tt.value), int64(tt.max))
			logger.Info("MaxInt64:", result)
		})
	}
}

func TestRandomNumber(t *testing.T) {
	result := RandomNumber(0, 100)
	logger.Info("RandomNumber:", result)
}
