package helper

import (
	"github.com/klassmann/cpfcnpj"
	"github.com/leekchan/accounting"
	"github.com/nyaruka/phonenumbers"
	"strconv"
)

// FormatPhoneNumber format string to international phone number ex: 47997576130 -> +55 47 99757-6130
func FormatPhoneNumber(v, defaultRegion string) string {
	num, err := phonenumbers.Parse(v, defaultRegion)
	if err == nil {
		return phonenumbers.Format(num, phonenumbers.INTERNATIONAL)
	}
	return v
}

// FormatPhoneNumberNational format string to national phone number ex: 47997576130 -> 47 99757-6130
func FormatPhoneNumberNational(v, defaultRegion string) string {
	num, err := phonenumbers.Parse(v, defaultRegion)
	if err == nil {
		return phonenumbers.Format(num, phonenumbers.NATIONAL)
	}
	return v
}

// FormatCpf format string to cpf if valid ex: 02104996642 -> 021.049.966-43
func FormatCpf(v string) string {
	if IsCpf(v) {
		cpf := cpfcnpj.NewCPF(cpfcnpj.Clean(v))
		return cpf.String()
	}
	return v
}

// FormatCnpj format string to cnpj if valid ex: 45991590000108 -> 45.991.590/0001-08
func FormatCnpj(v string) string {
	if IsCnpj(v) {
		cnpj := cpfcnpj.NewCNPJ(cpfcnpj.Clean(v))
		v = cnpj.String()
	}
	return v
}

// FormatFloat32 format float32 to string ex: 3.1415926535 -> "3.1415926535"
func FormatFloat32(v float32) string {
	return strconv.FormatFloat(float64(v), 'f', -1, 32)
}

// FormatFloat64 format float64 to string ex: 3.1415926535 -> "3.1415926535"
func FormatFloat64(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

// FormatEFloat32 format float64 to string ex: 3.1415926535 -> "3.1415927E+00"
func FormatEFloat32(v float32) string {
	return strconv.FormatFloat(float64(v), 'E', -1, 32)
}

// FormatEFloat64 format float64 to string ex: 3.1415926535 -> "3.1415926535E+00"
func FormatEFloat64(v float64) string {
	return strconv.FormatFloat(v, 'E', -1, 64)
}

// FormatMoney format float to string money
//
// Usage:
//
// result := FormatMoney(12.23, 2, "R$" , ".", "," )
//
// logger.Info("money result:", result)
//
// Output:
//
// [INFO 2023/12/24 08:26:38] money result: R$ 12,23
func FormatMoney(v float64, precision int, symbol, thousand, decimal string) string {
	ac := accounting.Accounting{Symbol: symbol, Precision: precision, Thousand: thousand, Decimal: decimal}
	return ac.FormatMoney(v)
}

// FormatPercentage format float to string percentage ex: 12.23 -> "12,23%"
func FormatPercentage(v float64, precision int) string {
	return accounting.FormatNumberFloat64(v, precision, ".", ",") + "%"
}

// HideCpf formats and partially hides the value ex: 02104996643 to 021.049.***-*
func HideCpf(v string, maskStart bool) string {
	if IsCpf(v) {
		formatValue := FormatCpf(v)
		v = formatValue
	}
	return MaskStartOrEndOfString(v, '*', maskStart)
}

// HideCnpj formats and partially hides the value ex: 45991590000108 to 45.991.590/****-**
func HideCnpj(v string, maskStart bool) string {
	if IsCpf(v) {
		formatValue := FormatCnpj(v)
		v = formatValue
	}
	return MaskStartOrEndOfString(v, '*', maskStart)
}

// HidePhoneNumber international format of the phone number and partially hide the value if it is a valid phone number
// ex: 47997576130 to +55 47 99757-****
func HidePhoneNumber(v, defaultRegion string, maskStart bool) string {
	if IsPhoneNumber(v, defaultRegion) {
		formatValue := FormatPhoneNumber(v, defaultRegion)
		v = formatValue
	}
	return MaskStartOrEndOfString(v, '*', maskStart)
}

// HidePhoneNumberNational national format of the phone number and partially hide the value if it is a valid phone number
// ex: 47997576130 to 47 99757-****
func HidePhoneNumberNational(v, defaultRegion string, maskStart bool) string {
	if IsPhoneNumber(v, defaultRegion) {
		formatValue := FormatPhoneNumberNational(v, defaultRegion)
		v = formatValue
	}
	return MaskStartOrEndOfString(v, '*', maskStart)
}

// HideEmail partially hide the value email if it is a valid
// if maskStart param is true
//
// ex: gabrielcataldo@gmail.com to ***********do@gmail.com
//
// if masStart param is false (default)
//
// ex: gabrielcataldo@gmail.com to gabrielcatal**@*****.***
func HideEmail(v string, maskStart bool) string {
	return MaskStartOrEndOfString(v, '*', maskStart)
}

// MaskStartOrEndOfString mask start or end of string
//
// # Parameters
//
// - s: string value
// - mask: character to replace
// - maskStart: if true mask start value, otherwise mask end value
func MaskStartOrEndOfString(s string, mask rune, maskStart bool) string {
	if len(s) == 1 {
		return string(mask)
	}
	maskedString := make([]rune, len(s))
	copy(maskedString, []rune(s))
	maskIndex := len(s) / 2
	if maskStart {
		for i := 0; i < maskIndex; i++ {
			si := string(s[i])
			if IsNumeric(si) || IsLetter(si) {
				maskedString[i] = mask
			} else {
				maskedString[i] = rune(s[i])
				maskIndex = MaxInt(maskIndex+1, len(s))
			}
		}
	} else {
		for i := maskIndex; i < len(s); i++ {
			si := string(s[i])
			if IsNumeric(si) || IsLetter(si) {
				maskedString[i] = mask
			} else {
				maskedString[i] = rune(s[i])
				maskIndex = MaxInt(maskIndex+1, len(s))
			}
		}
	}
	return string(maskedString)
}
