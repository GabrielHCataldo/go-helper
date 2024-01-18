package helper

import (
	"log"
	"testing"
)

func TestIsNil(t *testing.T) {
	for _, tt := range initListTestIsNil() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsNil(tt.value)
			log.Println("IsNil:", empty)
		})
	}
}

func TestIsNotNil(t *testing.T) {
	for _, tt := range initListTestIsNotNil() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsNotNil(tt.value)
			log.Println("IsNotNil:", empty)
		})
	}
}

func TestIsEmpty(t *testing.T) {
	for _, tt := range initListTestIsEmpty() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsEmpty(tt.value)
			log.Println("IsEmpty:", empty)
		})
	}
}

func TestIsNotEmpty(t *testing.T) {
	for _, tt := range initListTestIsNotEmpty() {
		t.Run(tt.name, func(t *testing.T) {
			empty := IsNotEmpty(tt.value)
			log.Println("IsNotEmpty:", empty)
		})
	}
}
