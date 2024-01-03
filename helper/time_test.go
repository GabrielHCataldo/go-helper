package helper

import (
	"github.com/GabrielHCataldo/go-logger/logger"
	"testing"
)

func TestIsBeforeNow(t *testing.T) {
	for _, tt := range initListTestTime() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr && tt.name != "failed value type b")
			result := IsBeforeNow(tt.value[0])
			logger.Info("IsBeforeNow:", result)
		})
	}
}

func TestIsBeforeDateToday(t *testing.T) {
	for _, tt := range initListTestTime() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr && tt.name != "failed value type b")
			result := IsBeforeDateToday(tt.value[0])
			logger.Info("IsBeforeDateToday:", result)
		})
	}
}

func TestIsBeforeDate(t *testing.T) {
	for _, tt := range initListTestTime() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			result := IsBeforeDate(tt.value[0], tt.value[1])
			logger.Info("IsBeforeDate:", result)
		})
	}
}

func TestIsBefore(t *testing.T) {
	for _, tt := range initListTestTime() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			result := IsBefore(tt.value[0], tt.value[1])
			logger.Info("IsBefore:", result)
		})
	}
}

func TestIsAfterNow(t *testing.T) {
	for _, tt := range initListTestTime() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr && tt.name != "failed value type b")
			result := IsAfterNow(tt.value[0])
			logger.Info("IsAfterNow:", result)
		})
	}
}

func TestIsAfterDateToday(t *testing.T) {
	for _, tt := range initListTestTime() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr && tt.name != "failed value type b")
			result := IsAfterDateToday(tt.value[0])
			logger.Info("IsAfterDateToday:", result)
		})
	}
}

func TestIsAfterDate(t *testing.T) {
	for _, tt := range initListTestTime() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			result := IsAfterDate(tt.value[0], tt.value[1])
			logger.Info("IsAfterDate:", result)
		})
	}
}

func TestIsAfter(t *testing.T) {
	for _, tt := range initListTestTime() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			result := IsAfter(tt.value[0], tt.value[1])
			logger.Info("IsAfter:", result)
		})
	}
}

func TestIsToday(t *testing.T) {
	for _, tt := range initListTestTime() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr && tt.name != "failed value type b")
			result := IsToday(tt.value[0])
			logger.Info("IsToday:", result)
		})
	}
}

func TestIsNow(t *testing.T) {
	for _, tt := range initListTestTime() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr && tt.name != "failed value type b")
			result := IsNow(tt.value[0])
			logger.Info("IsNow:", result)
		})
	}
}

func TestIsFullNow(t *testing.T) {
	for _, tt := range initListTestTime() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr && tt.name != "failed value type b")
			result := IsFullNow(tt.value[0])
			logger.Info("IsFullNow:", result)
		})
	}
}
