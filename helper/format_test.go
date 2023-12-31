package helper

import (
	"github.com/GabrielHCataldo/go-logger/logger"
	"testing"
)

func TestFormatPhoneNumber(t *testing.T) {
	s := "47997576130"
	result := FormatPhoneNumber(s, "BR")
	logger.Info("FormatPhoneNumber:", result)
}

func TestFormatPhoneNumberNational(t *testing.T) {
	s := "47997576130"
	result := FormatPhoneNumberNational(s, "BR")
	logger.Info("FormatPhoneNumberNational:", result)
}

func TestFormatCpf(t *testing.T) {
	s := "02104996643"
	result := FormatCpf(s)
	logger.Info("FormatCpf:", result)
}

func TestFormatCnpj(t *testing.T) {
	s := "45991590000108"
	result := FormatCnpj(s)
	logger.Info("FormatCnpj:", result)
}

func TestFormatFloat32(t *testing.T) {
	f := 1.23
	result := FormatFloat32(float32(f))
	logger.Info("FormatFloat32:", result)
}

func TestFormatFloat64(t *testing.T) {
	f := 12.23
	result := FormatFloat64(f)
	logger.Info("FormatFloat64:", result)
}

func TestFormatEFloat32(t *testing.T) {
	f := 1.23
	result := FormatEFloat32(float32(f))
	logger.Info("FormatEFloat32:", result)
}

func TestFormatEFloat64(t *testing.T) {
	f := 12.23
	result := FormatEFloat64(f)
	logger.Info("FormatEFloat64:", result)
}

func TestFormatMoney(t *testing.T) {
	f := 12.23
	result := FormatMoney(f, 2, "R$ ", ".", ",")
	logger.Info("FormatMoney:", result)
}

func TestFormatPercentage(t *testing.T) {
	f := 12.23
	result := FormatPercentage(f, 2)
	logger.Info("FormatPercentage:", result)
}

func TestHideCpf(t *testing.T) {
	s := "02104996643"
	result := HideCpf(s, true)
	logger.Info("HideCpf:", result)
}

func TestHideCnpj(t *testing.T) {
	s := "45991590000108"
	result := HideCnpj(s, false)
	logger.Info("HideCnpj:", result)
}

func TestHidePhoneNumber(t *testing.T) {
	s := "47997576130"
	result := HidePhoneNumber(s, "BR", true)
	logger.Info("HidePhoneNumber:", result)
}

func TestHidePhoneNumberNational(t *testing.T) {
	s := "47997576130"
	result := HidePhoneNumberNational(s, "BR", false)
	logger.Info("HidePhoneNumberNational:", result)
}

func TestHideEmail(t *testing.T) {
	s := "gabrielcataldo@gmail.com"
	result := HideEmail(s, false)
	logger.Info("HideEmail:", result)
}

func TestMaskStartOrEndOfString(t *testing.T) {
	s := "g"
	result := MaskStartOrEndOfString(s, '*', true)
	logger.Info("MaskStartOrEndOfString:", result)
}
