package helper

import (
	"encoding/json"
	"errors"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

// ConvertByteUnit convert byte unit text to int ex: 1KB = 1024
func ConvertByteUnit(v any) (int, error) {
	s := ConvertToString(v)
	errDefault := errors.New("byte unit mal formatted ex: 100MB")
	regex := regexp.MustCompile(`^(\d+)\s?(\w+)?$`)
	match := regex.FindAllStringSubmatch(s, -1)
	if len(match) == 0 || len(match[0]) == 0 || len(match[0]) < 3 {
		return 0, errDefault
	}
	vInt, _ := strconv.Atoi(match[0][1])
	switch match[0][2] {
	case "B":
		return vInt, nil
	case "KB":
		return vInt * 1024, nil
	case "MB":
		return vInt * int(math.Pow(1024, 2)), nil
	case "GB":
		return vInt * int(math.Pow(1024, 3)), nil
	case "TB":
		return vInt * int(math.Pow(1024, 4)), nil
	case "PB":
		return vInt * int(math.Pow(1024, 5)), nil
	case "EB":
		return vInt * int(math.Pow(1024, 6)), nil
	case "ZB":
		return vInt * int(math.Pow(1024, 7)), nil
	case "YB":
		return vInt * int(math.Pow(1024, 8)), nil
	default:
		return 0, errDefault
	}
}

// ConvertMegaByteUnit convert megabyte unit text to int ex: 1GB = 1024
func ConvertMegaByteUnit(a any) (int, error) {
	s := ConvertToString(a)
	errDefault := errors.New("byte unit mal formatted ex: 100MB")
	regex := regexp.MustCompile(`^(\d+)\s?(\w+)?$`)
	match := regex.FindAllStringSubmatch(s, -1)
	if len(match) == 0 || len(match[0]) == 0 || len(match[0]) < 3 {
		return 0, errDefault
	}
	vInt, _ := strconv.Atoi(match[0][1])
	switch match[0][2] {
	case "MB":
		return vInt, nil
	case "GB":
		return vInt * 1024, nil
	case "TB":
		return vInt * int(math.Pow(1024, 2)), nil
	case "PB":
		return vInt * int(math.Pow(1024, 3)), nil
	case "EB":
		return vInt * int(math.Pow(1024, 4)), nil
	case "ZB":
		return vInt * int(math.Pow(1024, 5)), nil
	case "YB":
		return vInt * int(math.Pow(1024, 6)), nil
	default:
		return 0, errDefault
	}
}

// ConvertToString convert any value to beautiful string
func ConvertToString(a any) string {
	v := reflect.ValueOf(a)
	if v.Type().Kind() == reflect.Pointer {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		b, _ := json.Marshal(v.Interface())
		return string(b)
	default:
		return convertToStringByType(v.Interface())
	}
}

// ConvertToTime convert any value to time
func ConvertToTime(a any) (time.Time, error) {
	v := reflect.ValueOf(a)
	if v.Type().Kind() == reflect.Pointer {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		if t, ok := v.Interface().(time.Time); ok {
			return t, nil
		}
		return time.Time{}, errors.New("error type to parse time: " + reflect.TypeOf(a).Kind().String())
	default:
		return convertToTimeByType(v.Interface())
	}
}

// GetRealValue get real value by any, if pointer or interface return value non pointer or interface
func GetRealValue(v any) any {
	elem := reflect.ValueOf(v)
	if IsPointer(v) || IsInterface(v) {
		elem = elem.Elem()
	}
	return elem.Interface()
}

func convertToStringByType(a any) string {
	switch t := a.(type) {
	case int:
		return strconv.Itoa(t)
	case int8:
		return strconv.Itoa(int(t))
	case uint:
		return strconv.Itoa(int(t))
	case uint8:
		return strconv.Itoa(int(t))
	case int16:
		return strconv.Itoa(int(t))
	case uint16:
		return strconv.Itoa(int(t))
	case int32:
		return strconv.Itoa(int(t))
	case uint32:
		return strconv.Itoa(int(t))
	case int64:
		return strconv.Itoa(int(t))
	case uint64:
		return strconv.Itoa(int(t))
	case bool:
		return strconv.FormatBool(t)
	case float32:
		return strconv.FormatFloat(float64(t), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64)
	case string:
		return t
	default:
		return ""
	}
}

func convertToTimeByType(a any) (time.Time, error) {
	switch t := a.(type) {
	case int:
		return time.Unix(int64(t), 0), nil
	case int8:
		return time.Unix(int64(t), 0), nil
	case uint:
		return time.Unix(int64(t), 0), nil
	case uint8:
		return time.Unix(int64(t), 0), nil
	case int16:
		return time.Unix(int64(t), 0), nil
	case uint16:
		return time.Unix(int64(t), 0), nil
	case int32:
		return time.Unix(int64(t), 0), nil
	case uint32:
		return time.Unix(int64(t), 0), nil
	case int64:
		return time.Unix(t, 0), nil
	case uint64:
		return time.Unix(int64(t), 0), nil
	case float32:
		return time.Unix(int64(t), 0), nil
	case float64:
		return time.Unix(int64(t), 0), nil
	case string:
		return time.Parse(time.RFC3339, t)
	default:
		return time.Time{}, errors.New("error type to parse time: " + reflect.TypeOf(a).Kind().String())
	}
}
