package helper

import (
	"github.com/GabrielHCataldo/go-logger/logger"
	"testing"
)

func TestConvertByteUnit(t *testing.T) {
	for _, tt := range initListTestConvertByteUnit() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertByteUnit(tt.value)
			if (err != nil) != tt.wantErr {
				logger.Errorf("ConvertByteUnit() error = %v, wantErr %v", err, tt.wantErr)
				t.Fail()
			} else if err != nil {
				logger.Info("err expected:", err)
			} else {
				logger.Info("result:", result)
			}
		})
	}
}

func TestConvertMegaByteUnit(t *testing.T) {
	for _, tt := range initListTestConvertMegaByteUnit() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertMegaByteUnit(tt.value)
			if (err != nil) != tt.wantErr {
				logger.Errorf("ConvertMegaByteUnit() error = %v, wantErr %v", err, tt.wantErr)
				t.Fail()
			} else if err != nil {
				logger.Info("err expected:", err)
			} else {
				logger.Info("result:", result)
			}
		})
	}
}

func TestConvertToString(t *testing.T) {
	for _, tt := range initListTestConvertToString() {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertToString(tt.value)
			if (IsEmpty(result)) != tt.wantEmpty {
				logger.Errorf("ConvertToString() result = %v, wantEmpty %v", result, tt.wantEmpty)
				t.Fail()
			} else {
				logger.Info("result:", result)
			}
		})
	}
}
