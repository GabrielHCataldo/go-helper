package helper

import (
	"log"
	"testing"
)

func TestGetCallerInfo(t *testing.T) {
	name, line, funcName := testReturn()
	log.Println(name+":"+line, "funcName:", funcName)
	name, line, funcName = GetCallerInfo(10)
	log.Println(name+":"+line, "funcName:", funcName)
}

func testReturn() (string, string, string) {
	return GetCallerInfo(1)
}
