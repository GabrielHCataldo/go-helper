package helper

import (
	"log"
	"testing"
)

func TestRandomBool(t *testing.T) {
	b := RandomBool()
	log.Println("result:", b)
}
