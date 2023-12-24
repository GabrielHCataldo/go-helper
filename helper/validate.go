package helper

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/klassmann/cpfcnpj"
	"github.com/nyaruka/phonenumbers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"
)

func IsUrl(v string) bool {
	_, err := url.ParseRequestURI(v)
	return err == nil
}

func IsPhoneNumber(v string) bool {
	num, err := phonenumbers.Parse(v, "")
	if err == nil {
		return phonenumbers.IsValidNumber(num)
	}
	return false
}

func IsPhoneNumberPerRegion(v, defaultRegion string) bool {
	num, err := phonenumbers.Parse(v, defaultRegion)
	if err == nil {
		return phonenumbers.IsValidNumberForRegion(num, defaultRegion)
	}
	return false
}

func IsEmail(v string) bool {
	validate := validator.New()
	if err := validate.Var(v, "email,max=180"); err == nil {
		return true
	}
	return false
}

func IsCpf(v string) bool {
	return cpfcnpj.ValidateCPF(cpfcnpj.Clean(v))
}

func IsCnpj(v string) bool {
	return cpfcnpj.ValidateCNPJ(cpfcnpj.Clean(v))
}

func IsPostalCode(v string) bool {
	if !IsString(v) {
		panic("value type is not string")
	}
	elem := reflect.ValueOf(v)
	if IsPointer(v) || IsInterface(v) {
		elem = elem.Elem()
	}
	s := elem.String()
	var postalCodes []map[string]string
	_ = GetFileJson("postal-codes.json", &postalCodes)
	for _, postalCode := range postalCodes {
		regexStr := postalCode["Regex"]
		if IsNotEmpty(regexStr) {
			regex := regexp.MustCompile(regexStr)
			if regex.MatchString(s) {
				return true
			}
		}
	}
	return false
}

func IsPostalCodePerCountry(v any, countryIso string) bool {
	if !IsString(v) {
		panic("value type is not string")
	}
	elem := reflect.ValueOf(v)
	if IsPointer(v) || IsInterface(v) {
		elem = elem.Elem()
	}
	s := elem.String()
	var postalCodes []map[string]string
	_ = GetFileJson("postal-codes.json", &postalCodes)
	for _, postalCode := range postalCodes {
		if postalCode["ISO"] == countryIso {
			regexStr := postalCode["Regex"]
			if IsEmpty(regexStr) {
				return true
			}
			regex := regexp.MustCompile(regexStr)
			return regex.MatchString(s)
		}
	}
	return false
}

func IsObjectID(ID *string) bool {
	if ID != nil {
		_, err := primitive.ObjectIDFromHex(*ID)
		return err == nil
	}
	return false
}

func IsBase64(v any) bool {
	if IsString(v) {
		elem := reflect.ValueOf(v)
		if IsPointer(v) || IsInterface(v) {
			elem = elem.Elem()
		}
		regex := regexp.MustCompile(`^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$`)
		return regex.MatchString(elem.String())
	}
	return false
}

func IsPrivateIP(v string) bool {
	var privateIPBlocks []*net.IPNet
	for _, cidr := range []string{
		"127.0.0.0/8",
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"169.254.0.0/16",
		"::1/128",
		"fe80::/10",
		"fc00::/7",
	} {
		_, block, err := net.ParseCIDR(cidr)
		if err != nil {
			panic(fmt.Errorf("parse error on %q: %v", cidr, err))
		}
		privateIPBlocks = append(privateIPBlocks, block)
	}

	ip := net.ParseIP(v)
	if ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() {
		return true
	}
	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}

func ValidateFullName(v string) bool {
	regex := regexp.MustCompile(`^([a-zA-Z]{2,}\s[a-zA-Z]+'?-?[a-zA-Z]+\s?([a-zA-Z]+)?)`)
	return regex.MatchString(v)
}

func ValidateBirthDate(values time.Time) bool {
	return time.Now().After(values)
}

func ValidateIOSDeviceID(v string) bool {
	regex := regexp.MustCompile(`[A-F0-9]{8}-[A-F0-9]{4}-[A-F0-9]{4}-[A-F0-9]{4}-[A-F0-9]{12}`)
	return regex.MatchString(v)
}

func ValidateAndroidDeviceID(v string) bool {
	regex := regexp.MustCompile(`[0-9a-fA-F]`)
	return regex.MatchString(v)
}

func ValidateMobilePlatform(v string) bool {
	platform := strings.ToLower(v)
	return platform == "android" || platform == "ios" || platform == "iphone os"
}
