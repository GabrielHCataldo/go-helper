package helper

import (
	"github.com/go-playground/validator/v10"
	"github.com/klassmann/cpfcnpj"
	"github.com/nyaruka/phonenumbers"
	"golang.org/x/crypto/bcrypt"
	"net"
	"net/url"
	"regexp"
	"strings"
)

// IsUrl If value is url return true, otherwise return false ex: "google.com" = true
func IsUrl(a any) bool {
	s := SimpleConvertToString(a)
	_, err := url.ParseRequestURI(s)
	return err == nil
}

// IsNotUrl If value is not url return true, otherwise return false ex: "google.com" = false
func IsNotUrl(a any) bool {
	return !IsUrl(a)
}

// IsPhoneNumber If value is phone number by region return true, otherwise return false
func IsPhoneNumber(a any, defaultRegion string) bool {
	s := SimpleConvertToString(a)
	num, _ := phonenumbers.Parse(s, defaultRegion)
	return num != nil && phonenumbers.IsValidNumber(num)
}

// IsNotPhoneNumber If value is not phone number by region return true, otherwise return false
func IsNotPhoneNumber(a any, defaultRegion string) bool {
	return !IsPhoneNumber(a, defaultRegion)
}

// IsEmail If value is email return true, otherwise return false
func IsEmail(a any) bool {
	s := SimpleConvertToString(a)
	validate := validator.New()
	err := validate.Var(s, "email,max=180")
	return err == nil
}

// IsNotEmail If value is not email return true, otherwise return false
func IsNotEmail(a any) bool {
	return !IsEmail(a)
}

// IsCpf If value is cpf return true, otherwise return false
func IsCpf(a any) bool {
	s := SimpleConvertToString(a)
	return cpfcnpj.ValidateCPF(s)
}

// IsNotCpf If value is not cpf return true, otherwise return false
func IsNotCpf(a any) bool {
	return !IsCpf(a)
}

// IsCnpj If value is cnpj return true, otherwise return false
func IsCnpj(a any) bool {
	s := SimpleConvertToString(a)
	return cpfcnpj.ValidateCNPJ(s)
}

// IsNotCnpj If value is not cnpj return true, otherwise return false
func IsNotCnpj(a any) bool {
	return !IsCnpj(a)
}

// IsCpfCnpj If value is cpf or cnpj return true, otherwise return false
func IsCpfCnpj(a any) bool {
	return IsCpf(a) || IsCnpj(a)
}

// IsNotCpfCnpj If value is not cpf or cnpj return true, otherwise return false
func IsNotCpfCnpj(a any) bool {
	return !IsCpfCnpj(a)
}

