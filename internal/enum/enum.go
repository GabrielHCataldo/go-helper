package enum

import "time"

type ExampleEnumStr string
type ExampleEnumInt int
type ExampleEnumFloat float64
type ExampleEnumBool bool
type ExampleTime time.Time

const (
	ExampleEnumStr1 ExampleEnumStr = "test1"
)
const (
	ExampleEnumInt1 ExampleEnumInt = iota + 1
)
const (
	ExampleEnumFloat1 ExampleEnumFloat = iota + 1.3
)
const (
	ExampleEnumBool1 ExampleEnumBool = true
)
