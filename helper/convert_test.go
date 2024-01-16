package helper

import (
	"github.com/GabrielHCataldo/go-logger/logger"
	"io"
	"os"
	"testing"
)

func TestConvertToPointer(t *testing.T) {
	v := "string text"
	result := ConvertToPointer(v)
	logger.Info("ConvertToPointer:", result)
}

func TestConvertPointerToValue(t *testing.T) {
	v := "string text"
	result := ConvertPointerToValue(&v)
	logger.Info("ConvertPointerToValue:", result)
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
				logger.Info("result:", string(result), "error:", err)
			}
		})
	}
}

func TestSimpleConvertToBytes(t *testing.T) {
	result := SimpleConvertToBytes("")
	logger.Info("SimpleConvertToBytes:", result)
}

func TestConvertFileToBytes(t *testing.T) {
	for _, tt := range initListTestConvertFileToBytes() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertFileToBytes(tt.value)
			if (err != nil) != tt.wantErr {
				logger.Errorf("ConvertFileToBytes() err = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
			} else {
				logger.Info("result:", result, "error:", err)
			}
		})
	}
}

func TestSimpleConvertFileToBytes(t *testing.T) {
	result := SimpleConvertFileToBytes(os.NewFile(0, "test"))
	logger.Info("SimpleConvertFileToBytes:", result)
}

func TestConvertToFile(t *testing.T) {
	for _, tt := range initListTestConvertToFile() {
		t.Run(tt.name, func(t *testing.T) {
			var resultPrint any
			result, err := ConvertToFile(tt.value)
			defer closeFile(result)
			if (err != nil) != tt.wantErr {
				logger.Errorf("ConvertToFile() err = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
			} else if result != nil {
				bs, _ := io.ReadAll(result)
				resultPrint = string(bs)
			}
			logger.Info("result:", resultPrint, "error:", err)
		})
	}
}

func TestSimpleConvertToFile(t *testing.T) {
	result := SimpleConvertToFile("test file")
	bs, _ := io.ReadAll(result)
	logger.Info("SimpleConvertToFile:", string(bs))
}

func TestConvertToReader(t *testing.T) {
	for _, tt := range initListTestConvertFileToBytes() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertToReader(tt.value)
			if (err != nil) != tt.wantErr {
				logger.Errorf("ConvertToReader() err = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
			} else if result != nil {
				bs, _ := io.ReadAll(result)
				logger.Info("result:", string(bs), "error:", err)
			}
		})
	}
}

func TestSimpleConvertToReader(t *testing.T) {
	result := SimpleConvertToReader("test file")
	bs, _ := io.ReadAll(result)
	logger.Info("SimpleConvertToReader:", string(bs))
}

func TestConvertToBuffer(t *testing.T) {
	for _, tt := range initListTestConvertFileToBytes() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertToBuffer(tt.value)
			if (err != nil) != tt.wantErr {
				logger.Errorf("ConvertToBuffer() err = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
			} else if result != nil {
				bs, _ := io.ReadAll(result)
				logger.Info("result:", string(bs), "error:", err)
			}
		})
	}
}

func TestSimpleConvertToBuffer(t *testing.T) {
	result := SimpleConvertToBuffer("test file")
	bs, _ := io.ReadAll(result)
	logger.Info("SimpleConvertToBuffer:", string(bs))
}

func TestConvertToBase64(t *testing.T) {
	for _, tt := range initListTestConvertToString() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertToBase64(tt.value)
			if (err != nil) != tt.wantErr {
				logger.Errorf("ConvertToBase64() err = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
				return
			}
			logger.Info("result:", result, "error:", err)
		})
	}
}

func TestSimpleConvertToBase64(t *testing.T) {
	result := SimpleConvertToBase64("test file")
	logger.Info("SimpleConvertToBase64:", result)
}

func TestConvertToDest(t *testing.T) {
	for _, tt := range initListTestConvertToDest() {
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
