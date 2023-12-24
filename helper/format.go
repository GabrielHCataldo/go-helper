package helper

import (
	"github.com/klassmann/cpfcnpj"
	"github.com/leekchan/accounting"
	"github.com/nyaruka/phonenumbers"
	"regexp"
	"strconv"
)

func FormatPhoneNumber(v string) string {
	num, err := phonenumbers.Parse(v, "")
	if err == nil {
		return phonenumbers.Format(num, phonenumbers.INTERNATIONAL)
	}
	return v
}

func FormatPhoneNumberNational(v, defaultRegion string) string {
	num, err := phonenumbers.Parse(v, defaultRegion)
	if err == nil {
		return phonenumbers.Format(num, phonenumbers.NATIONAL)
	}
	return v
}

func FormatCpf(v string) string {
	if IsCpf(v) {
		cpf := cpfcnpj.NewCPF(cpfcnpj.Clean(v))
		return cpf.String()
	}
	return v
}

func FormatCnpj(v string) string {
	if IsCnpj(v) {
		cnpj := cpfcnpj.NewCNPJ(cpfcnpj.Clean(v))
		return cnpj.String()
	}
	return v
}

func FormatFloat32(v float32) string {
	return strconv.FormatFloat(float64(v), 'f', -1, 32)
}

func FormatFloat64(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func FormatMoney(v float64, precision int, symbol, thousand, decimal string) string {
	ac := accounting.Accounting{Symbol: symbol, Precision: precision, Thousand: thousand, Decimal: decimal}
	return ac.FormatMoney(v)
}

func FormatPercentage(v float64, precision int) string {
	return accounting.FormatNumberFloat64(v, precision, ".", ",") + "%"
}

func HideCpf(v string) string {
	//021.049.966-43 to 021.049.***-*
	if IsCpf(v) {
		formatValue := FormatCpf(v)
		return formatValue[0:7] + ".***-**"
	}
	return v
}

func HideCnpj(v string) string {
	//45.991.590/0001-08 to 45.991.590/****-**
	if IsCpf(v) {
		formatValue := FormatCnpj(v)
		return formatValue[0:10] + "/****-**"
	}
	return v
}

func HidePhoneNumber(v string) string {
	//+55 47 99757-6130 to +55 47 99757-****
	if IsPhoneNumber(v) {
		formatValue := FormatPhoneNumber(v)
		return formatValue[0:12] + "-****"
	}
	return v
}

func HidePhoneNumberNational(v, defaultRegion string) string {
	//47 99757-6130 to 47 99757-****
	if IsPhoneNumber(v) {
		formatValue := FormatPhoneNumberNational(v, defaultRegion)
		return formatValue[0:8] + "-****"
	}
	return v
}

func HideEmail(v string) string {
	//gabrielcataldo@gmail.com to gabri***@gmail.com
	if IsEmail(v) {
		regex := regexp.MustCompile(`(\w{3})[\w.-]+@([\w.]+\w)`)
		return regex.ReplaceAllString(v, `$1***@$2`)
	}
	return v
}
