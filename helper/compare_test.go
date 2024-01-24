package helper

import (
	"log"
	"testing"
)

func TestIsEqual(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsEqual(tt.value...)
			log.Println("IsEqual:", empty)
		})
	}
}

func TestIsNotEqual(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsNotEqual(tt.value...)
			log.Println("IsNotEqual:", empty)
		})
	}
}

func TestIsEqualIgnoreCase(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsEqualIgnoreCase(tt.value...)
			log.Println("IsEqualIgnoreCase:", empty)
		})
	}
}

func TestIsNotEqualIgnoreCase(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsNotEqualIgnoreCase(tt.value...)
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
			empty := IsGreater(a, values...)
			log.Println("IsGreater:", empty)
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
			empty := IsGreaterEqual(a, values...)
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
			empty := IsSmaller(a, values...)
			log.Println("IsSmaller:", empty)
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
			empty := IsSmallerEqual(a, values...)
			log.Println("IsSmallerEqual:", empty)
		})
	}
}
