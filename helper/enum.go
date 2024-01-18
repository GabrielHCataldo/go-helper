package helper

import "time"

type exampleEnumStr string
type exampleEnumInt int
type exampleEnumFloat float64
type exampleEnumBool bool
type exampleTime time.Time

const (
	exampleEnumStr1 exampleEnumStr = "test1"
)
const (
	exampleEnumInt1 exampleEnumInt = iota + 1
)
const (
	exampleEnumFloat1 exampleEnumFloat = iota + 1.3
)
const (
	exampleEnumBool1 exampleEnumBool = true
)

type BaseEnum interface {
	// IsEnumValid return true if value is valid, using on tag validate "enum"
	IsEnumValid() bool
}
