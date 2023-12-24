package helper

import (
	"encoding/json"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func IsNumeric(v any) bool {
	if !IsString(v) {
		panic("value type is not a string")
	}
	s := GetRealValue(v).(string)
	regex := regexp.MustCompile(`^[0-9]+$`)
	return regex.MatchString(s)
}

func IsLetter(v any) bool {
	if !IsString(v) {
		panic("value type is not a string")
	}
	s := GetRealValue(v).(string)
	regex := regexp.MustCompile(`^[A-Za-z]+$`)
	return regex.MatchString(s)
}

func GetFirstLastName(v string) string {
	split := strings.Split(v, " ")
	firstName := split[0]
	lastName := split[len(split)-1]
	if IsNotEmpty(lastName) && firstName != lastName {
		return firstName + " " + lastName
	}
	return firstName
}

func GetFileString(uri string) (string, error) {
	bytes, err := os.ReadFile(uri)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func GetFileJson(uri string, dest any) error {
	bytes, err := os.ReadFile(uri)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, dest)
}

func RandomNumberStr(min, max int) string {
	return strconv.Itoa(rand.Intn(MinInt(min-max, 1)+1) + min)
}

func CleanAllRepeatSpace(v string) string {
	regex := regexp.MustCompile(` +`)
	regex2 := regexp.MustCompile(`^\s+|\s+$|\s+[\r\t\f\v]`)
	v = regex.ReplaceAllString(v, " ")
	v = regex2.ReplaceAllString(v, "")
	return v
}
