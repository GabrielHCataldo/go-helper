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

func TestIfNilReturns(t *testing.T) {
	var a1 *string
	var b1 string
	b1 = "b value non nil"
	result := IfNilReturns(a1, b1)
	log.Println("IfNilReturns:", result)
	var a2 string
	var b2 *string
	b2 = ConvertToPointer("a value non nil")
	result = IfNilReturns(b2, a2)
	log.Println("IfNilReturns:", result)
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

func TestIfEmptyReturns(t *testing.T) {
	var a1 string
	var b1 string
	b1 = "b value non empty"
	result := IfEmptyReturns(a1, b1)
	log.Println("IfEmptyReturns:", result)
	var a2 string
	var b2 string
	a2 = "a value non empty"
	result = IfEmptyReturns(a2, b2)
	log.Println("IfEmptyReturns:", result)
}

func TestReturnNonNilValue(t *testing.T) {
	var a *string
	var b *string
	var c *string
	value := "b value non empty"
	b = &value
	result := ReturnNonNilValue(a, b, c)
	log.Println("ReturnNonNilValue:", result)
}

func TestReturnNonEmptyValue(t *testing.T) {
	var a string
	var b string
	var c string
	c = "c value non empty"
	result := ReturnNonEmptyValue(a, b, c)
	log.Println("ReturnNonEmptyValue:", result)
}
