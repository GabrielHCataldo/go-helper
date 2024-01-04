package helper

import (
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

// ConvertToPointer convert any value to pointer
func ConvertToPointer[T any](t T) *T {
	return &t
}

// ConvertToObjectId convert any value to primitive.ObjectID, if err return empty value, check using primitive.NilObjectID
func ConvertToObjectId(a any) primitive.ObjectID {
	s, _ := ConvertToString(a)
	r, _ := primitive.ObjectIDFromHex(s)
	return r
}

// ConvertToObjectIdWithErr convert any value to primitive.ObjectID
func ConvertToObjectIdWithErr(a any) (primitive.ObjectID, error) {
	s, _ := ConvertToString(a)
	return primitive.ObjectIDFromHex(s)
}

// ConvertByteUnit convert byte unit text to int ex: 1KB = 1024
func ConvertByteUnit(v any) (int, error) {
	s, _ := ConvertToString(v)
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
	s, _ := ConvertToString(a)
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
func ConvertToString(a any) (string, error) {
	if IsNil(a) {
		return "", errors.New("error convert to string: value is nil")
	}
	v := reflect.ValueOf(a)
	if v.Type().Kind() == reflect.Pointer {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		if s, ok := v.Interface().([]byte); ok {
			return string(s), nil
		}
		b, err := json.Marshal(v.Interface())
		return string(b), err
	default:
		return convertToStringByType(v.Interface())
	}
}

// ConvertToInt convert any value to int
func ConvertToInt(a any) (int, error) {
	if IsNil(a) {
		return 0, errors.New("error convert to int: value is nil")
	}
	v := reflect.ValueOf(a)
	if v.Type().Kind() == reflect.Pointer {
		v = v.Elem()
	}
	return convertToIntByType(v.Interface())
}

// ConvertToFloat convert any value to float
func ConvertToFloat(a any) (float64, error) {
	if IsNil(a) {
		return 0, errors.New("error convert to float: value is nil")
	}
	v := reflect.ValueOf(a)
	if v.Type().Kind() == reflect.Pointer {
		v = v.Elem()
	}
	return convertToFloatByType(v.Interface())
}

// ConvertToBool convert any value to float
func ConvertToBool(a any) (bool, error) {
	if IsNil(a) {
		return false, errors.New("error convert to bool: value is nil")
	}
	v := reflect.ValueOf(a)
	if v.Type().Kind() == reflect.Pointer {
		v = v.Elem()
	}
	return convertToBoolByType(v.Interface())
}

// ConvertToTime convert any value to time
func ConvertToTime(a any) (time.Time, error) {
	if IsNil(a) {
		return time.Time{}, errors.New("error convert to time: value is nil")
	}
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

// ConvertToBytes convert any value to float
func ConvertToBytes(a any) ([]byte, error) {
	if IsNil(a) {
		return []byte{}, errors.New("error convert to bool: value is nil")
	} else if IsJson(a) && IsNotTime(a) && IsNotBytes(a) {
		return json.Marshal(a)
	} else {
		s, err := ConvertToString(a)
		return []byte(s), err
	}
}

// ConvertToDest convert value to dest param
func ConvertToDest(a, dest any) error {
	if IsNil(a) {
		return errors.New("error convert string to dest: value is nil")
	} else if IsNotPointer(dest) {
		return errors.New("error convert string to dest: dest is not a pointer")
	}
	v := reflect.ValueOf(a)
	if v.Type().Kind() == reflect.Pointer {
		v = v.Elem()
	}
	vInterface := v.Interface()
	reflectDest := reflect.ValueOf(dest)
	if IsJson(dest) && IsNotTime(dest) {
		b, _ := ConvertToBytes(vInterface)
		return json.Unmarshal(b, dest)
	} else if IsInt(dest) {
		i, err := ConvertToInt(vInterface)
		reflectDest.Elem().Set(reflect.ValueOf(i))
		return err
	} else if IsFloat(dest) {
		f, err := ConvertToFloat(vInterface)
		reflectDest.Elem().Set(reflect.ValueOf(f))
		return err
	} else if IsBool(dest) {
		b, err := ConvertToBool(vInterface)
		reflectDest.Elem().Set(reflect.ValueOf(b))
		return err
	} else if IsString(dest) {
		s, err := ConvertToString(vInterface)
		reflectDest.Elem().Set(reflect.ValueOf(s))
		return err
	} else if IsTime(dest) {
		tm, err := ConvertToTime(vInterface)
		reflectDest.Elem().Set(reflect.ValueOf(tm))
		return err
	} else {
		return errors.New("error convert to dest: unknown value dest")
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

func convertToStringByType(a any) (string, error) {
	switch t := a.(type) {
	case int:
		return strconv.Itoa(t), nil
	case int8:
		return strconv.Itoa(int(t)), nil
	case uint:
		return strconv.Itoa(int(t)), nil
	case uint8:
		return strconv.Itoa(int(t)), nil
	case int16:
		return strconv.Itoa(int(t)), nil
	case uint16:
		return strconv.Itoa(int(t)), nil
	case int32:
		return strconv.Itoa(int(t)), nil
	case uint32:
		return strconv.Itoa(int(t)), nil
	case int64:
		return strconv.Itoa(int(t)), nil
	case uint64:
		return strconv.Itoa(int(t)), nil
	case bool:
		return strconv.FormatBool(t), nil
	case float32:
		return strconv.FormatFloat(float64(t), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64), nil
	case string:
		return t, nil
	default:
		return "", errors.New("error convert to string from type: " + reflect.TypeOf(a).Kind().String())
	}
}

func convertToIntByType(a any) (int, error) {
	switch t := a.(type) {
	case int:
		return t, nil
	case int8:
		return int(t), nil
	case uint:
		return int(t), nil
	case uint8:
		return int(t), nil
	case int16:
		return int(t), nil
	case uint16:
		return int(t), nil
	case int32:
		return int(t), nil
	case uint32:
		return int(t), nil
	case int64:
		return int(t), nil
	case uint64:
		return int(t), nil
	case float32:
		return int(t), nil
	case float64:
		return int(t), nil
	case string:
		f, err := strconv.ParseFloat(t, 64)
		return int(f), err
	default:
		return 0, errors.New("error convert to int from type: " + reflect.TypeOf(a).Kind().String())
	}
}

func convertToFloatByType(a any) (float64, error) {
	switch t := a.(type) {
	case int:
		return float64(t), nil
	case int8:
		return float64(t), nil
	case uint:
		return float64(t), nil
	case uint8:
		return float64(t), nil
	case int16:
		return float64(t), nil
	case uint16:
		return float64(t), nil
	case int32:
		return float64(t), nil
	case uint32:
		return float64(t), nil
	case int64:
		return float64(t), nil
	case uint64:
		return float64(t), nil
	case float32:
		return float64(t), nil
	case float64:
		return t, nil
	case string:
		return strconv.ParseFloat(t, 64)
	default:
		return 0, errors.New("error convert to float from type: " + reflect.TypeOf(a).Kind().String())
	}
}

func convertToBoolByType(a any) (bool, error) {
	switch t := a.(type) {
	case bool:
		return t, nil
	case string:
		return strconv.ParseBool(t)
	default:
		return false, errors.New("error convert to bool from type: " + reflect.TypeOf(a).Kind().String())
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
		tm, err := time.Parse(time.Layout, t)
		if err == nil {
			return tm, err
		}
		tm, err = time.Parse(time.ANSIC, t)
		if err == nil {
			return tm, err
		}
		tm, err = time.Parse(time.UnixDate, t)
		if err == nil {
			return tm, err
		}
		tm, err = time.Parse(time.RubyDate, t)
		if err == nil {
			return tm, err
		}
		tm, err = time.Parse(time.RFC822, t)
		if err == nil {
			return tm, err
		}
		tm, err = time.Parse(time.RFC822Z, t)
		if err == nil {
			return tm, err
		}
		tm, err = time.Parse(time.RFC850, t)
		if err == nil {
			return tm, err
		}
		tm, err = time.Parse(time.RFC1123, t)
		if err == nil {
			return tm, err
		}
		tm, err = time.Parse(time.RFC1123Z, t)
		if err == nil {
			return tm, err
		}
		tm, err = time.Parse(time.RFC3339, t)
		if err == nil {
			return tm, err
		}
		tm, err = time.Parse(time.Kitchen, t)
		if err == nil {
			return tm, err
		}
		tm, err = time.Parse(time.Stamp, t)
		if err == nil {
			return tm, err
		}
		tm, err = time.Parse(time.DateTime, t)
		if err == nil {
			return tm, err
		}
		tm, err = time.Parse(time.DateOnly, t)
		if err == nil {
			return tm, err
		}
		tm, err = time.Parse(time.TimeOnly, t)
		return tm, err
	default:
		return time.Time{}, errors.New("error convert to parse time from type: " + reflect.TypeOf(a).Kind().String())
	}
}
