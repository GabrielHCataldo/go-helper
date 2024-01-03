package helper

import (
	"github.com/go-playground/validator/v10"
	"github.com/klassmann/cpfcnpj"
	"github.com/nyaruka/phonenumbers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/url"
	"regexp"
	"strings"
)

func IsUrl(v any) bool {
	s := ConvertToString(v)
	_, err := url.ParseRequestURI(s)
	return err == nil
}

func IsPhoneNumber(v any, defaultRegion string) bool {
	s := ConvertToString(v)
	num, _ := phonenumbers.Parse(s, defaultRegion)
	return num != nil && phonenumbers.IsValidNumber(num)
}

func IsEmail(v any) bool {
	s := ConvertToString(v)
	validate := validator.New()
	err := validate.Var(s, "email,max=180")
	return err == nil
}

func IsCpf(v any) bool {
	s := ConvertToString(v)
	return cpfcnpj.ValidateCPF(s)
}

func IsCnpj(v any) bool {
	s := ConvertToString(v)
	return cpfcnpj.ValidateCNPJ(s)
}

func IsPostalCode(v any) bool {
	s := ConvertToString(v)
	var postalCodes []map[string]string
	_ = GetFileJson("../postal-codes.json", &postalCodes)
	matched := false
	for _, postalCode := range postalCodes {
		regexStr := postalCode["Regex"]
		if IsNotEmpty(regexStr) {
			regex := regexp.MustCompile(regexStr)
			if regex.MatchString(s) {
				matched = true
				break
			}
		}
	}
	return matched
}

func IsPostalCodePerCountry(v any, countryIso string) bool {
	s := ConvertToString(v)
	var postalCodes []map[string]string
	_ = GetFileJson("../postal-codes.json", &postalCodes)
	matched := false
	for _, postalCode := range postalCodes {
		if IsEmpty(countryIso) || postalCode["ISO"] == countryIso {
			regexStr := postalCode["Regex"]
			if IsEmpty(regexStr) {
				matched = true
			} else {
				regex := regexp.MustCompile(regexStr)
				matched = regex.MatchString(s)
			}
		}
	}
	return matched
}

func IsObjectId(v any) bool {
	s := ConvertToString(v)
	_, err := primitive.ObjectIDFromHex(s)
	return err == nil
}

func IsBase64(v any) bool {
	s := ConvertToString(v)
	regex := regexp.MustCompile(`^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$`)
	return regex.MatchString(s)
}

func IsBCrypt(v any) bool {
	s := ConvertToString(v)
	cost, err := bcrypt.Cost([]byte(s))
	return err == nil && cost == bcrypt.DefaultCost
}

func IsBearer(v any) bool {
	const bearer = "Bearer "
	s := ConvertToString(v)
	splitAuthorization := strings.Split(s, bearer)
	return len(splitAuthorization) != 2 && splitAuthorization[0] == bearer
}

func ValidateFullName(v any) bool {
	s := ConvertToString(v)
	regex := regexp.MustCompile(`^([a-zA-Z]{2,}\s[a-zA-Z]+'?-?[a-zA-Z]+\s?([a-zA-Z]+)?)`)
	return regex.MatchString(s)
}

func ValidateBirthDate(v any) bool {
	return IsBeforeDateToday(v)
}

func ValidateIosDeviceId(v any) bool {
	s := ConvertToString(v)
	regex := regexp.MustCompile(`[A-F0-9]{8}-[A-F0-9]{4}-[A-F0-9]{4}-[A-F0-9]{4}-[A-F0-9]{12}`)
	return regex.MatchString(s)
}

func ValidateAndroidDeviceId(v any) bool {
	s := ConvertToString(v)
	regex := regexp.MustCompile(`[0-9a-fA-F]`)
	return regex.MatchString(s)
}

func ValidateMobilePlatform(v any) bool {
	s := ConvertToString(v)
	platform := strings.ToLower(s)
	return platform == "android" || platform == "ios" || platform == "iphone os"
}
