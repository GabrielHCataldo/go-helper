package helper

import (
	"log"
	"testing"
)

func TestRound(t *testing.T) {
	v := 1.6
	result := Round(v, 2)
	log.Println("Round:", result)
}

func TestRoundUp(t *testing.T) {
	v := 1.5
	result := RoundUp(v, 2)
	log.Println("RoundUp:", result)
}

func TestRoundDown(t *testing.T) {
	v := 1.5
	result := RoundDown(v, 2)
	log.Println("RoundDown:", result)
}
