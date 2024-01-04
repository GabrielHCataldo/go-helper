package helper

import (
	"github.com/GabrielHCataldo/go-logger/logger"
	"testing"
)

func TestConvertToPointer(t *testing.T) {
	v := "string text"
	result := ConvertToPointer(v)
	logger.Info("ConvertToPointer:", result)
}

func TestConvertToObjectId(t *testing.T) {
	v := ""
	result := ConvertToObjectId(v)
	logger.Info("ConvertToObjectId:", result)
}

func TestConvertToObjectIdWithErr(t *testing.T) {
	v := ""
	result, err := ConvertToObjectIdWithErr(v)
	logger.Info("ConvertToObjectId:", result, "err:", err)
}

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
				logger.Errorf("ConvertMegaByteUnit() error = %v, wantErr = %v", err, tt.wantErr)
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
			result, err := ConvertToString(tt.value)
			if (err != nil) != tt.wantErr {
				logger.Errorf("ConvertToString() err = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
			} else {
				logger.Info("result:", result, "err:", err)
			}
		})
	}
}

func TestConvertToInt(t *testing.T) {
	for _, tt := range initListTestConvertToInt() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertToInt(tt.value)
			if (err != nil) != tt.wantErr {
				logger.Errorf("ConvertToInt() err = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
			} else {
				logger.Info("result:", result, "err:", err)
			}
		})
	}
}

func TestConvertToFloat(t *testing.T) {
	for _, tt := range initListTestConvertToInt() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertToFloat(tt.value)
			if (err != nil) != tt.wantErr {
				logger.Errorf("ConvertToFloat() err = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
			} else {
				logger.Info("result:", result, "err:", err)
			}
		})
	}
}

func TestConvertToBool(t *testing.T) {
	for _, tt := range initListTestConvertToBool() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertToBool(tt.value)
			if (err != nil) != tt.wantErr {
				logger.Errorf("ConvertToBool() err = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
			} else {
				logger.Info("result:", result, "err:", err)
			}
		})
	}
}

func TestConvertToTime(t *testing.T) {
	for _, tt := range initListTestConvertToTime() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertToTime(tt.value)
			if (err != nil) != tt.wantErr {
				logger.Errorf("ConvertToString() err = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
			} else {
				logger.Info("result:", result)
			}
		})
	}
}

func TestConvertToBytes(t *testing.T) {
	for _, tt := range initListTestConvertToString() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertToBytes(tt.value)
			if (err != nil) != tt.wantErr {
				logger.Errorf("ConvertToBytes() err = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
			} else {
				logger.Info("result:", result, "error:", err)
			}
		})
	}
}

func TestConvertToDest(t *testing.T) {
	for _, tt := range initListTestConvertStringToDest() {
		t.Run(tt.name, func(t *testing.T) {
			err := ConvertToDest(tt.value, tt.dest)
			if (err != nil) != tt.wantErr {
				logger.Errorf("ConvertToDest() err = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
			} else {
				logger.Info("result:", tt.dest, "error:", err)
			}
		})
	}
}
