package helper

import (
	"log"
	"testing"
)

func TestEquals(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := Equals(tt.value[0], tt.value[1])
			log.Println("Equals:", empty)
		})
	}
}

func TestIsNotEqualTo(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsNotEqualTo(tt.value[0], tt.value[1])
			log.Println("IsNotEqualTo:", empty)
		})
	}
}

func TestEqualsIgnoreCase(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := EqualsIgnoreCase(tt.value[0], tt.value[1])
			log.Println("EqualsIgnoreCase:", empty)
		})
	}
}

func TestIsNotEqualToIgnoreCase(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsNotEqualToIgnoreCase(tt.value[0], tt.value[1])
			log.Println("IsNotEqualToIgnoreCase:", empty)
		})
	}
}

func TestEqualsLen(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := EqualsLen(tt.value[0], 4)
			log.Println("EqualsLen:", empty)
		})
	}
}

func TestIsNotEqualsLen(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsNotEqualsLen(tt.value[0], 2)
			log.Println("IsNotEqualsLen:", empty)
		})
	}
}

func TestContains(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := Contains(tt.value[0], tt.value[1])
			log.Println("Contains:", empty)
		})
	}
}

func TestContainsIgnoreCase(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := ContainsIgnoreCase(tt.value[0], tt.value[1])
			log.Println("ContainsIgnoreCase:", empty)
		})
	}
}

func TestNotContains(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := NotContains(tt.value[0], tt.value[1])
			log.Println("NotContains:", empty)
		})
	}
}

func TestNotContainsIgnoreCase(t *testing.T) {
	for _, tt := range initListTestEquals() {
		t.Run(tt.name, func(t *testing.T) {
			empty := NotContainsIgnoreCase(tt.value[0], tt.value[1])
			log.Println("NotContainsIgnoreCase:", empty)
		})
	}
}

func TestIsGreaterThan(t *testing.T) {
	for _, tt := range initListTestGreaterThan() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsGreaterThan(tt.value[0], tt.value[1])
			log.Println("IsGreaterThan:", empty)
		})
	}
}

func TestIsGreaterThanOrEqual(t *testing.T) {
	for _, tt := range initListTestGreaterThan() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsGreaterThanOrEqual(tt.value[0], tt.value[1])
			log.Println("IsGreaterThanOrEqual:", empty)
		})
	}
}

func TestIsLessThan(t *testing.T) {
	for _, tt := range initListTestLessThan() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsLessThan(tt.value[0], tt.value[1])
			log.Println("IsLessThan:", empty)
		})
	}
}

func TestIsLessThanOrEqual(t *testing.T) {
	for _, tt := range initListTestLessThan() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsLessThanOrEqual(tt.value[0], tt.value[1])
			log.Println("IsLessThanOrEqual:", empty)
		})
	}
}
