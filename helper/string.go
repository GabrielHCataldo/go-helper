package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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

// IsHttpMethod checks whether the given value represents a valid HTTP method.
// It converts the value to a string and compares it with predefined HTTP method constants.
// Returns true if the value is a valid HTTP method, otherwise returns false.
func IsHttpMethod(a any) bool {
	method := SimpleConvertToString(a)
	switch method {
	case http.MethodGet, http.MethodPost, http.MethodHead, http.MethodPut, http.MethodDelete, http.MethodConnect,
		http.MethodOptions, http.MethodTrace, http.MethodPatch:
		return true
	}
	return false
}

// IsNotHttpMethod checks whether the given value is not a valid HTTP method.
// It negates the result of the IsHttpMethod function.
// Returns true if the value is not a valid HTTP method, otherwise returns false.
func IsNotHttpMethod(a any) bool {
	return !IsHttpMethod(a)
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
	bs, err := os.ReadFile(uri)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

// GetFileJson get file by uri string param, fill dest by file json
func GetFileJson(uri string, dest any) error {
	bs, err := os.ReadFile(uri)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, dest)
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

func CamelCaseString(a any) string {
	s := SimpleConvertToString(a)

	wordBoundary := regexp.MustCompile("[^a-zA-Z0-9]+")
	words := wordBoundary.Split(s, -1)

	for i := 0; i < len(words); i++ {
		if words[i] == "" {
			continue
		}
		if i == 0 {
			words[i] = strings.ToLower(words[i])
		} else {
			words[i] = strings.Title(words[i])
		}
	}

	return strings.Join(words, "")
}

func SnakeCaseString(a any) string {
	s := SimpleConvertToString(a)

	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(s, "${1}_$2")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	//remove all non-alphanumeric character except underscores
	snake = regexp.MustCompile("[^a-zA-Z0-9_]+").ReplaceAllString(snake, "_")

	return strings.ToLower(snake)
}

func KebabCaseString(a any) string {
	s := SimpleConvertToString(a)

	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")

	kebab := matchFirstCap.ReplaceAllString(s, "${1}-$2")
	kebab = matchAllCap.ReplaceAllString(kebab, "${1}-${2}")
	kebab = regexp.MustCompile("[^a-zA-Z0-9-]+").ReplaceAllString(kebab, "-")

	return strings.ToLower(kebab)
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
