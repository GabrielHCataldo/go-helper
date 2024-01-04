package helper

import "reflect"

// Equal compare values if are equals return true, otherwise return false
func Equal(a ...any) bool {
	a2 := a
	for _, v := range a {
		for _, v2 := range a2 {
			rv := reflect.ValueOf(v)
			rv2 := reflect.ValueOf(v2)
			if IsPointer(v) {
				rv = rv.Elem()
			}
			if IsPointer(v2) {
				rv2 = rv2.Elem()
			}
			if ConvertToString(rv.Interface()) != ConvertToString(rv2.Interface()) {
				return false
			}
		}
	}
	return true
}

// NotEqual compare values if aren't equals return true, otherwise return false
func NotEqual(a ...any) bool {
	return !Equal(a...)
}
