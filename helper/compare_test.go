package helper

import (
	"log"
	"testing"
)

func TestIsEqual(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := Equals(tt.value...)
			log.Println("IsEqual:", empty)
		})
	}
}

func TestIsNotEqual(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsNotEqualTo(tt.value...)
			log.Println("IsNotEquals:", empty)
		})
	}
}

func TestIsEqualIgnoreCase(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := EqualsIgnoreCase(tt.value...)
			log.Println("IsEqualIgnoreCase:", empty)
		})
	}
}

func TestIsNotEqualIgnoreCase(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsNotEqualToIgnoreCase(tt.value...)
			log.Println("IsNotEqualIgnoreCase:", empty)
		})
	}
}

func TestIsGreater(t *testing.T) {
	for _, tt := range initListTestGreater() {
		t.Run(tt.name, func(t *testing.T) {
			var a any = 2
			values := tt.value
			if IsNotEmpty(tt.value) {
				a = tt.value[0]
				values = tt.value[1:]
			}
			empty := IsGreaterThan(a, values...)
			log.Println("IsGreaterThan:", empty)
		})
	}
}

func TestIsGreaterEqual(t *testing.T) {
	for _, tt := range initListTestGreater() {
		t.Run(tt.name, func(t *testing.T) {
			var a any = 2
			values := tt.value
			if IsNotEmpty(tt.value) {
				a = tt.value[0]
				values = tt.value[1:]
			}
			empty := IsGreaterThanOrEqual(a, values...)
			log.Println("IsGreaterEqual:", empty)
		})
	}
}

func TestIsSmaller(t *testing.T) {
	for _, tt := range initListTestSmaller() {
		t.Run(tt.name, func(t *testing.T) {
			var a any = 2
			values := tt.value
			if IsNotEmpty(tt.value) {
				a = tt.value[0]
				values = tt.value[1:]
			}
			empty := IsLessThan(a, values...)
			log.Println("IsLessThan:", empty)
		})
	}
}

func TestIsSmallerEqual(t *testing.T) {
	for _, tt := range initListTestSmaller() {
		t.Run(tt.name, func(t *testing.T) {
			var a any = 2
			values := tt.value
			if IsNotEmpty(tt.value) {
				a = tt.value[0]
				values = tt.value[1:]
			}
			empty := IsLessThanOrEqual(a, values...)
			log.Println("IsSmallerEqual:", empty)
		})
	}
}
