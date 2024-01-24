package helper

import (
	"log"
	"testing"
)

func TestEquals(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := Equals(tt.value...)
			log.Println("Equals:", empty)
		})
	}
}

func TestIsNotEqualTo(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsNotEqualTo(tt.value...)
			log.Println("IsNotEqualTo:", empty)
		})
	}
}

func TestEqualsIgnoreCase(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := EqualsIgnoreCase(tt.value...)
			log.Println("EqualsIgnoreCase:", empty)
		})
	}
}

func TestIsNotEqualToIgnoreCase(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsNotEqualToIgnoreCase(tt.value...)
			log.Println("IsNotEqualToIgnoreCase:", empty)
		})
	}
}

func TestIsGreaterThan(t *testing.T) {
	for _, tt := range initListTestGreaterThan() {
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

func TestIsGreaterThanOrEqual(t *testing.T) {
	for _, tt := range initListTestGreaterThan() {
		t.Run(tt.name, func(t *testing.T) {
			var a any = 2
			values := tt.value
			if IsNotEmpty(tt.value) {
				a = tt.value[0]
				values = tt.value[1:]
			}
			empty := IsGreaterThanOrEqual(a, values...)
			log.Println("IsGreaterThanOrEqual:", empty)
		})
	}
}

func TestIsLessThan(t *testing.T) {
	for _, tt := range initListTestLessThan() {
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

func TestIsLessThanOrEqual(t *testing.T) {
	for _, tt := range initListTestLessThan() {
		t.Run(tt.name, func(t *testing.T) {
			var a any = 2
			values := tt.value
			if IsNotEmpty(tt.value) {
				a = tt.value[0]
				values = tt.value[1:]
			}
			empty := IsLessThanOrEqual(a, values...)
			log.Println("IsLessThanOrEqual:", empty)
		})
	}
}
