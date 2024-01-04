package helper

import (
	"github.com/GabrielHCataldo/go-logger/logger"
	"testing"
	"time"
)

func TestIsUrl(t *testing.T) {
	v := "string text"
	result := IsUrl(v)
	logger.Info("IsUrl:", result)
}

func TestIsPhoneNumber(t *testing.T) {
	v := "47997576130"
	result := IsPhoneNumber(v, "BR")
	logger.Info("IsPhoneNumber:", result)
}

func TestIsEmail(t *testing.T) {
	v := "test@gmail.com"
	result := IsEmail(v)
	logger.Info("IsEmail:", result)
}

func TestIsCpf(t *testing.T) {
	v := "11664947051"
	result := IsCpf(v)
	logger.Info("IsCpf:", result)
}

func TestIsCnpj(t *testing.T) {
	v := "52977110000101"
	result := IsCnpj(v)
	logger.Info("IsCnpj:", result)
}

func TestIsPostalCode(t *testing.T) {
	v := "89041-001"
	result := IsPostalCode(v)
	logger.Info("IsPostalCode:", result)
	v2 := "8904100132sas"
	result = IsPostalCode(v2)
	logger.Info("IsPostalCode:", result)
}

func TestIsPostalCodePerCountry(t *testing.T) {
	v := "89041-001"
	result := IsPostalCodePerCountry(v, "BR")
	logger.Info("IsPostalCodePerCountry:", result)
	v2 := "8904100132sas"
	result = IsPostalCodePerCountry(v2, "MR")
	logger.Info("IsPostalCodePerCountry:", result)
}

func TestIsObjectIs(t *testing.T) {
	v := ""
	result := IsObjectId(v)
	logger.Info("IsObjectId:", result)
}

func TestIsBase64(t *testing.T) {
	v := ""
	result := IsBase64(v)
	logger.Info("IsBase64:", result)
}

func TestIsBCrypt(t *testing.T) {
	v := ""
	result := IsBCrypt(v)
	logger.Info("IsBCrypt:", result)
}

func TestIsBearer(t *testing.T) {
	v := ""
	result := IsBearer(v)
	logger.Info("IsBearer:", result)
}

func TestIsPrivateIP(t *testing.T) {
	v := "127.0.0.1"
	result := IsPrivateIP(v)
	logger.Info("IsPrivateIP:", result)
}

func TestValidateFullName(t *testing.T) {
	v := ""
	result := ValidateFullName(v)
	logger.Info("ValidateFullName:", result)
}

func TestValidateBirthDate(t *testing.T) {
	v := time.Date(1998, 1, 11, 0, 0, 0, 0, time.Local)
	result := ValidateBirthDate(v)
	logger.Info("ValidateBirthDate:", result)
}

func TestValidateIosDeviceId(t *testing.T) {
	v := ""
	result := ValidateIosDeviceId(v)
	logger.Info("ValidateIosDeviceId:", result)
}

func TestValidateAndroidDeviceId(t *testing.T) {
	v := ""
	result := ValidateAndroidDeviceId(v)
	logger.Info("ValidateAndroidDeviceId:", result)
}

func TestValidateMobilePlatform(t *testing.T) {
	v := "ios"
	result := ValidateMobilePlatform(v)
	logger.Info("ValidateMobilePlatform:", result)
}
