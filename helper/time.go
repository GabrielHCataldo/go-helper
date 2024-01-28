package helper

import "time"

// IsBeforeNow if the parameter entered was before now it will return true, otherwise it will return false
func IsBeforeNow(a any) bool {
	t, err := ConvertToTime(a)
	if err != nil {
		panic("the parameter A error to convert to time:" + err.Error())
	}
	return t.Before(time.Now())
}

// IsBeforeDateToday If the parameter entered is before today it will return true, otherwise it will return false
func IsBeforeDateToday(a any) bool {
	t, err := ConvertToTime(a)
	if err != nil {
		panic("the parameter A error to convert to time:" + err.Error())
	}
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, t.Location())
	return t.Before(time.Now())
}

// IsBeforeDate If parameter A has a date before the date of parameter B, it returns true, otherwise it will return false.
func IsBeforeDate(a, b any) bool {
	ta, err := ConvertToTime(a)
	if err != nil {
		panic("the parameter A error to convert to time:" + err.Error())
	}
	tb, err := ConvertToTime(b)
	if err != nil {
		panic("the parameter B error to convert to time:" + err.Error())
	}
	ta = time.Date(ta.Year(), ta.Month(), ta.Day(), 0, 0, 0, 0, ta.Location())
	tb = time.Date(tb.Year(), tb.Month(), tb.Day(), 0, 0, 0, 0, tb.Location())
	return ta.Before(tb)
}

// IsBefore If parameter A is before parameter B it will return true, otherwise it will return false.
func IsBefore(a, b any) bool {
	ta, err := ConvertToTime(a)
	if err != nil {
		panic("the parameter A error to convert to time:" + err.Error())
	}
	tb, err := ConvertToTime(b)
	if err != nil {
		panic("the parameter B error to convert to time:" + err.Error())
	}
	return ta.Before(tb)
}

// IsAfterNow if the parameter entered was after now it will return true, otherwise it will return false
func IsAfterNow(a any) bool {
	t, err := ConvertToTime(a)
	if err != nil {
		panic("the parameter A error to convert to time:" + err.Error())
	}
	return t.After(time.Now())
}

// IsAfterDateToday If the parameter entered is after today it will return true, otherwise it will return false
func IsAfterDateToday(a any) bool {
	t, err := ConvertToTime(a)
	if err != nil {
		panic("the parameter A error to convert to time:" + err.Error())
	}
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, t.Location())
	return t.After(time.Now())
}

// IsAfterDate If parameter A has a date after the date of parameter B, it returns true, otherwise it will return false.
func IsAfterDate(a, b any) bool {
	ta, err := ConvertToTime(a)
	if err != nil {
		panic("the parameter A error to convert to time:" + err.Error())
	}
	tb, err := ConvertToTime(b)
	if err != nil {
		panic("the parameter B error to convert to time:" + err.Error())
	}
	ta = time.Date(ta.Year(), ta.Month(), ta.Day(), 0, 0, 0, 0, ta.Location())
	tb = time.Date(tb.Year(), tb.Month(), tb.Day(), 0, 0, 0, 0, tb.Location())
	return ta.After(tb)
}

// IsAfter If parameter A is after parameter B it will return true, otherwise it will return false.
func IsAfter(a, b any) bool {
	ta, err := ConvertToTime(a)
	if err != nil {
		panic("the parameter A error to convert to time:" + err.Error())
	}
	tb, err := ConvertToTime(b)
	if err != nil {
		panic("the parameter B error to convert to time:" + err.Error())
	}
	return ta.After(tb)
}

// IsToday If the parameter entered is today it will return true, otherwise it will return false
func IsToday(a any) bool {
	t, err := ConvertToTime(a)
	if err != nil {
		panic("the parameter A error to convert to time:" + err.Error())
	}
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, t.Location())
	return t.Equal(now)
}

// IsNow If the parameter entered is now (not counting the nanosecond) it will return true, otherwise it will return false
func IsNow(a any) bool {
	t, err := ConvertToTime(a)
	if err != nil {
		panic("the parameter A error to convert to time:" + err.Error())
	}
	t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, t.Location())
	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), t.Second(), 0, t.Location())
	return t.Equal(now)
}

// IsFullNow If the parameter entered is now it will return true, otherwise it will return false
func IsFullNow(a any) bool {
	t, err := ConvertToTime(a)
	if err != nil {
		panic("the parameter A error to convert to time:" + err.Error())
	}
	return t.Equal(time.Now())
}
