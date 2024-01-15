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

func TestIsNotUrl(t *testing.T) {
	v := "string text"
	result := IsNotUrl(v)
	logger.Info("IsNotUrl:", result)
}

func TestIsPhoneNumber(t *testing.T) {
	v := "47997576130"
	result := IsPhoneNumber(v, "BR")
	logger.Info("IsPhoneNumber:", result)
}

func TestIsNotPhoneNumber(t *testing.T) {
	v := "47997576130"
	result := IsNotPhoneNumber(v, "BR")
	logger.Info("IsNotPhoneNumber:", result)
}

func TestIsEmail(t *testing.T) {
	v := "test@gmail.com"
	result := IsEmail(v)
	logger.Info("IsEmail:", result)
}

func TestIsNotEmail(t *testing.T) {
	v := "test@gmail.com"
	result := IsNotEmail(v)
	logger.Info("IsNotEmail:", result)
}

func TestIsCpf(t *testing.T) {
	v := "11664947051"
	result := IsCpf(v)
	logger.Info("IsCpf:", result)
}

func TestIsNotCpf(t *testing.T) {
	v := "11664947051"
	result := IsNotCpf(v)
	logger.Info("IsNotCpf:", result)
}

func TestIsNotCnpj(t *testing.T) {
	v := "52977110000101"
	result := IsNotCnpj(v)
	logger.Info("IsNotCnpj:", result)
}

func TestIsCpfCnpj(t *testing.T) {
	v := "11664947051"
	result := IsCpfCnpj(v)
	logger.Info("IsCpfCnpj:", result)
}

func TestIsNotCpfCnpj(t *testing.T) {
	v := "11664947051"
	result := IsNotCpfCnpj(v)
	logger.Info("IsNotCpf:", result)
}

func TestIsPostalCode(t *testing.T) {
	v := "89041-001"
	result := IsPostalCode(v)
	logger.Info("IsPostalCode:", result)
	v2 := "8904100132sas"
	result = IsPostalCode(v2)
	logger.Info("IsPostalCode:", result)
}

func TestIsNotPostalCode(t *testing.T) {
	v := "89041-001"
	result := IsNotPostalCode(v)
	logger.Info("IsNotPostalCode:", result)
}

func TestIsPostalCodePerCountry(t *testing.T) {
	v := "89041-001"
	result := IsPostalCodePerCountry(v, "BR")
	logger.Info("IsPostalCodePerCountry:", result)
	v2 := "8904100132sas"
	result = IsPostalCodePerCountry(v2, "MR")
	logger.Info("IsPostalCodePerCountry:", result)
}

func TestIsNotPostalCodePerCountry(t *testing.T) {
	v := "89041-001"
	result := IsNotPostalCodePerCountry(v, "BR")
	logger.Info("IsNotPostalCodePerCountry:", result)
}

func TestIsObjectId(t *testing.T) {
	v := ""
	result := IsObjectId(v)
	logger.Info("IsObjectId:", result)
}

func TestIsNotObjectId(t *testing.T) {
	v := ""
	result := IsNotObjectId(v)
	logger.Info("IsNotObjectId:", result)
}

func TestIsBase64(t *testing.T) {
	v := ""
	result := IsBase64(v)
	logger.Info("IsBase64:", result)
}

func TestIsNotBase64(t *testing.T) {
	v := ""
	result := IsNotBase64(v)
	logger.Info("IsNotBase64:", result)
}

func TestIsBCrypt(t *testing.T) {
	v := ""
	result := IsBCrypt(v)
	logger.Info("IsBCrypt:", result)
}

func TestIsNotBCrypt(t *testing.T) {
	v := ""
	result := IsNotBCrypt(v)
	logger.Info("IsNotBCrypt:", result)
}

func TestIsBearer(t *testing.T) {
	v := ""
	result := IsBearer(v)
	logger.Info("IsBearer:", result)
}

func TestIsNotBearer(t *testing.T) {
	v := ""
	result := IsNotBearer(v)
	logger.Info("IsNotBearer:", result)
}

func TestIsPrivateIp(t *testing.T) {
	v := "127.0.0.1"
	result := IsPrivateIp(v)
	logger.Info("IsPrivateIp:", result)
}

func TestIsNotPrivateIp(t *testing.T) {
	v := "127.0.0.1"
	result := IsNotPrivateIp(v)
	logger.Info("IsNotPrivateIp:", result)
}

func TestIsFullName(t *testing.T) {
	v := ""
	result := IsFullName(v)
	logger.Info("IsFullName:", result)
}

func TestIsNotFullName(t *testing.T) {
	v := ""
	result := IsNotFullName(v)
	logger.Info("IsNotFullName:", result)
}

func TestIsBirthDate(t *testing.T) {
	v := time.Date(1998, 1, 11, 0, 0, 0, 0, time.Local)
	result := IsBirthDate(v)
	logger.Info("IsBirthDate:", result)
}

func TestIsNotBirthDate(t *testing.T) {
	v := time.Date(1998, 1, 11, 0, 0, 0, 0, time.Local)
	result := IsNotBirthDate(v)
	logger.Info("IsNotBirthDate:", result)
}

func TestIsIOSDeviceId(t *testing.T) {
	v := ""
	result := IsIosDeviceId(v)
	logger.Info("IsIosDeviceId:", result)
}

func TestIsNotIOSDeviceId(t *testing.T) {
	v := ""
	result := IsNotIOSDeviceId(v)
	logger.Info("IsNotIOSDeviceId:", result)
}

func TestIsAndroidDeviceId(t *testing.T) {
	v := ""
	result := IsAndroidDeviceId(v)
	logger.Info("IsAndroidDeviceId:", result)
}

func TestIsNotAndroidDeviceId(t *testing.T) {
	v := ""
	result := IsNotAndroidDeviceId(v)
	logger.Info("IsNotAndroidDeviceId:", result)
}

func TestIsMobilePlatform(t *testing.T) {
	v := "ios"
	result := IsMobilePlatform(v)
	logger.Info("IsMobilePlatform:", result)
}

func TestIsNotMobilePlatform(t *testing.T) {
	v := "ios"
	result := IsNotMobilePlatform(v)
	logger.Info("IsNotMobilePlatform:", result)
}
