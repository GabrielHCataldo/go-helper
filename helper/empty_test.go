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
	for _, tt := range initListTestIsNonNil() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsNonNil(tt.value)
			logger.Info("IsNonNil:", empty)
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

func TestIsPointerNil(t *testing.T) {
	for _, tt := range initListTestIsPointerNil() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			isNil := IsPointerNil(tt.value)
			logger.Info("IsPointerNil:", isNil)
		})
	}
}

func TestIsPointerNonNil(t *testing.T) {
	for _, tt := range initListTestIsPointerNonNil() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			isNonNil := IsPointerNonNil(tt.value)
			logger.Info("IsPointerNonNil:", isNonNil)
		})
	}
}

func TestIsJsonEmpty(t *testing.T) {
	for _, tt := range initListTestIsJsonEmpty() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			isEmpty := IsJsonEmpty(tt.value)
			logger.Info("IsJsonEmpty:", isEmpty)
		})
	}
}

func TestIsJsonNotEmpty(t *testing.T) {
	for _, tt := range initListTestIsJsonNotEmpty() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			isNotEmpty := IsJsonNotEmpty(tt.value)
			logger.Info("IsJsonNotEmpty:", isNotEmpty)
		})
	}
}

func TestIsMapEmpty(t *testing.T) {
	for _, tt := range initListTestIsMapEmpty() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			isEmpty := IsMapEmpty(tt.value)
			logger.Info("IsMapEmpty:", isEmpty)
		})
	}
}

func TestIsMapNotEmpty(t *testing.T) {
	for _, tt := range initListTestIsMapNotEmpty() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			isNotEmpty := IsMapNotEmpty(tt.value)
			logger.Info("IsMapNotEmpty:", isNotEmpty)
		})
	}
}

func TestIsStructEmpty(t *testing.T) {
	for _, tt := range initListTestIsStructEmpty() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			isEmpty := IsStructEmpty(tt.value)
			logger.Info("IsStructEmpty:", isEmpty)
		})
	}
}

func TestIsStructNotEmpty(t *testing.T) {
	for _, tt := range initListTestIsStructNotEmpty() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			isNotEmpty := IsStructNotEmpty(tt.value)
			logger.Info("IsStructNotEmpty:", isNotEmpty)
		})
	}
}
