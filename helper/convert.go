package helper

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"log"
	"math"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Len retrieves the size of the passed value, if it is not a slice, struct or map, the size of the parameter converted to
// a string is returned.
func Len(a any) int {
	if IsNil(a) {
		return 0
	}
	var i int
	a = getRealValue(a)
	r := reflect.ValueOf(a)
	if IsSlice(a) {
		i = r.Len()
	} else if IsStruct(a) {
		i = r.NumField()
	} else if IsMap(a) {
		i = len(r.MapKeys())
	} else {
		i = len(SimpleConvertToString(a))
	}
	return i
}

// ConvertToPointer convert any value to pointer
func ConvertToPointer[T any](t T) *T {
	return &t
}

// ConvertPointerToValue convert pointer value to value
func ConvertPointerToValue[T any](t *T) T {
	var result T
	if IsNotNil(t) {
		result = *t
	}
	return result
}

// ConvertToObjectId convert any value to primitive.ObjectID
func ConvertToObjectId(a any) (primitive.ObjectID, error) {
	if IsNil(a) {
		return primitive.NilObjectID, errors.New("error convert to objectId: value is nil")
	}
	v := reflect.ValueOf(a)
	if IsPointer(a) {
		v = v.Elem()
	}
	var objectId primitive.ObjectID
	canConvert := v.CanConvert(reflect.TypeOf(primitive.ObjectID{}))
	if canConvert {
		objectId = v.Convert(reflect.TypeOf(primitive.ObjectID{})).Interface().(primitive.ObjectID)
	} else {
		s := SimpleConvertToString(a)
		objectId, _ = primitive.ObjectIDFromHex(s)
	}
	return objectId, nil
}

// SimpleConvertToObjectId convert any value to primitive.ObjectID, if err return empty value, check using primitive.NilObjectID
func SimpleConvertToObjectId(a any) primitive.ObjectID {
	r, _ := ConvertToObjectId(a)
	return r
}

// ConvertToPrimitiveDateTime convert any value to primitive.DateTime
func ConvertToPrimitiveDateTime(a any) (primitive.DateTime, error) {
	if IsNil(a) {
		return primitive.NewDateTimeFromTime(time.Time{}), errors.New("error convert to primitive dateTime: value is nil")
	}
	v := reflect.ValueOf(a)
	if v.Type().Kind() == reflect.Pointer {
		v = v.Elem()
	}
	var dateTime primitive.DateTime
	canConvert := v.CanConvert(reflect.TypeOf(primitive.DateTime(0)))
	if canConvert {
		dateTime = v.Convert(reflect.TypeOf(primitive.DateTime(0))).Interface().(primitive.DateTime)
	} else if IsTime(a) {
		dateTime = primitive.NewDateTimeFromTime(v.Interface().(time.Time))
	}
	return dateTime, nil
}

// SimpleConvertToPrimitiveDateTime convert any value to primitive.DateTime, if err return empty value
func SimpleConvertToPrimitiveDateTime(a any) primitive.DateTime {
	d, _ := ConvertToPrimitiveDateTime(a)
	return d
}

