package helper

import (
	"log"
	"testing"
)

func TestGetCallerInfo(t *testing.T) {
	name, line, funcName := GetCallerInfo(1)
	log.Println(name+":"+line, "funcName:", funcName)
}
