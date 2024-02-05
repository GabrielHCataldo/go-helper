package helper

import (
	"log"
	"testing"
)

func TestFormatPhoneNumber(t *testing.T) {
	s := "47997576130"
	result := FormatPhoneNumber(s, "BR")
	log.Println("FormatPhoneNumber:", result)
}

func TestFormatPhoneNumberNational(t *testing.T) {
	s := "47997576130"
	result := FormatPhoneNumberNational(s, "BR")
	log.Println("FormatPhoneNumberNational:", result)
}

func TestFormatCpf(t *testing.T) {
	s := "02104996643"
	result := FormatCpf(s)
	log.Println("FormatCpf:", result)
}

func TestFormatCnpj(t *testing.T) {
	s := "45991590000108"
	result := FormatCnpj(s)
	log.Println("FormatCnpj:", result)
}

func TestFormatFloat32(t *testing.T) {
	f := 1.23
	result := FormatFloat32(float32(f))
	log.Println("FormatFloat32:", result)
}

func TestFormatFloat64(t *testing.T) {
	f := 12.23
	result := FormatFloat64(f)
	log.Println("FormatFloat64:", result)
}

func TestFormatEFloat32(t *testing.T) {
	f := 1.23
	result := FormatEFloat32(float32(f))
	log.Println("FormatEFloat32:", result)
}

func TestFormatEFloat64(t *testing.T) {
	f := 12.23
	result := FormatEFloat64(f)
	log.Println("FormatEFloat64:", result)
}

func TestFormatMoney(t *testing.T) {
	f := 12.23
	result := FormatMoney(f, 2, "R$ ", ".", ",")
	log.Println("FormatMoney:", result)
}

func TestFormatMoneyBr(t *testing.T) {
	f := 12.23
	result := FormatMoneyBr(f)
	log.Println("FormatMoneyBr:", result)
}

func TestFormatPercentage(t *testing.T) {
	f := 12.23
	result := FormatPercentage(f, 2)
	log.Println("FormatPercentage:", result)
}

func TestHideCpf(t *testing.T) {
	s := "02104996643"
	result := HideCpf(s, true)
	log.Println("HideCpf:", result)
}

func TestHideCnpj(t *testing.T) {
	s := "45991590000108"
	result := HideCnpj(s, false)
	log.Println("HideCnpj:", result)
}

func TestHidePhoneNumber(t *testing.T) {
	s := "47997576130"
	result := HidePhoneNumber(s, "BR", true)
	log.Println("HidePhoneNumber:", result)
}

func TestHidePhoneNumberNational(t *testing.T) {
	s := "47997576130"
	result := HidePhoneNumberNational(s, "BR", false)
	log.Println("HidePhoneNumberNational:", result)
}

func TestHideEmail(t *testing.T) {
	s := "gabrielcataldo@gmail.com"
	result := HideEmail(s, false)
	log.Println("HideEmail:", result)
}

func TestMaskStartOrEndOfString(t *testing.T) {
	s := "g"
	result := MaskStartOrEndOfString(s, '*', true)
	log.Println("MaskStartOrEndOfString:", result)
}
