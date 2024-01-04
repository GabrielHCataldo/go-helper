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

func TestAllNil(t *testing.T) {
	for _, tt := range initListTestIsNil() {
		t.Run(tt.name, func(t *testing.T) {
			result := AllNil(tt.value)
			logger.Info("AllNil:", result)
		})
	}
}

func TestAllNotNil(t *testing.T) {
	for _, tt := range initListTestIsNotNil() {
		t.Run(tt.name, func(t *testing.T) {
			result := AllNotNil(tt.value)
			logger.Info("AllNotNil:", result)
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

func TestAllEmpty(t *testing.T) {
	for _, tt := range initListTestIsEmpty() {
		t.Run(tt.name, func(t *testing.T) {
			empty := AllEmpty(tt.value)
			logger.Info("AllEmpty:", empty)
		})
	}
}

func TestAllNotEmpty(t *testing.T) {
	for _, tt := range initListTestIsNotEmpty() {
		t.Run(tt.name, func(t *testing.T) {
			empty := AllNotEmpty(tt.value)
			logger.Info("AllNotEmpty:", empty)
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

func TestIsSliceEmpty(t *testing.T) {
	for _, tt := range initListTestIsSliceEmpty() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			isEmpty := IsSliceEmpty(tt.value)
			logger.Info("IsSliceEmpty:", isEmpty)
		})
	}
}

func TestIsSliceNotEmpty(t *testing.T) {
	for _, tt := range initListTestIsSliceNotEmpty() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			isNotEmpty := IsSliceNotEmpty(tt.value)
			logger.Info("IsSliceNotEmpty:", isNotEmpty)
		})
	}
}

func TestIsStringEmpty(t *testing.T) {
	for _, tt := range initListTestIsStringEmpty() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			isEmpty := IsStringEmpty(tt.value)
			logger.Info("IsStringEmpty:", isEmpty)
		})
	}
}

func TestIsStringNotEmpty(t *testing.T) {
	for _, tt := range initListTestIsStringNotEmpty() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			isNotEmpty := IsStringNotEmpty(tt.value)
			logger.Info("IsStringNotEmpty:", isNotEmpty)
		})
	}
}

func TestIsIntEmpty(t *testing.T) {
	for _, tt := range initListTestIsIntEmpty() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			isEmpty := IsIntEmpty(tt.value)
			logger.Info("IsIntEmpty:", isEmpty)
		})
	}
}

func TestIsIntNotEmpty(t *testing.T) {
	for _, tt := range initListTestIsIntNotEmpty() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			isNotEmpty := IsIntNotEmpty(tt.value)
			logger.Info("IsIntNotEmpty:", isNotEmpty)
		})
	}
}

func TestIsFloatEmpty(t *testing.T) {
	for _, tt := range initListTestIsFloatEmpty() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			isEmpty := IsFloatEmpty(tt.value)
			logger.Info("IsFloatEmpty:", isEmpty)
		})
	}
}

func TestIsFloatNotEmpty(t *testing.T) {
	for _, tt := range initListTestIsFloatNotEmpty() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			isNotEmpty := IsFloatNotEmpty(tt.value)
			logger.Info("IsFloatNotEmpty:", isNotEmpty)
		})
	}
}

func TestIsBoolEmpty(t *testing.T) {
	for _, tt := range initListTestIsBoolEmpty() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			isEmpty := IsBoolEmpty(tt.value)
			logger.Info("IsBoolEmpty:", isEmpty)
		})
	}
}

func TestIsBoolNotEmpty(t *testing.T) {
	for _, tt := range initListTestIsBoolNotEmpty() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			isNotEmpty := IsBoolNotEmpty(tt.value)
			logger.Info("IsBoolNotEmpty:", isNotEmpty)
		})
	}
}
