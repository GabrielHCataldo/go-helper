package helper

import (
	"log"
	"testing"
)

func TestIsNumeric(t *testing.T) {
	for _, tt := range initListTestIsNumeric() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNumeric(tt.value)
			log.Println("IsNumeric:", result)
		})
	}
}

func TestIsLetter(t *testing.T) {
	for _, tt := range initListTestIsLetter() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsLetter(tt.value)
			log.Println("IsLetter:", result)
		})
	}
}

func TestIsStringJson(t *testing.T) {
	for _, tt := range initListTestIsStringJson() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsStringJson(tt.value)
			log.Println("IsStringJson:", result)
		})
	}
}

func TestIsNotStringJson(t *testing.T) {
	for _, tt := range initListTestIsStringJson() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNotStringJson(tt.value)
			log.Println("IsNotStringJson:", result)
		})
	}
}

func TestIsStringInt(t *testing.T) {
	for _, tt := range initListTestIsStringInt() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsStringInt(tt.value)
			log.Println("IsStringInt:", result)
		})
	}
}

func TestIsNotStringInt(t *testing.T) {
	for _, tt := range initListTestIsStringInt() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNotStringInt(tt.value)
			log.Println("IsNotStringInt:", result)
		})
	}
}

func TestIsStringBool(t *testing.T) {
	for _, tt := range initListTestIsStringBool() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsStringBool(tt.value)
			log.Println("IsStringBool:", result)
		})
	}
}

func TestIsNotStringBool(t *testing.T) {
	for _, tt := range initListTestIsStringBool() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNotStringBool(tt.value)
			log.Println("IsNotStringBool:", result)
		})
	}
}

func TestIsStringFloat(t *testing.T) {
	for _, tt := range initListTestIsStringFloat() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsStringFloat(tt.value)
			log.Println("IsStringFloat:", result)
		})
	}
}

func TestIsNotStringFloat(t *testing.T) {
	for _, tt := range initListTestIsStringFloat() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNotStringFloat(tt.value)
			log.Println("IsNotStringFloat:", result)
		})
	}
}

func TestIsStringTime(t *testing.T) {
	for _, tt := range initListTestIsStringTime() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsStringTime(tt.value)
			log.Println("IsStringTime:", result)
		})
	}
}

func TestIsNotStringTime(t *testing.T) {
	for _, tt := range initListTestIsStringTime() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNotStringTime(tt.value)
			log.Println("IsNotStringTime:", result)
		})
	}
}

func TestGetFirstLastName(t *testing.T) {
	for _, tt := range initListTestGetFirstLastName() {
		t.Run(tt.name, func(t *testing.T) {
			result := GetFirstLastName(tt.value)
			log.Println("GetFirstLastName:", result)
		})
	}
}

func TestGetFileString(t *testing.T) {
	for _, tt := range initListTestGetFileString() {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetFileString(tt.uri)
			if (err != nil) != tt.wantErr {
				log.Printf("GetFileString() error = %v, wantErr %v", err, tt.wantErr)
				t.Fail()
			} else if err != nil {
				log.Println("GetFileString() err expected:", err)
			} else {
				log.Println("GetFileString() result:", result)
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
				log.Printf("GetFileJson() error = %v, wantErr %v", err, tt.wantErr)
				t.Fail()
			} else if err != nil {
				log.Println("GetFileJson() err expected:", err)
			} else {
				log.Println("GetFileJson() result:", dest)
			}
		})
	}
}

func TestRandomNumberStr(t *testing.T) {
	result := RandomNumberStr(1, 100)
	log.Println("RandomNumberStr:", result)
}

func TestCleanAllRepeatSpaces(t *testing.T) {
	s := "Test text string  if   spaces repeat"
	result := CleanAllRepeatSpaces(s)
	log.Println("CleanAllRepeatSpaces:", result)
}

func TestSprintln(t *testing.T) {
	log.Println("Sprintln:", Sprintln("test string value:", 2, "with bool:", true, "with struct:", initTestStruct()))
}
