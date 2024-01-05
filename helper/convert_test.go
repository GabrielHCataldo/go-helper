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
	result, err := ConvertToObjectId("")
	logger.Info("ConvertToObjectId:", result, "err:", err)
}

func TestSimpleConvertToObjectId(t *testing.T) {
	v := ""
	result := SimpleConvertToObjectId(v)
	logger.Info("SimpleConvertToObjectId:", result)
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

func TestSimpleConvertByteUnit(t *testing.T) {
	result := SimpleConvertByteUnit("")
	logger.Info("SimpleConvertByteUnit:", result)
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

func TestSimpleConvertMegaByteUnit(t *testing.T) {
	result := SimpleConvertMegaByteUnit("")
	logger.Info("SimpleConvertMegaByteUnit:", result)
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

func TestSimpleConvertToString(t *testing.T) {
	result := SimpleConvertToString("")
	logger.Info("SimpleConvertToString:", result)
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

func TestSimpleConvertToInt(t *testing.T) {
	result := SimpleConvertToInt("")
	logger.Info("SimpleConvertToInt:", result)
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

func TestSimpleConvertToFloat(t *testing.T) {
	result := SimpleConvertToFloat("")
	logger.Info("SimpleConvertToFloat:", result)
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

func TestSimpleConvertToBool(t *testing.T) {
	result := SimpleConvertToBool("")
	logger.Info("SimpleConvertToBool:", result)
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

func TestSimpleConvertToTime(t *testing.T) {
	result := SimpleConvertToTime("")
	logger.Info("SimpleConvertToTime:", result)
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

func TestSimpleConvertToBytes(t *testing.T) {
	result := SimpleConvertToBytes("")
	logger.Info("SimpleConvertToBytes:", result)
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
