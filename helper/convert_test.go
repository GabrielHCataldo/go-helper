package helper

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"log"
	"os"
	"testing"
	"time"
)

func TestConvertToPointer(t *testing.T) {
	v := "string text"
	result := ConvertToPointer(v)
	log.Println("ConvertToPointer:", result)
}

func TestConvertPointerToValue(t *testing.T) {
	v := "string text"
	result := ConvertPointerToValue(&v)
	log.Println("ConvertPointerToValue:", result)
}

func TestConvertToObjectId(t *testing.T) {
	result, err := ConvertToObjectId(nil)
	log.Println("ConvertToObjectId:", result, "err:", err)
	result, err = ConvertToObjectId(ConvertToPointer(primitive.NewObjectID().Hex()))
	log.Println("ConvertToObjectId:", result, "err:", err)
}

func TestSimpleConvertToObjectId(t *testing.T) {
	v := ""
	result := SimpleConvertToObjectId(v)
	log.Println("SimpleConvertToObjectId:", result)
}

func TestConvertToPrimitiveDateTime(t *testing.T) {
	result, err := ConvertToPrimitiveDateTime(nil)
	log.Println("ConvertToPrimitiveDateTime:", result, "err:", err)
	result, err = ConvertToPrimitiveDateTime(ConvertToPointer(primitive.NewDateTimeFromTime(time.Now())))
	log.Println("ConvertToPrimitiveDateTime:", result, "err:", err)
	result, err = ConvertToPrimitiveDateTime(ConvertToPointer(10))
	log.Println("ConvertToPrimitiveDateTime:", result, "err:", err)
}

func TestSimpleConvertToPrimitiveDateTime(t *testing.T) {
	v := ""
	result := SimpleConvertToPrimitiveDateTime(v)
	log.Println("SimpleConvertToPrimitiveDateTime:", result)
}

func TestConvertByteUnit(t *testing.T) {
	for _, tt := range initListTestConvertByteUnit() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertByteUnit(tt.value)
			if (err != nil) != tt.wantErr {
				log.Printf("ConvertByteUnit() error = %v, wantErr %v", err, tt.wantErr)
				t.Fail()
			} else if err != nil {
				log.Println("err expected:", err)
			} else {
				log.Println("result:", result)
			}
		})
	}
}

func TestSimpleConvertByteUnit(t *testing.T) {
	result := SimpleConvertByteUnit("")
	log.Println("SimpleConvertByteUnit:", result)
}

func TestConvertMegaByteUnit(t *testing.T) {
	for _, tt := range initListTestConvertMegaByteUnit() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertMegaByteUnit(tt.value)
			if (err != nil) != tt.wantErr {
				log.Printf("ConvertMegaByteUnit() error = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
			} else if err != nil {
				log.Println("err expected:", err)
			} else {
				log.Println("result:", result)
			}
		})
	}
}

func TestSimpleConvertMegaByteUnit(t *testing.T) {
	result := SimpleConvertMegaByteUnit("")
	log.Println("SimpleConvertMegaByteUnit:", result)
}

func TestConvertToString(t *testing.T) {
	for _, tt := range initListTestConvertToString() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertToString(tt.value)
			if (err != nil) != tt.wantErr {
				log.Printf("ConvertToString() err = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
			} else {
				log.Println("result:", result, "err:", err)
			}
		})
	}
}

func TestSimpleConvertToString(t *testing.T) {
	result := SimpleConvertToString("")
	log.Println("SimpleConvertToString:", result)
}

func TestConvertToInt(t *testing.T) {
	for _, tt := range initListTestConvertToInt() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertToInt(tt.value)
			if (err != nil) != tt.wantErr {
				log.Printf("ConvertToInt() err = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
			} else {
				log.Println("result:", result, "err:", err)
			}
		})
	}
}

func TestLen(t *testing.T) {
	for _, tt := range initListTestLen() {
		t.Run(tt.name, func(t *testing.T) {
			result := Len(tt.value)
			log.Println("result:", result)
		})
	}
}

func TestSimpleConvertToInt(t *testing.T) {
	result := SimpleConvertToInt("")
	log.Println("SimpleConvertToInt:", result)
}

func TestConvertToFloat(t *testing.T) {
	for _, tt := range initListTestConvertToFloat() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertToFloat(tt.value)
			if (err != nil) != tt.wantErr {
				log.Printf("ConvertToFloat() err = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
			} else {
				log.Println("result:", result, "err:", err)
			}
		})
	}
}

func TestSimpleConvertToFloat(t *testing.T) {
	result := SimpleConvertToFloat("")
	log.Println("SimpleConvertToFloat:", result)
}

