package helper

import (
	"encoding/json"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// IsNumeric check any value is numeric, ex: v any is int return true, if string numeric return true, if bool return false.
func IsNumeric(v any) bool {
	s, _ := ConvertToString(v)
	regex := regexp.MustCompile(`^[0-9]+$`)
	return regex.MatchString(s)
}

// IsLetter check any value is letter, ex: v any is int return false, if string letter return true, if bool return true.
func IsLetter(v any) bool {
	s, _ := ConvertToString(v)
	regex := regexp.MustCompile(`^[A-Za-z]+$`)
	return regex.MatchString(s)
}

// IsStringJson check if string value is json return true, otherwise return false.
func IsStringJson(v any) bool {
	s, _ := ConvertToString(v)
	var js json.RawMessage
	return json.Unmarshal([]byte(s), &js) == nil
}

// IsNotStringJson check if string value is not json return true, otherwise return false.
func IsNotStringJson(v any) bool {
	return !IsStringJson(v)
}

// IsStringInt check if string value is int return true, otherwise return false.
func IsStringInt(v any) bool {
	s, _ := ConvertToString(v)
	_, err := strconv.Atoi(s)
	return err == nil
}

// IsNotStringInt check if string value is not int return true, otherwise return false.
func IsNotStringInt(v any) bool {
	return !IsStringInt(v)
}

// IsStringBool check if string value is bool return true, otherwise return false.
func IsStringBool(v any) bool {
	s, _ := ConvertToString(v)
	_, err := strconv.ParseBool(s)
	return err == nil
}

// IsNotStringBool check if string value is not bool return true, otherwise return false.
func IsNotStringBool(v any) bool {
	return !IsStringBool(v)
}

// IsStringFloat check if string value is float return true, otherwise return false.
func IsStringFloat(v any) bool {
	s, _ := ConvertToString(v)
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// IsNotStringFloat check if string value is not float return true, otherwise return false.
func IsNotStringFloat(v any) bool {
	return !IsStringFloat(v)
}

// IsStringTime check if string value is time return true, otherwise return false.
func IsStringTime(v any) bool {
	s, _ := ConvertToString(v)
	_, err := ConvertToTime(s)
	return err == nil
}

// IsNotStringTime check if string value is not time return true, otherwise return false.
func IsNotStringTime(v any) bool {
	return !IsStringTime(v)
}

// GetFirstLastName get first and last name by string value or string pointer, ex: Gabriel Henrique Cataldo -> Gabriel Cataldo.
func GetFirstLastName(v any) string {
	s, _ := ConvertToString(v)
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
