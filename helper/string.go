package helper

import (
	"encoding/json"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetFirstLastName(v string) string {
	split := strings.Split(v, " ")
	if len(split) == 0 {
		return v
	}
	firstName := split[0]
	lastName := split[len(split)-1]
	if firstName != lastName {
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

func RandomNumber(min, max int) string {
	return strconv.Itoa(rand.Intn(min-max+1) + min)
}

func ReplaceAllRepeatSpace(v string) string {
	regex := regexp.MustCompile(` +`)
	regex2 := regexp.MustCompile(`^\s+|\s+$|\s+[\r\t\f\v]`)
	v = regex.ReplaceAllString(v, " ")
	v = regex2.ReplaceAllString(v, "")
	return v
}
