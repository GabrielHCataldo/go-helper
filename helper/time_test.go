package helper

import (
	"log"
	"testing"
)

func TestIsBeforeNow(t *testing.T) {
	for _, tt := range initListTestTime() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr && tt.name != "failed value type b")
			result := IsBeforeNow(tt.value[0])
			log.Println("IsBeforeNow:", result)
		})
	}
}

func TestIsBeforeDateToday(t *testing.T) {
	for _, tt := range initListTestTime() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr && tt.name != "failed value type b")
			result := IsBeforeDateToday(tt.value[0])
			log.Println("IsBeforeDateToday:", result)
		})
	}
}

func TestIsBeforeDate(t *testing.T) {
	for _, tt := range initListTestTime() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			result := IsBeforeDate(tt.value[0], tt.value[1])
			log.Println("IsBeforeDate:", result)
		})
	}
}

func TestIsBefore(t *testing.T) {
	for _, tt := range initListTestTime() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			result := IsBefore(tt.value[0], tt.value[1])
			log.Println("IsBefore:", result)
		})
	}
}

func TestIsAfterNow(t *testing.T) {
	for _, tt := range initListTestTime() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr && tt.name != "failed value type b")
			result := IsAfterNow(tt.value[0])
			log.Println("IsAfterNow:", result)
		})
	}
}

func TestIsAfterDateToday(t *testing.T) {
	for _, tt := range initListTestTime() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr && tt.name != "failed value type b")
			result := IsAfterDateToday(tt.value[0])
			log.Println("IsAfterDateToday:", result)
		})
	}
}

func TestIsAfterDate(t *testing.T) {
	for _, tt := range initListTestTime() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			result := IsAfterDate(tt.value[0], tt.value[1])
			log.Println("IsAfterDate:", result)
		})
	}
}

func TestIsAfter(t *testing.T) {
	for _, tt := range initListTestTime() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr)
			result := IsAfter(tt.value[0], tt.value[1])
			log.Println("IsAfter:", result)
		})
	}
}

func TestIsToday(t *testing.T) {
	for _, tt := range initListTestTime() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr && tt.name != "failed value type b")
			result := IsToday(tt.value[0])
			log.Println("IsToday:", result)
		})
	}
}

func TestIsNow(t *testing.T) {
	for _, tt := range initListTestTime() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr && tt.name != "failed value type b")
			result := IsNow(tt.value[0])
			log.Println("IsNow:", result)
		})
	}
}

func TestIsFullNow(t *testing.T) {
	for _, tt := range initListTestTime() {
		t.Run(tt.name, func(t *testing.T) {
			defer panicRecovery(t, tt.wantErr && tt.name != "failed value type b")
			result := IsFullNow(tt.value[0])
			log.Println("IsFullNow:", result)
		})
	}
}