// IsPostalCode If value is postal code return true, otherwise return false
func IsPostalCode(a any) bool {
	s := SimpleConvertToString(a)
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

// IsNotPostalCode If value is not postal code return true, otherwise return false
func IsNotPostalCode(a any) bool {
	return !IsPostalCode(a)
}

// IsPostalCodePerCountry If value is postal code per country return true, otherwise return false
func IsPostalCodePerCountry(a any, countryIso string) bool {
	s := SimpleConvertToString(a)
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

// IsNotPostalCodePerCountry If value is not postal code per country return true, otherwise return false
func IsNotPostalCodePerCountry(a any, countryIso string) bool {
	return !IsPostalCodePerCountry(a, countryIso)
}

// IsBase64 If value is string base64 return true, otherwise return false
func IsBase64(a any) bool {
	s := SimpleConvertToString(a)
	regex := regexp.MustCompile(`^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{4}|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$`)
	return regex.MatchString(s)
}

// IsNotBase64 If value is not string base64 return true, otherwise return false
func IsNotBase64(a any) bool {
	return !IsBase64(a)
}

// IsBCrypt If value is string bcrypt return true, otherwise return false
func IsBCrypt(a any) bool {
	s := SimpleConvertToString(a)
	cost, err := bcrypt.Cost([]byte(s))
	return err == nil && cost == bcrypt.DefaultCost
}

// IsNotBCrypt If value is not string bcrypt return true, otherwise return false
func IsNotBCrypt(a any) bool {
	return !IsBCrypt(a)
}

// IsBearer If value is string bearer return true, otherwise return false
func IsBearer(a any) bool {
	const bearer = "Bearer "
	s := SimpleConvertToString(a)
	splitAuthorization := strings.Split(s, bearer)
	return len(splitAuthorization) != 2 && splitAuthorization[0] == bearer
}

// IsNotBearer If value is not string bearer return true, otherwise return false
func IsNotBearer(a any) bool {
	return !IsBearer(a)
}

// IsPrivateIp check value is private ip
func IsPrivateIp(a any) bool {
	s := SimpleConvertToString(a)
	var privateIPBlocks []*net.IPNet
	for _, cidr := range []string{
		"127.0.0.0/8",    // IPv4
		"10.0.0.0/8",     // RFC1918
		"172.16.0.0/12",  // RFC1918
		"192.168.0.0/16", // RFC1918
		"169.254.0.0/16", // RFC3927 link-local
		"::1/128",        // IPv6
		"fe80::/10",      // IPv6 link-local
		"fc00::/7",       // IPv6 unique local addr
	} {
		_, block, _ := net.ParseCIDR(cidr)
		if block != nil {
			privateIPBlocks = append(privateIPBlocks, block)
		}
	}
	ip := net.ParseIP(s)
	result := ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast()
	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			result = true
			break
		}
	}
	return result
}

// IsNotPrivateIp check value is private ip
func IsNotPrivateIp(a any) bool {
	return !IsPrivateIp(a)
}

// IsFullName If value contains first name and last name return true, otherwise return false
func IsFullName(a any) bool {
	s := SimpleConvertToString(a)
	regex := regexp.MustCompile(`^([a-zA-Z]{2,}\s[a-zA-Z]+'?-?[a-zA-Z]+\s?([a-zA-Z]+)?)`)
	return regex.MatchString(s)
}

// IsNotFullName If value not contains first name and last name return true, otherwise return false
func IsNotFullName(a any) bool {
	return !IsFullName(a)
}

// IsBirthDate If value time is before today return true, otherwise return false
func IsBirthDate(a any) bool {
	return IsBeforeDateToday(a)
}

// IsNotBirthDate If value time is before today return true, otherwise return false
func IsNotBirthDate(a any) bool {
	return !IsBirthDate(a)
}

// IsIOSDeviceId If value string is ios device id hex return true, otherwise return false
func IsIOSDeviceId(a any) bool {
	s := SimpleConvertToString(a)
	regex := regexp.MustCompile(`[A-F0-9]{8}-[A-F0-9]{4}-[A-F0-9]{4}-[A-F0-9]{4}-[A-F0-9]{12}`)
	return regex.MatchString(s)
}

// IsNotIOSDeviceId If value string is not ios device id hex return true, otherwise return false
func IsNotIOSDeviceId(a any) bool {
	return !IsIOSDeviceId(a)
}

// IsAndroidDeviceId If value string is android device id hex return true, otherwise return false
func IsAndroidDeviceId(a any) bool {
	s := SimpleConvertToString(a)
	regex := regexp.MustCompile(`[0-9a-fA-F]`)
	return regex.MatchString(s)
}

// IsNotAndroidDeviceId If value string is not android device id hex return true, otherwise return false
func IsNotAndroidDeviceId(a any) bool {
	return !IsAndroidDeviceId(a)
}

// IsMobilePlatform If value string is "android", "ios" or "iphone os" (independently we always count lowercase) return true, otherwise return false
func IsMobilePlatform(a any) bool {
	s := SimpleConvertToString(a)
	platform := strings.ToLower(s)
	return platform == "android" || platform == "ios" || platform == "iphone os"
}

// IsNotMobilePlatform If value string is not "android", "ios" or "iphone os" (independently we always count lowercase) return true, otherwise return false
func IsNotMobilePlatform(a any) bool {
	return !IsMobilePlatform(a)
}