// ConvertByteUnit convert byte unit text to int ex: 1KB = 1024
func ConvertByteUnit(a any) (int, error) {
	s := SimpleConvertToString(a)
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

// SimpleConvertByteUnit convert byte unit text to int ex: 1KB = 1024, if err return empty value
func SimpleConvertByteUnit(a any) int {
	r, _ := ConvertByteUnit(a)
	return r
}

// ConvertMegaByteUnit convert megabyte unit text to int ex: 1GB = 1024
func ConvertMegaByteUnit(a any) (int, error) {
	s := SimpleConvertToString(a)
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

// SimpleConvertMegaByteUnit convert megabyte unit text to int ex: 1GB = 1024, if err return empty value
func SimpleConvertMegaByteUnit(a any) int {
	r, _ := ConvertMegaByteUnit(a)
	return r
}

// ConvertToString convert any value to beautiful string
func ConvertToString(a any) (string, error) {
	if IsNil(a) {
		return "", errors.New("error convert to string: value is nil")
	}
	v := reflect.ValueOf(a)
	if IsPointer(a) {
		v = v.Elem()
	}
	if IsBytes(a) {
		return string(v.Interface().([]byte)), nil
	} else if IsTime(a) {
		t := v.Interface().(time.Time)
		return t.Format(time.RFC3339Nano), nil
	} else if IsFile(a) {
		f := v.Interface().(os.File)
		b, err := ConvertFileToBytes(&f)
		return string(b), err
	} else if IsBuffer(a) {
		bf := v.Interface().(bytes.Buffer)
		return string(bf.Bytes()), nil
	} else if IsReader(a) {
		if r, ok := v.Interface().(bytes.Reader); ok {
			bs, err := io.ReadAll(&r)
			return string(bs), err
		}
		if r, ok := v.Interface().(strings.Reader); ok {
			bs, err := io.ReadAll(&r)
			return string(bs), err
		}
		if r, ok := a.(io.Reader); ok {
			bs, err := io.ReadAll(r)
			return string(bs), err
		}
	} else if IsError(a) {
		return a.(error).Error(), nil
	} else if IsObjectId(a) {
		return v.Interface().(primitive.ObjectID).Hex(), nil
	} else if IsPrimitiveDateTime(a) {
		t := v.Interface().(primitive.DateTime)
		return t.Time().UTC().Format(time.RFC3339Nano), nil
	} else if IsJson(a) || IsInterface(a) {
		b, err := json.Marshal(v.Interface())
		return string(b), err
	}
	return convertToStringByType(v.Interface())
}

// SimpleConvertToString convert any value to beautiful string, if err return empty value
func SimpleConvertToString(a any) string {
	s, _ := ConvertToString(a)
	return s
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

// SimpleConvertToInt convert any value to int, if err return empty value
func SimpleConvertToInt(a any) int {
	i, _ := ConvertToInt(a)
	return i
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

// SimpleConvertToFloat convert any value to float, if err return empty value
func SimpleConvertToFloat(a any) float64 {
	f, _ := ConvertToFloat(a)
	return f
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

// SimpleConvertToBool convert any value to float, if err return empty value
func SimpleConvertToBool(a any) bool {
	b, _ := ConvertToBool(a)
	return b
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
	if IsJson(a) || IsInterface(a) || IsTime(a) {
		canConvert := v.CanConvert(reflect.TypeOf(time.Time{}))
		if canConvert {
			return v.Convert(reflect.TypeOf(time.Time{})).Interface().(time.Time), nil
		}
		return time.Time{}, errors.New("error type to parse time: " + v.Kind().String())
	} else {
		return convertToTimeByType(v.Interface())
	}
}

// SimpleConvertToTime convert any value to time, if err return empty value
func SimpleConvertToTime(a any) time.Time {
	t, _ := ConvertToTime(a)
	return t
}

// ConvertToBytes convert any value to bytes
func ConvertToBytes(a any) ([]byte, error) {
	if IsNil(a) {
		return []byte{}, errors.New("error convert to bytes: value is nil")
	} else if (IsJson(a) || IsInterface(a)) && IsNotError(a) && IsNotBytes(a) {
		return json.Marshal(a)
	} else {
		s, err := ConvertToString(a)
		return []byte(s), err
	}
}

// SimpleConvertToBytes convert any value to bytes, if err return empty value
func SimpleConvertToBytes(a any) []byte {
	bs, _ := ConvertToBytes(a)
	return bs
}

// ConvertFileToBytes convert os.File value to slice byte
func ConvertFileToBytes(file *os.File) ([]byte, error) {
	if IsNil(file) {
		return nil, errors.New("error convert file to bytes: value is nil")
	}
	stat, _ := file.Stat()
	bs := make([]byte, stat.Size())
	_, err := bufio.NewReader(file).Read(bs)
	return bs, err
}

// SimpleConvertFileToBytes convert os.File value to slice byte, without error
func SimpleConvertFileToBytes(file *os.File) []byte {
	b, _ := ConvertFileToBytes(file)
	return b
}

// ConvertToFile convert any value to os.File .txt base64
func ConvertToFile(a any) (*os.File, error) {
	if IsNil(a) {
		return nil, errors.New("error convert to file: value is nil")
	}
	s, err := ConvertToString(a)
	if IsNotNil(err) {
		return nil, err
	}
	s64 := base64.StdEncoding.EncodeToString([]byte(s))
	f, err := os.CreateTemp("", "tmp-helper-")
	var fr *os.File
	if IsNil(err) {
		defer closeFile(f)
		_, err = f.WriteString(s64)
		if IsNil(err) {
			err = f.Sync()
			if IsNil(err) {
				fr, err = os.Open(f.Name())
			}
		}
	}
	return fr, err
}

// SimpleConvertToFile convert any value to os.File .txt base64, without error
func SimpleConvertToFile(a any) *os.File {
	f, _ := ConvertToFile(a)
	return f
}

// ConvertToReader convert any value to io.Reader
func ConvertToReader(a any) (*bytes.Reader, error) {
	if IsNil(a) {
		return nil, errors.New("error convert to reader: value is nil")
	}
	bs, err := ConvertToBytes(a)
	return bytes.NewReader(bs), err
}

// SimpleConvertToReader convert any value to bytes.Reader, without error
func SimpleConvertToReader(a any) *bytes.Reader {
	r, _ := ConvertToReader(a)
	return r
}

// ConvertToBuffer convert any value to bytes.Buffer
func ConvertToBuffer(a any) (*bytes.Buffer, error) {
	if IsNil(a) {
		return nil, errors.New("error convert to buffer: value is nil")
	}
	bs, err := ConvertToBytes(a)
	return bytes.NewBuffer(bs), err
}

// SimpleConvertToBuffer convert any value to bytes.Buffer, without error
func SimpleConvertToBuffer(a any) *bytes.Buffer {
	bf, _ := ConvertToBuffer(a)
	return bf
}

// ConvertToBase64 convert any value to string base64
func ConvertToBase64(a any) (string, error) {
	if IsNil(a) {
		return "", errors.New("error convert to base64: value is nil")
	}
	bs, err := ConvertToBytes(a)
	return base64.StdEncoding.EncodeToString(bs), err
}

// SimpleConvertToBase64 convert any value to string base64, without error
func SimpleConvertToBase64(a any) string {
	s, _ := ConvertToBase64(a)
	return s
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
	rDest := reflect.ValueOf(dest)
	if IsInt(dest) {
		i, err := ConvertToInt(vInterface)
		rDest.Elem().Set(reflect.ValueOf(i))
		return err
	} else if IsFloat(dest) {
		f, err := ConvertToFloat(vInterface)
		rDest.Elem().Set(reflect.ValueOf(f))
		return err
	} else if IsBool(dest) {
		b, err := ConvertToBool(vInterface)
		rDest.Elem().Set(reflect.ValueOf(b))
		return err
	} else if IsString(dest) {
		s, err := ConvertToString(vInterface)
		rDest.Elem().Set(reflect.ValueOf(s).Convert(rDest.Type()))
		return err
	} else if IsTime(dest) {
		tm, err := ConvertToTime(vInterface)
		rDest.Elem().Set(reflect.ValueOf(tm))
		return err
	} else if IsFile(dest) {
		f, err := ConvertToFile(vInterface)
		if IsNil(err) {
			rDest.Elem().Set(reflect.ValueOf(ConvertPointerToValue(f)))
		}
		return err
	} else if IsReader(dest) {
		r, err := ConvertToReader(vInterface)
		rDest.Elem().Set(reflect.ValueOf(ConvertPointerToValue(r)))
		return err
	} else if IsBuffer(dest) {
		bf, err := ConvertToBuffer(vInterface)
		rDest.Elem().Set(reflect.ValueOf(ConvertPointerToValue(bf)))
		return err
	} else if IsObjectId(dest) {
		bf, err := ConvertToObjectId(vInterface)
		rDest.Elem().Set(reflect.ValueOf(bf))
		return err
	} else if IsPrimitiveDateTime(dest) {
		bf, err := ConvertToPrimitiveDateTime(vInterface)
		rDest.Elem().Set(reflect.ValueOf(bf))
		return err
	} else if IsJson(dest) {
		b, _ := ConvertToBytes(a)
		return json.Unmarshal(b, dest)
	} else if IsInterface(dest) {
		if IsJson(a) || IsStringMap(a) || IsStringSlice(a) {
			b, _ := ConvertToBytes(a)
			return json.Unmarshal(b, dest)
		} else {
			rDest.Elem().Set(reflect.ValueOf(a))
			return nil
		}
	} else {
		return errors.New("error convert to dest: unknown value dest")
	}
}

// SimpleConvertToDest convert value to dest param without error
func SimpleConvertToDest(a, dest any) {
	_ = ConvertToDest(a, dest)
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
		v := reflect.ValueOf(a)
		canConvert := v.CanConvert(reflect.TypeOf(""))
		if canConvert {
			return v.Convert(reflect.TypeOf("")).Interface().(string), nil
		}
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
		v := reflect.ValueOf(a)
		canConvert := v.CanConvert(reflect.TypeOf(0))
		if canConvert {
			return v.Convert(reflect.TypeOf(0)).Interface().(int), nil
		}
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
		v := reflect.ValueOf(a)
		canConvert := v.CanConvert(reflect.TypeOf(0.0))
		if canConvert {
			return v.Convert(reflect.TypeOf(0.0)).Interface().(float64), nil
		}
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
		v := reflect.ValueOf(a)
		canConvert := v.CanConvert(reflect.TypeOf(true))
		if canConvert {
			return v.Convert(reflect.TypeOf(true)).Interface().(bool), nil
		}
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
	case primitive.DateTime:
		return t.Time(), nil
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

func closeFile(f *os.File) {
	err := f.Close()
	if IsNotNil(err) {
		log.Println("error close file to convert:", err)
	}
}
