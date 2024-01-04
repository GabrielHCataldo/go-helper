package helper

import (
	"github.com/GabrielHCataldo/go-logger/logger"
	"testing"
)

func TestIsNumeric(t *testing.T) {
	for _, tt := range initListTestIsNumeric() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNumeric(tt.value)
			logger.Info("IsNumeric:", result)
		})
	}
}

func TestIsLetter(t *testing.T) {
	for _, tt := range initListTestIsLetter() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsLetter(tt.value)
			logger.Info("IsLetter:", result)
		})
	}
}

func TestIsStringJson(t *testing.T) {
	for _, tt := range initListTestIsStringJson() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsStringJson(tt.value)
			logger.Info("IsStringJson:", result)
		})
	}
}

func TestIsNotStringJson(t *testing.T) {
	for _, tt := range initListTestIsStringJson() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNotStringJson(tt.value)
			logger.Info("IsNotStringJson:", result)
		})
	}
}

func TestIsStringInt(t *testing.T) {
	for _, tt := range initListTestIsStringInt() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsStringInt(tt.value)
			logger.Info("IsStringInt:", result)
		})
	}
}

func TestIsNotStringInt(t *testing.T) {
	for _, tt := range initListTestIsStringInt() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNotStringInt(tt.value)
			logger.Info("IsNotStringInt:", result)
		})
	}
}

func TestIsStringBool(t *testing.T) {
	for _, tt := range initListTestIsStringBool() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsStringBool(tt.value)
			logger.Info("IsStringBool:", result)
		})
	}
}

func TestIsNotStringBool(t *testing.T) {
	for _, tt := range initListTestIsStringBool() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNotStringBool(tt.value)
			logger.Info("IsNotStringBool:", result)
		})
	}
}

func TestIsStringFloat(t *testing.T) {
	for _, tt := range initListTestIsStringFloat() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsStringFloat(tt.value)
			logger.Info("IsStringFloat:", result)
		})
	}
}

func TestIsNotStringFloat(t *testing.T) {
	for _, tt := range initListTestIsStringFloat() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNotStringFloat(tt.value)
			logger.Info("IsNotStringFloat:", result)
		})
	}
}

func TestIsStringTime(t *testing.T) {
	for _, tt := range initListTestIsStringTime() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsStringTime(tt.value)
			logger.Info("IsStringTime:", result)
		})
	}
}

func TestIsNotStringTime(t *testing.T) {
	for _, tt := range initListTestIsStringTime() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNotStringTime(tt.value)
			logger.Info("IsNotStringTime:", result)
		})
	}
}

func TestGetFirstLastName(t *testing.T) {
	for _, tt := range initListTestGetFirstLastName() {
		t.Run(tt.name, func(t *testing.T) {
			result := GetFirstLastName(tt.value)
			logger.Info("GetFirstLastName:", result)
		})
	}
}

func TestGetFileString(t *testing.T) {
	for _, tt := range initListTestGetFileString() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetFileString(tt.uri)
			if (err != nil) != tt.wantErr {
				logger.Errorf("GetFileString() error = %v, wantErr %v", err, tt.wantErr)
				t.Fail()
			} else if err != nil {
				logger.Info("GetFileString() err expected:", err)
			} else {
				logger.Info("GetFileString() result:", result)
			}
		})
	}
}

func TestGetFileJson(t *testing.T) {
	for _, tt := range initListTestGetFileJson() {
		t.Run(tt.name, func(t *testing.T) {
			var dest []map[string]any
			err := GetFileJson(tt.uri, &dest)
			if (err != nil) != tt.wantErr {
				logger.Errorf("GetFileJson() error = %v, wantErr %v", err, tt.wantErr)
				t.Fail()
			} else if err != nil {
				logger.Info("GetFileJson() err expected:", err)
			} else {
				logger.Info("GetFileJson() result:", dest)
			}
		})
	}
}

func TestRandomNumberStr(t *testing.T) {
	result := RandomNumberStr(1, 100)
	logger.Info("RandomNumberStr:", result)
}

func TestCleanAllRepeatSpaces(t *testing.T) {
	s := "Test text string  if   spaces repeat"
	result := CleanAllRepeatSpaces(s)
	logger.Info("CleanAllRepeatSpaces:", result)
}
