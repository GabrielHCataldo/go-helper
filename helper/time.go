package helper

import "time"

// IsBeforeNow if the parameter entered was before now it will return true, otherwise it will return false
func IsBeforeNow(a any) bool {
	if !IsTime(a) {
		panic("the parameter A is not of type time")
	}
	t := GetRealValue(a).(time.Time)
	return t.Before(time.Now())
}

// IsBeforeDateToday If the parameter entered is before today it will return true, otherwise it will return false
func IsBeforeDateToday(a any) bool {
	if !IsTime(a) {
		panic("the parameter A is not of type time")
	}
	t := GetRealValue(a).(time.Time)
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, t.Location())
	return t.Before(time.Now())
}

// IsBeforeDate If parameter A has a date before the date of parameter B, it returns true, otherwise it will return false.
func IsBeforeDate(a, b any) bool {
	if !IsTime(a) {
		panic("the parameter A is not of type time")
	}
	if !IsTime(b) {
		panic("the parameter B is not of type time")
	}
	ta := GetRealValue(a).(time.Time)
	ta = time.Date(ta.Year(), ta.Month(), ta.Day(), 0, 0, 0, 0, ta.Location())
	tb := GetRealValue(b).(time.Time)
	tb = time.Date(tb.Year(), tb.Month(), tb.Day(), 0, 0, 0, 0, tb.Location())
	return ta.Before(tb)
}

// IsBefore If parameter A is before parameter B it will return true, otherwise it will return false.
func IsBefore(a, b any) bool {
	if !IsTime(a) {
		panic("the parameter A is not of type time")
	}
	if !IsTime(b) {
		panic("the parameter B is not of type time")
	}
	ta := GetRealValue(a).(time.Time)
	tb := GetRealValue(b).(time.Time)
	return ta.Before(tb)
}

// IsAfterNow if the parameter entered was after now it will return true, otherwise it will return false
func IsAfterNow(a any) bool {
	if !IsTime(a) {
		panic("the parameter A is not of type time")
	}
	t := GetRealValue(a).(time.Time)
	return t.After(time.Now())
}

// IsAfterDateToday If the parameter entered is after today it will return true, otherwise it will return false
func IsAfterDateToday(a any) bool {
	if !IsTime(a) {
		panic("the parameter A is not of type time")
	}
	t := GetRealValue(a).(time.Time)
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, t.Location())
	return t.Before(time.Now())
}

// IsAfterDate If parameter A has a date after the date of parameter B, it returns true, otherwise it will return false.
func IsAfterDate(a, b any) bool {
	if !IsTime(a) {
		panic("the parameter A is not of type time")
	}
	if !IsTime(b) {
		panic("the parameter B is not of type time")
	}
	ta := GetRealValue(a).(time.Time)
	ta = time.Date(ta.Year(), ta.Month(), ta.Day(), 0, 0, 0, 0, ta.Location())
	tb := GetRealValue(b).(time.Time)
	tb = time.Date(tb.Year(), tb.Month(), tb.Day(), 0, 0, 0, 0, tb.Location())
	return ta.After(tb)
}

// IsAfter If parameter A is after parameter B it will return true, otherwise it will return false.
func IsAfter(a, b any) bool {
	if !IsTime(a) {
		panic("the parameter A is not of type time")
	}
	if !IsTime(b) {
		panic("the parameter B is not of type time")
	}
	ta := GetRealValue(a).(time.Time)
	tb := GetRealValue(b).(time.Time)
	return ta.After(tb)
}
