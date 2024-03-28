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
			result := IsMap(tt.value)
			log.Println("IsMap:", result)
		})
	}
}

func TestIsNotStringJson(t *testing.T) {
	for _, tt := range initListTestIsStringJson() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNotMap(tt.value)
			log.Println("IsNotStringMap:", result)
		})
	}
}

func TestIsStringSlice(t *testing.T) {
	for _, tt := range initListTestIsStringSlice() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsSlice(tt.value)
			log.Println("IsStringSlice:", result)
		})
	}
}

func TestIsNotStringSlice(t *testing.T) {
	for _, tt := range initListTestIsStringSlice() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNotSlice(tt.value)
			log.Println("IsNotSlice:", result)
		})
	}
}

func TestIsStringInt(t *testing.T) {
	for _, tt := range initListTestIsStringInt() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsInt(tt.value)
			log.Println("IsStringInt:", result)
		})
	}
}

func TestIsNotStringInt(t *testing.T) {
	for _, tt := range initListTestIsStringInt() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNotInt(tt.value)
			log.Println("IsNotStringInt:", result)
		})
	}
}

func TestIsStringBool(t *testing.T) {
	for _, tt := range initListTestIsStringBool() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsBool(tt.value)
			log.Println("IsStringBool:", result)
		})
	}
}

func TestIsNotStringBool(t *testing.T) {
	for _, tt := range initListTestIsStringBool() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNotBool(tt.value)
			log.Println("IsNotStringBool:", result)
		})
	}
}

func TestIsStringFloat(t *testing.T) {
	for _, tt := range initListTestIsStringFloat() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsFloat(tt.value)
			log.Println("IsStringFloat:", result)
		})
	}
}

func TestIsNotStringFloat(t *testing.T) {
	for _, tt := range initListTestIsStringFloat() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNotFloat(tt.value)
			log.Println("IsNotFloat:", result)
		})
	}
}

func TestIsStringTime(t *testing.T) {
	for _, tt := range initListTestIsStringTime() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsTime(tt.value)
			log.Println("IsStringTime:", result)
		})
	}
}

func TestIsNotStringTime(t *testing.T) {
	for _, tt := range initListTestIsStringTime() {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNotTime(tt.value)
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
