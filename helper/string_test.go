package helper

import (
	"github.com/GabrielHCataldo/go-logger/logger"
	"testing"
)

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

func TestCleanAllRepeatSpace(t *testing.T) {
	s := "Test text string  if   spaces repeat"
	result := CleanAllRepeatSpace(s)
	logger.Info("CleanAllRepeatSpace:", result)
}
