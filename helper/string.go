package helper

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// IsNumeric check any value is numeric, ex: v any is int return true, if string numeric return true, if bool return false.
func IsNumeric(a any) bool {
	s, _ := ConvertToString(a)
	regex := regexp.MustCompile(`^[0-9]+$`)
	return regex.MatchString(s)
}

// IsLetter check any value is letter, ex: v any is int return false, if string letter return true, if bool return true.
func IsLetter(a any) bool {
	s, _ := ConvertToString(a)
	regex := regexp.MustCompile(`^[A-Za-z]+$`)
	return regex.MatchString(s)
}

// IsStringJson check if string value is json return true, otherwise return false.
func IsStringJson(a any) bool {
	s, _ := ConvertToString(a)
	var js json.RawMessage
	return json.Unmarshal([]byte(s), &js) == nil
}

// IsNotStringJson check if string value is not json return true, otherwise return false.
func IsNotStringJson(a any) bool {
	return !IsStringJson(a)
}

// IsStringInt check if string value is int return true, otherwise return false.
func IsStringInt(a any) bool {
	s, _ := ConvertToString(a)
	_, err := strconv.Atoi(s)
	return err == nil
}

// IsNotStringInt check if string value is not int return true, otherwise return false.
func IsNotStringInt(a any) bool {
	return !IsStringInt(a)
}

// IsStringBool check if string value is bool return true, otherwise return false.
func IsStringBool(a any) bool {
	s, _ := ConvertToString(a)
	_, err := strconv.ParseBool(s)
	return err == nil
}

// IsNotStringBool check if string value is not bool return true, otherwise return false.
func IsNotStringBool(a any) bool {
	return !IsStringBool(a)
}

// IsStringFloat check if string value is float return true, otherwise return false.
func IsStringFloat(a any) bool {
	s, _ := ConvertToString(a)
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// IsNotStringFloat check if string value is not float return true, otherwise return false.
func IsNotStringFloat(a any) bool {
	return !IsStringFloat(a)
}

// IsStringTime check if string value is time return true, otherwise return false.
func IsStringTime(a any) bool {
	s, _ := ConvertToString(a)
	_, err := ConvertToTime(s)
	return err == nil
}

// IsNotStringTime check if string value is not time return true, otherwise return false.
func IsNotStringTime(a any) bool {
	return !IsStringTime(a)
}

// GetFirstLastName get first and last name by string value or string pointer, ex: Gabriel Henrique Cataldo -> Gabriel Cataldo.
func GetFirstLastName(a any) string {
	s, _ := ConvertToString(a)
	split := strings.Split(s, " ")
	firstName := split[0]
	lastName := split[len(split)-1]
	if IsNotEmpty(lastName) && firstName != lastName {
		return firstName + " " + lastName
	}
	return firstName
}

// GetFileString get file by uri string param
func GetFileString(uri string) (string, error) {
	bytes, err := os.ReadFile(uri)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// GetFileJson get file by uri string param, fill dest by file json
func GetFileJson(uri string, dest any) error {
	bytes, err := os.ReadFile(uri)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, dest)
}

// RandomNumberStr generate random numbers string by min and max parameters
func RandomNumberStr(min, max int) string {
	return strconv.Itoa(rand.Intn(MinInt(min-max, 1)+1) + min)
}

// CleanAllRepeatSpaces clean all repeat space, ex: Get All Girls in   the   \n party -> Get All Girls in the party
func CleanAllRepeatSpaces(v string) string {
	regex := regexp.MustCompile(` +`)
	regex2 := regexp.MustCompile(`^\s+|\s+$|\s+[\r\t\f\v]`)
	v = regex.ReplaceAllString(v, " ")
	v = regex2.ReplaceAllString(v, "")
	return v
}

// Sprintln get all values convert to text string
func Sprintln(a ...any) string {
	var ac []any
	for _, v := range a {
		s := SimpleConvertToString(v)
		if IsNotEmpty(s) {
			ac = append(ac, s)
		}
	}
	return strings.TrimRight(fmt.Sprintln(ac...), "\n")
}