func TestConvertToBool(t *testing.T) {
	for _, tt := range initListTestConvertToBool() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertToBool(tt.value)
			if (err != nil) != tt.wantErr {
				log.Printf("ConvertToBool() err = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
			} else {
				log.Println("result:", result, "err:", err)
			}
		})
	}
}

func TestSimpleConvertToBool(t *testing.T) {
	result := SimpleConvertToBool("")
	log.Println("SimpleConvertToBool:", result)
}

func TestConvertToTime(t *testing.T) {
	for _, tt := range initListTestConvertToTime() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertToTime(tt.value)
			if (err != nil) != tt.wantErr {
				log.Printf("ConvertToString() err = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
			} else {
				log.Println("result:", result)
			}
		})
	}
}

func TestSimpleConvertToTime(t *testing.T) {
	result := SimpleConvertToTime("")
	log.Println("SimpleConvertToTime:", result)
}

func TestConvertToBytes(t *testing.T) {
	for _, tt := range initListTestConvertToString() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertToBytes(tt.value)
			if (err != nil) != tt.wantErr {
				log.Printf("ConvertToBytes() err = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
			} else {
				log.Println("result:", string(result), "error:", err)
			}
		})
	}
}

func TestSimpleConvertToBytes(t *testing.T) {
	result := SimpleConvertToBytes("")
	log.Println("SimpleConvertToBytes:", result)
}

func TestConvertFileToBytes(t *testing.T) {
	for _, tt := range initListTestConvertFileToBytes() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertFileToBytes(tt.value)
			if (err != nil) != tt.wantErr {
				log.Printf("ConvertFileToBytes() err = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
			} else {
				log.Println("result:", result, "error:", err)
			}
		})
	}
}

func TestSimpleConvertFileToBytes(t *testing.T) {
	result := SimpleConvertFileToBytes(os.NewFile(0, "test"))
	log.Println("SimpleConvertFileToBytes:", result)
}

func TestConvertToFile(t *testing.T) {
	for _, tt := range initListTestConvertToFile() {
		t.Run(tt.name, func(t *testing.T) {
			var resultPrint any
			result, err := ConvertToFile(tt.value)
			defer closeFile(result)
			if (err != nil) != tt.wantErr {
				log.Printf("ConvertToFile() err = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
			} else if result != nil {
				bs, _ := io.ReadAll(result)
				resultPrint = string(bs)
			}
			log.Println("result:", resultPrint, "error:", err)
		})
	}
}

func TestSimpleConvertToFile(t *testing.T) {
	result := SimpleConvertToFile("test file")
	bs, _ := io.ReadAll(result)
	log.Println("SimpleConvertToFile:", string(bs))
}

func TestConvertToReader(t *testing.T) {
	for _, tt := range initListTestConvertFileToBytes() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertToReader(tt.value)
			if (err != nil) != tt.wantErr {
				log.Printf("ConvertToReader() err = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
			} else if result != nil {
				bs, _ := io.ReadAll(result)
				log.Println("result:", string(bs), "error:", err)
			}
		})
	}
}

func TestSimpleConvertToReader(t *testing.T) {
	result := SimpleConvertToReader("test file")
	bs, _ := io.ReadAll(result)
	log.Println("SimpleConvertToReader:", string(bs))
}

func TestConvertToBuffer(t *testing.T) {
	for _, tt := range initListTestConvertFileToBytes() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertToBuffer(tt.value)
			if (err != nil) != tt.wantErr {
				log.Printf("ConvertToBuffer() err = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
			} else if result != nil {
				bs, _ := io.ReadAll(result)
				log.Println("result:", string(bs), "error:", err)
			}
		})
	}
}

func TestSimpleConvertToBuffer(t *testing.T) {
	result := SimpleConvertToBuffer("test file")
	bs, _ := io.ReadAll(result)
	log.Println("SimpleConvertToBuffer:", string(bs))
}

func TestConvertToBase64(t *testing.T) {
	for _, tt := range initListTestConvertToString() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertToBase64(tt.value)
			if (err != nil) != tt.wantErr {
				log.Printf("ConvertToBase64() err = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
				return
			}
			log.Println("result:", result, "error:", err)
		})
	}
}

func TestSimpleConvertToBase64(t *testing.T) {
	result := SimpleConvertToBase64("test file")
	log.Println("SimpleConvertToBase64:", result)
}

func TestConvertToDest(t *testing.T) {
	for _, tt := range initListTestConvertToDest() {
		t.Run(tt.name, func(t *testing.T) {
			err := ConvertToDest(tt.value, tt.dest)
			if (err != nil) != tt.wantErr {
				log.Printf("ConvertToDest() err = %v, wantErr = %v", err, tt.wantErr)
				t.Fail()
			} else {
				log.Println("result:", SimpleConvertToString(tt.dest), "error:", err)
			}
		})
	}
}
