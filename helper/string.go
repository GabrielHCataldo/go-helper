package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// IsNumeric check any value is numeric, ex: v any is int return true, if string numeric return true, if bool return false.
func IsNumeric(a any) bool {
	s := SimpleConvertToString(a)
	regex := regexp.MustCompile(`^[0-9]+$`)
	return regex.MatchString(s)
}

// IsLetter check any value is letter, ex: v any is int return false, if string letter return true, if bool return true.
func IsLetter(a any) bool {
	s := SimpleConvertToString(a)
	regex := regexp.MustCompile(`^[A-Za-z]+$`)
	return regex.MatchString(s)
}

// IsUrlPath check any value is url path ex: "/api/users"
func IsUrlPath(a any) bool {
	s := SimpleConvertToString(a)
	regex := regexp.MustCompile(`^/([^/\s]*)+(/[^/\s]+)*$`)
	return regex.MatchString(s)
}

// IsNotUrlPath check any value is not url path
func IsNotUrlPath(a any) bool {
	return !IsUrlPath(a)
}

// GetFirstLastName get first and last name by string value or string pointer, ex: Gabriel Henrique Cataldo -> Gabriel Cataldo.
func GetFirstLastName(a any) string {
	s := SimpleConvertToString(a)
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
	return strconv.Itoa(RandomNumber(min, max))
}

// CleanAllRepeatSpaces clean all repeat space, ex: Get All Girls in   the   \n party -> Get All Girls in the party
func CleanAllRepeatSpaces(v string) string {
	regex := regexp.MustCompile(` +`)
	regex2 := regexp.MustCompile(`^\s+|\s+$|\s+[\r\t\f\v]`)
	v = regex.ReplaceAllString(v, " ")
	v = regex2.ReplaceAllString(v, "")
	return v
}

// CompactString compact string representation of a value.
// If the value is a valid JSON, it will be compacted to a single line of JSON string.
// Otherwise, it removes all repeated spaces from the string representation of the value.
func CompactString(a any) string {
	bs := SimpleConvertToBytes(a)
	if IsJson(bs) {
		var buffer bytes.Buffer
		err := json.Compact(&buffer, bs)
		if IsNil(err) {
			return buffer.String()
		}
	}
	return CleanAllRepeatSpaces(string(bs))
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
