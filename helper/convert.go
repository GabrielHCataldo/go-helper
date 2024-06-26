package helper

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
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

var ErrInvalidFormatByteUnit = errors.New("byte unit invalid format ex: 123.3MB")
var ErrInvalidFormatMegaByteUnit = errors.New("megabyte unit invalid format ex: 102.3MB")

// Len retrieves the size of the passed value, if it is not a slice, struct or map, the size of the parameter converted to
// a string is returned.
func Len(a any) int {
	if IsNil(a) {
		return 0
	}
	var i int
	a = getRealValue(a)
	r := reflect.ValueOf(a)
	if IsSliceType(a) {
		i = r.Len()
	} else if IsStructType(a) {
		i = r.NumField()
	} else if IsMapType(a) {
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
	if IsPointerType(a) {
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
	} else if IsTimeType(a) {
		dateTime = primitive.NewDateTimeFromTime(v.Interface().(time.Time))
	}
	return dateTime, nil
}

// SimpleConvertToPrimitiveDateTime convert any value to primitive.DateTime, if err return empty value
func SimpleConvertToPrimitiveDateTime(a any) primitive.DateTime {
	d, _ := ConvertToPrimitiveDateTime(a)
	return d
}

// ConvertByteUnitStrToFloat convert byte unit text to int ex: 1KB = 1024.0
func ConvertByteUnitStrToFloat(a any) (float64, error) {
	s := SimpleConvertToString(a)

	regex := regexp.MustCompile(`^(\d+)\s?(\w+)?$`)
	match := regex.FindAllStringSubmatch(s, -1)

	if len(match) == 0 || len(match[0]) < 3 {
		return 0, ErrInvalidFormatByteUnit
	}

	qty, _ := strconv.ParseFloat(match[0][1], 64)
	unit := match[0][2]

	return convertToBytes(qty, unit)
}

// SimpleConvertByteUnitStrToFloat convert byte unit text to int ex: 1KB = 1024, if err return empty value
func SimpleConvertByteUnitStrToFloat(a any) float64 {
	r, _ := ConvertByteUnitStrToFloat(a)
	return r
}

// ConvertToByteUnitStr takes an integer value in bytes and converts it to a human-readable byte unit string.
// If the value is less than 1024, it returns the value with the unit "B" appended.
// If the value is greater than or equal to 1024, it calculates the appropriate unit (KB, MB, GB, etc.) and
// returns the value with the corresponding unit appended.
// The value is rounded to one decimal place.
// This function does not handle negative values.
func ConvertToByteUnitStr(a any) string {
	in := SimpleConvertToFloat(a)
	const unit = 1024.0
	if in < unit {
		return fmt.Sprintf("%.1fB", in)
	}
	div, exp := unit, 0
	for n := in / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%cB", in/div, "KMGTPEZY"[exp])
}

// ConvertMegaByteUnitStrToFloat convert megabyte unit text to int ex: 1GB = 1024
func ConvertMegaByteUnitStrToFloat(a any) (float64, error) {
	s := SimpleConvertToString(a)

	regex := regexp.MustCompile(`^(\d+)\s?(\w+)?$`)
	match := regex.FindAllStringSubmatch(s, -1)

	if len(match) == 0 || len(match[0]) < 3 {
		return 0, ErrInvalidFormatMegaByteUnit
	}

	qty, _ := strconv.ParseFloat(match[0][1], 64)
	unit := match[0][2]

	return convertToMegaBytes(qty, unit)
}

// SimpleConvertMegaByteUnitStrToFloat convert megabyte unit text to int ex: 1GB = 1024, if err return empty value
func SimpleConvertMegaByteUnitStrToFloat(a any) float64 {
	r, _ := ConvertMegaByteUnitStrToFloat(a)
	return r
}

// ConvertToMegaByteUnitStr converts an integer value to a string representation
// in megabyte units.
// It takes in a parameter `a` which can be of any type that can be converted to an integer.
// The function returns a string representation of the passed value in megabyte units.
// If the value is less than 1MB, it returns the value followed by the string "B" (e.g. "256B").
// If the value is greater than or equal to 1MB, it converts the value to the nearest megabyte,
// appends the appropriate unit symbol (K, M, G, T, P, E, Z, Y) based on the magnitude of the value,
// and appends the string "B" at the end (e.g. "4.5MB", "10.2GB", "3.7TB", etc.).
func ConvertToMegaByteUnitStr(a any) string {
	const unit = 1024.0 * 1024.0

	in := SimpleConvertToFloat(a)
	if in < unit {
		return fmt.Sprintf("%.1fB", in)
	}
	div, exp := unit, 0
	for n := in / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%cB", in/div, "KMGTPEZY"[exp])
}

// ConvertToString convert any value to beautiful string
func ConvertToString(a any) (string, error) {
	if IsNil(a) {
		return "", errors.New("error convert to string: value is nil")
	}
	v := reflect.ValueOf(a)
	if IsPointerType(a) {
		v = v.Elem()
	}
	if IsStringType(a) {
		return v.String(), nil
	} else if IsBytesType(a) {
		return string(v.Interface().([]byte)), nil
	} else if IsTimeType(a) {
		t := v.Interface().(time.Time)
		return t.Format(time.RFC3339Nano), nil
	} else if IsFileType(a) {
		f := v.Interface().(os.File)
		b, err := ConvertFileToBytes(&f)
		return string(b), err
	} else if IsBufferType(a) {
		bf := v.Interface().(bytes.Buffer)
		return string(bf.Bytes()), nil
	} else if IsReaderType(a) {
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
	} else if IsErrorType(a) {
		return a.(error).Error(), nil
	} else if IsObjectIdType(a) {
		return v.Interface().(primitive.ObjectID).Hex(), nil
	} else if IsPrimitiveDateTimeType(a) {
		t := v.Interface().(primitive.DateTime)
		return t.Time().UTC().Format(time.RFC3339Nano), nil
	} else if IsJsonType(a) || IsInterfaceType(a) {
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
	va := v.Interface()
	if IsJsonType(a) {
		va = SimpleConvertToString(va)
	}
	return convertToIntByType(va)
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
	va := v.Interface()
	if IsJsonType(a) {
		va = SimpleConvertToString(va)
	}
	return convertToFloatByType(va)
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
	va := v.Interface()
	if IsJsonType(a) {
		va = SimpleConvertToString(va)
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
	if IsJsonType(a) || IsInterfaceType(a) || IsTimeType(a) {
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

// ConvertToTimeDuration convert any value to time duration
func ConvertToTimeDuration(a any) (time.Duration, error) {
	if IsNil(a) {
		return 0, errors.New("error convert to time: value is nil")
	} else if IsIntType(a) {
		return time.Duration(SimpleConvertToInt(a)), nil
	}

	s := SimpleConvertToString(a)
	return time.ParseDuration(s)
}

// SimpleConvertToTimeDuration convert any value to time duration, if err return empty value
func SimpleConvertToTimeDuration(a any) time.Duration {
	d, _ := ConvertToTimeDuration(a)
	return d
}

// ConvertToBytes convert any value to bytes
func ConvertToBytes(a any) ([]byte, error) {
	if IsNil(a) {
		return []byte{}, errors.New("error convert to bytes: value is nil")
	} else if (IsJsonType(a) || IsInterfaceType(a)) && IsNotErrorType(a) && IsNotBytesType(a) {
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

func ConvertBase64ToString(a any) (string, error) {
	if IsNil(a) {
		return "", errors.New("error convert base64 to string: value is nil")
	}
	s, err := ConvertToString(a)
	if IsNotNil(err) {
		return "", err
	}
	decode, err := base64.StdEncoding.DecodeString(s)
	if IsNotNil(err) {
		return "", err
	}
	return string(decode), nil
}

func SimpleConvertBase64ToString(a any) string {
	s, _ := ConvertBase64ToString(a)
	return s
}

// ConvertToDest convert value to dest param
func ConvertToDest(a, dest any) error {
	if IsNil(a) {
		return errors.New("error convert string to dest: value is nil")
	} else if IsNotPointerType(dest) {
		return errors.New("error convert string to dest: dest is not a pointer")
	}

	v := reflect.ValueOf(a)
	if v.Type().Kind() == reflect.Pointer {
		v = v.Elem()
	}
	vInterface := v.Interface()
	rDest := reflect.ValueOf(dest)

	if IsIntType(dest) {
		i, err := ConvertToInt(vInterface)
		if IsNil(err) {
			setReflectDest(i, rDest)
		}
		return err
	} else if IsFloatType(dest) {
		f, err := ConvertToFloat(vInterface)
		if IsNil(err) {
			setReflectDest(f, rDest)
		}
		return err
	} else if IsBoolType(dest) {
		b, err := ConvertToBool(vInterface)
		if IsNil(err) {
			setReflectDest(b, rDest)
		}
		return err
	} else if IsStringType(dest) {
		s, err := ConvertToString(vInterface)
		if IsNil(err) {
			setReflectDest(s, rDest)
		}
		return err
	} else if IsTimeType(dest) {
		tm, err := ConvertToTime(vInterface)
		if IsNil(err) {
			setReflectDest(tm, rDest)
		}
		return err
	} else if IsFileType(dest) {
		f, err := ConvertToFile(vInterface)
		if IsNil(err) {
			setReflectDest(ConvertPointerToValue(f), rDest)
		}
		return err
	} else if IsReaderType(dest) {
		r, err := ConvertToReader(vInterface)
		if IsNil(err) {
			setReflectDest(ConvertPointerToValue(r), rDest)
		}
		return err
	} else if IsBufferType(dest) {
		bf, err := ConvertToBuffer(vInterface)
		if IsNil(err) {
			setReflectDest(ConvertPointerToValue(bf), rDest)
		}
		return err
	} else if IsObjectIdType(dest) {
		o, err := ConvertToObjectId(vInterface)
		if IsNil(err) {
			setReflectDest(o, rDest)
		}
		return err
	} else if IsPrimitiveDateTimeType(dest) {
		pdt, err := ConvertToPrimitiveDateTime(vInterface)
		if IsNil(err) {
			setReflectDest(pdt, rDest)
		}
		return err
	} else if IsJsonType(dest) {
		b, err := ConvertToBytes(a)
		if IsNil(err) {
			return json.Unmarshal(b, dest)
		}
		return err
	} else if IsInterfaceType(dest) {
		return convertToInterfaceDest(a, dest)
	} else {
		return errors.New("error convert to dest: unknown value dest")
	}
}

// SimpleConvertToDest convert value to dest param without error
func SimpleConvertToDest(a, dest any) {
	_ = ConvertToDest(a, dest)
}

func convertToInterfaceDest(a, dest any) error {
	rDest := reflect.ValueOf(dest)
	if IsInt(a) {
		i, err := ConvertToInt(a)
		if IsNil(err) {
			setReflectDest(i, rDest)
		}
	} else if IsFloat(a) {
		f, err := ConvertToFloat(a)
		if IsNil(err) {
			setReflectDest(f, rDest)
		}
	} else if IsBool(a) {
		b, err := ConvertToBool(a)
		if IsNil(err) {
			setReflectDest(b, rDest)
		}
	} else if IsTime(a) {
		tm, err := ConvertToTime(a)
		if IsNil(err) {
			setReflectDest(tm, rDest)
		}
		return err
	} else if IsJson(a) {
		b, _ := ConvertToBytes(a)
		return json.Unmarshal(b, dest)
	} else {
		rs := reflect.ValueOf(a)
		converted := rs.Convert(rDest.Elem().Type())
		rDest.Elem().Set(converted)
	}
	return nil
}

func convertToBytes(qty float64, unit string) (float64, error) {
	unitMap := map[string]float64{
		"B":  0,
		"KB": 1,
		"MB": 2,
		"GB": 3,
		"TB": 4,
		"PB": 5,
		"EB": 6,
		"ZB": 7,
		"YB": 8,
	}

	if exp, ok := unitMap[unit]; ok {
		return qty * math.Pow(1024, exp), nil
	}
	return 0, ErrInvalidFormatByteUnit
}

func convertToMegaBytes(qty float64, unit string) (float64, error) {
	unitMap := map[string]float64{
		"MB": 0,
		"GB": 1,
		"TB": 2,
		"PB": 3,
		"EB": 4,
		"ZB": 5,
		"YB": 6,
	}

	if exp, ok := unitMap[unit]; ok {
		return qty * math.Pow(1024, exp), nil
	}
	return 0, ErrInvalidFormatMegaByteUnit
}

func convertToStringByType(a any) (string, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(val.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(val.Uint(), 10), nil
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(val.Float(), 'f', -1, 64), nil
	case reflect.Bool:
		return strconv.FormatBool(val.Bool()), nil
	default:
		if val.CanInterface() {
			if stringer, ok := val.Interface().(fmt.Stringer); ok {
				return stringer.String(), nil
			}
		}
		return "", errors.New("error convert to string from type: " + reflect.TypeOf(a).Kind().String())
	}
}

func convertToIntByType(a any) (int, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int(val.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int(val.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return int(val.Float()), nil
	case reflect.String:
		v, err := strconv.ParseInt(val.String(), 10, 64)
		if err != nil {
			return 0, err
		}
		return int(v), nil
	default:
		if val.CanInterface() {
			if stringer, ok := val.Interface().(fmt.Stringer); ok {
				v, err := strconv.ParseInt(stringer.String(), 10, 64)
				if err != nil {
					return 0, err
				}
				return int(v), nil
			}
		}
	}
	return 0, fmt.Errorf("error convert to int from type: %s", val.Type().String())
}

func convertToFloatByType(a any) (float64, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(val.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(val.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return val.Float(), nil
	case reflect.String:
		v, err := strconv.ParseFloat(val.String(), 64)
		if err != nil {
			return 0, err
		}
		return v, nil
	default:
		if val.CanInterface() {
			if stringer, ok := val.Interface().(fmt.Stringer); ok {
				v, err := strconv.ParseFloat(stringer.String(), 64)
				if err != nil {
					return 0, err
				}
				return v, nil
			}
		}
	}
	return 0, fmt.Errorf("error convert to float from type: %s", val.Type().String())
}

func convertToBoolByType(a any) (bool, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Bool:
		return val.Bool(), nil
	case reflect.String:
		v, err := strconv.ParseBool(val.String())
		if err != nil {
			return false, err
		}
		return v, nil
	default:
		if val.CanInterface() {
			if stringer, ok := val.Interface().(fmt.Stringer); ok {
				v, err := strconv.ParseBool(stringer.String())
				if err != nil {
					return false, err
				}
				return v, nil
			}
		}
	}
	return false, fmt.Errorf("error convert to bool from type: %s", val.Type().String())
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

func setReflectDest(a any, rDest reflect.Value) {
	rs := reflect.ValueOf(a)
	converted := rs.Convert(rDest.Elem().Type())
	rDest.Elem().Set(converted)
}

func closeFile(f *os.File) {
	err := f.Close()
	if IsNotNil(err) {
		log.Println("error close file to convert:", err)
	}
}
