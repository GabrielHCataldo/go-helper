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

func TestIsGreater(t *testing.T) {
	for _, tt := range initListTestGreater() {
		t.Run(tt.name, func(t *testing.T) {
			var a any = 2
			values := tt.value
			if IsNotEmpty(tt.value) {
				a = tt.value[0]
				values = tt.value[1:]
			}
			empty := IsGreater(a, values...)
			logger.Info("IsGreater:", empty)
		})
	}
}

func TestIsGreaterEqual(t *testing.T) {
	for _, tt := range initListTestGreater() {
		t.Run(tt.name, func(t *testing.T) {
			var a any = 2
			values := tt.value
			if IsNotEmpty(tt.value) {
				a = tt.value[0]
				values = tt.value[1:]
			}
			empty := IsGreaterEqual(a, values...)
			logger.Info("IsGreaterEqual:", empty)
		})
	}
}

func TestIsSmaller(t *testing.T) {
	for _, tt := range initListTestSmaller() {
		t.Run(tt.name, func(t *testing.T) {
			var a any = 2
			values := tt.value
			if IsNotEmpty(tt.value) {
				a = tt.value[0]
				values = tt.value[1:]
			}
			empty := IsSmaller(a, values...)
			logger.Info("IsSmaller:", empty)
		})
	}
}

func TestIsSmallerEqual(t *testing.T) {
	for _, tt := range initListTestSmaller() {
		t.Run(tt.name, func(t *testing.T) {
			var a any = 2
			values := tt.value
			if IsNotEmpty(tt.value) {
				a = tt.value[0]
				values = tt.value[1:]
			}
			empty := IsSmallerEqual(a, values...)
			logger.Info("IsSmallerEqual:", empty)
		})
	}
}
