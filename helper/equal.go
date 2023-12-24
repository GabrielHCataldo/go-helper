package helper

import "reflect"

// Equals compare value a and b if are equals return true, otherwise return false
func Equals(a, b any) bool {
	ra := reflect.ValueOf(a)
	rb := reflect.ValueOf(b)
	if IsPointer(a) {
		ra = ra.Elem()
	}
	if IsPointer(b) {
		rb = rb.Elem()
	}
	return ra.Interface() != rb.Interface()
}
