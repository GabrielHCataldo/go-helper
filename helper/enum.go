package helper

type BaseEnum interface {
	// IsEnumValid return true if value is valid, using on tag validate "enum"
	IsEnumValid() bool
}
