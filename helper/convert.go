package helper

import (
	"encoding/json"
	"errors"
	"math"
	"reflect"
	"regexp"
	"strconv"
)

func ConvertByteUnit(v string) (int, error) {
	errDefault := errors.New("byte unit mal formatted ex: 100MB")
	regex := regexp.MustCompile(`^(\d+)\s?(\w+)?$`)
	match := regex.FindAllStringSubmatch(v, -1)
	if len(match) < 1 && len(match[0]) < 3 {
		return 0, errDefault
	}
	vInt, err := strconv.Atoi(match[0][1])
	if err != nil {
		return 0, err
	}
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

func ConvertMegaByteUnit(v string) (int, error) {
	errDefault := errors.New("byte unit mal formatted ex: 100MB")
	regex := regexp.MustCompile(`^(\d+)\s?(\w+)?$`)
	match := regex.FindAllStringSubmatch(v, -1)
	if len(match) < 1 && len(match[0]) < 3 {
		return 0, errDefault
	}
	vInt, err := strconv.Atoi(match[0][1])
	if err != nil {
		return 0, err
	}
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
		return 0, nil
	}
}

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
		return convertToStringByType(a)
	}
}

func convertToStringByType(a any) string {
	switch t := a.(type) {
	case int:
		return strconv.Itoa(t)
	case int8:
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
	case []byte:
		return string(t)
	case string:
		return t
	default:
		return ""
	}
}
