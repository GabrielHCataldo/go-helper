package helper

import (
	"github.com/GabrielHCataldo/go-logger/logger"
	"testing"
)

func TestIsTrue(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsTrue(tt.value...)
			logger.Info("IsTrue:", empty)
		})
	}
}

func TestIsNotTrue(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsNotTrue(tt.value...)
			logger.Info("IsNotTrue:", empty)
		})
	}
}

func TestIsEqual(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsEqual(tt.value...)
			logger.Info("IsEqual:", empty)
		})
	}
}

func TestIsNotEqual(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsNotEqual(tt.value...)
			logger.Info("IsNotEqual:", empty)
		})
	}
}

func TestIsEqualIgnoreCase(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsEqualIgnoreCase(tt.value...)
			logger.Info("IsEqualIgnoreCase:", empty)
		})
	}
}

func TestIsNotEqualIgnoreCase(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsNotEqualIgnoreCase(tt.value...)
			logger.Info("IsNotEqualIgnoreCase:", empty)
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
