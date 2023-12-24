package helper

import "math/rand"

// RandomBool return random value bool
func RandomBool() bool {
	return rand.Intn(2) == 1
}
