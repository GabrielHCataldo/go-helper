package helper

import (
	"github.com/go-playground/validator/v10"
	"github.com/klassmann/cpfcnpj"
	"github.com/nyaruka/phonenumbers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net"
	"net/url"
	"regexp"
	"strings"
)

// IsUrl If value is url return true, otherwise return false ex: "google.com" = true
func IsUrl(a any) bool {
	s, _ := ConvertToString(a)
	_, err := url.ParseRequestURI(s)
	return err == nil
}

// IsPhoneNumber If value is phone number by region return true, otherwise return false
func IsPhoneNumber(a any, defaultRegion string) bool {
	s, _ := ConvertToString(a)
	num, _ := phonenumbers.Parse(s, defaultRegion)
	return num != nil && phonenumbers.IsValidNumber(num)
}

// IsEmail If value is email return true, otherwise return false
func IsEmail(a any) bool {
	s, _ := ConvertToString(a)
	validate := validator.New()
	err := validate.Var(s, "email,max=180")
	return err == nil
}

// IsCpf If value is cpf return true, otherwise return false
func IsCpf(a any) bool {
	s, _ := ConvertToString(a)
	return cpfcnpj.ValidateCPF(s)
}

// IsCnpj If value is cnpj return true, otherwise return false
func IsCnpj(a any) bool {
	s, _ := ConvertToString(a)
	return cpfcnpj.ValidateCNPJ(s)
}

// IsPostalCode If value is postal code return true, otherwise return false
func IsPostalCode(a any) bool {
	s, _ := ConvertToString(a)
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

// IsPostalCodePerCountry If value is postal code per country return true, otherwise return false
func IsPostalCodePerCountry(a any, countryIso string) bool {
	s, _ := ConvertToString(a)
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

// IsObjectId If value is hex objectId return true, otherwise return false
func IsObjectId(a any) bool {
	s, _ := ConvertToString(a)
	_, err := primitive.ObjectIDFromHex(s)
	return err == nil
}

// IsBase64 If value is string base64 return true, otherwise return false
func IsBase64(a any) bool {
	s, _ := ConvertToString(a)
	regex := regexp.MustCompile(`^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$`)
	return regex.MatchString(s)
}

// IsBCrypt If value is string bcrypt return true, otherwise return false
func IsBCrypt(a any) bool {
	s, _ := ConvertToString(a)
	cost, err := bcrypt.Cost([]byte(s))
	return err == nil && cost == bcrypt.DefaultCost
}

// IsBearer If value is string bearer return true, otherwise return false
func IsBearer(a any) bool {
	const bearer = "Bearer "
	s, _ := ConvertToString(a)
	splitAuthorization := strings.Split(s, bearer)
	return len(splitAuthorization) != 2 && splitAuthorization[0] == bearer
}

// IsPrivateIP check value is private ip
func IsPrivateIP(a any) bool {
	s, _ := ConvertToString(a)
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

// ValidateFullName If value contains first name and last name return true, otherwise return false
func ValidateFullName(a any) bool {
	s, _ := ConvertToString(a)
	regex := regexp.MustCompile(`^([a-zA-Z]{2,}\s[a-zA-Z]+'?-?[a-zA-Z]+\s?([a-zA-Z]+)?)`)
	return regex.MatchString(s)
}

// ValidateBirthDate If value time is before today return true, otherwise return false
func ValidateBirthDate(a any) bool {
	return IsBeforeDateToday(a)
}

// ValidateIosDeviceId If value string is ios device id hex return true, otherwise return false
func ValidateIosDeviceId(a any) bool {
	s, _ := ConvertToString(a)
	regex := regexp.MustCompile(`[A-F0-9]{8}-[A-F0-9]{4}-[A-F0-9]{4}-[A-F0-9]{4}-[A-F0-9]{12}`)
	return regex.MatchString(s)
}

// ValidateAndroidDeviceId If value string is android device id hex return true, otherwise return false
func ValidateAndroidDeviceId(a any) bool {
	s, _ := ConvertToString(a)
	regex := regexp.MustCompile(`[0-9a-fA-F]`)
	return regex.MatchString(s)
}

// ValidateMobilePlatform If value string is "android", "ios" or "iphone os" (independently we always count lowercase) return true, otherwise return false
func ValidateMobilePlatform(a any) bool {
	s, _ := ConvertToString(a)
	platform := strings.ToLower(s)
	return platform == "android" || platform == "ios" || platform == "iphone os"
}
