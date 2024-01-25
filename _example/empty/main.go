package main

import (
	"github.com/GabrielHCataldo/go-helper/helper"
	"log"
	"time"
)

type exampleStruct struct {
	Name      string
	BirthDate time.Time
	Map       map[string]any
}

func main() {
	emptyAndNotEmpty()
	ifEmptyReturns()
	returnNonEmptyValue()
	nilAndNotNil()
	ifNilReturns()
	returnNonNilValue()
}

func nilAndNotNil() {
	var anyPointer *any
	var anotherValue []any
	isNil := helper.IsNil(anyPointer, anotherValue)
	log.Println("is nil?", isNil)
	anyPointer = helper.ConvertToPointer(any("value string"))
	anotherValue = []any{12, "test"}
	isNotNil := helper.IsNotNil(anyPointer, anotherValue)
	log.Println("is not nil?", isNotNil)
}

func emptyAndNotEmpty() {
	var anyPointer *any
	nStruct := exampleStruct{}
	nBlankString := "  "
	nMap := map[string]any{}
	var nSlice []any
	isEmpty := helper.IsEmpty(anyPointer, nStruct, nBlankString, nMap, nSlice)
	log.Println("is empty?", isEmpty)
	anyPointer = helper.ConvertToPointer(any("value string"))
	nBlankString = "test"
	nStruct.Name = "test name"
	nMap["test"] = "value string"
	nSlice = append(nSlice, "test string not empty")
	isNotEmpty := helper.IsNotEmpty(anyPointer, nStruct, nBlankString, nMap, nSlice)
	log.Println("is not empty?", isNotEmpty)
}

func ifEmptyReturns() {
	var a1 string
	var b1 string
	b1 = "b value non empty"
	result := helper.IfEmptyReturns(a1, b1)
	log.Println("value not empty returned:", result)
	var a2 string
	var b2 string
	a2 = "a value non empty"
	result = helper.IfEmptyReturns(b2, a2)
	log.Println("value not empty returned:", result)
}

func ifNilReturns() {
	var a1 *string
	var b1 string
	b1 = "b value non nil"
	result := helper.IfNilReturns(a1, b1)
	log.Println("value not nil returned:", result)
	var a2 string
	var b2 *string
	a2 = "a value non nil"
	result = helper.IfNilReturns(b2, a2)
	log.Println("value not nil returned:", result)
}

func returnNonNilValue() {
	var a *string
	var b *string
	var c *string
	value := "b value non empty"
	b = &value
	result := helper.ReturnNonNilValue(a, b, c)
	log.Println("value not nil returned:", result)
}

func returnNonEmptyValue() {
	var a string
	var b string
	var c string
	c = "c value non empty"
	result := helper.ReturnNonEmptyValue(a, b, c)
	log.Println("value not empty returned:", result)
}
