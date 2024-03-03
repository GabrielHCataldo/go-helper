package helper

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"testing"
	"time"
)

func TestIsUrl(t *testing.T) {
	v := "string text"
	result := IsUrl(v)
	log.Println("IsUrl:", result)
}

func TestIsNotUrl(t *testing.T) {
	v := "string text"
	result := IsNotUrl(v)
	log.Println("IsNotUrl:", result)
}

func TestIsPhoneNumber(t *testing.T) {
	v := "47997576130"
	result := IsPhoneNumber(v, "BR")
	log.Println("IsPhoneNumber:", result)
}

func TestIsNotPhoneNumber(t *testing.T) {
	v := "47997576130"
	result := IsNotPhoneNumber(v, "BR")
	log.Println("IsNotPhoneNumber:", result)
}

func TestIsEmail(t *testing.T) {
	v := "test@gmail.com"
	result := IsEmail(v)
	log.Println("IsEmail:", result)
}

func TestIsNotEmail(t *testing.T) {
	v := "test@gmail.com"
	result := IsNotEmail(v)
	log.Println("IsNotEmail:", result)
}

func TestIsCpf(t *testing.T) {
	v := "11664947051"
	result := IsCpf(v)
	log.Println("IsCpf:", result)
}

func TestIsNotCpf(t *testing.T) {
	v := "11664947051"
	result := IsNotCpf(v)
	log.Println("IsNotCpf:", result)
}

func TestIsNotCnpj(t *testing.T) {
	v := "52977110000101"
	result := IsNotCnpj(v)
	log.Println("IsNotCnpj:", result)
}

func TestIsCpfCnpj(t *testing.T) {
	v := "11664947051"
	result := IsCpfCnpj(v)
	log.Println("IsCpfCnpj:", result)
}

func TestIsNotCpfCnpj(t *testing.T) {
	v := "11664947051"
	result := IsNotCpfCnpj(v)
	log.Println("IsNotCpf:", result)
}

func TestIsPostalCode(t *testing.T) {
	v := "89041-001"
	result := IsPostalCode(v)
	log.Println("IsPostalCode:", result)
	v2 := "8904100132sas"
	result = IsPostalCode(v2)
	log.Println("IsPostalCode:", result)
}

func TestIsNotPostalCode(t *testing.T) {
	v := "89041-001"
	result := IsNotPostalCode(v)
	log.Println("IsNotPostalCode:", result)
}

func TestIsPostalCodePerCountry(t *testing.T) {
	v := "89041-001"
	result := IsPostalCodePerCountry(v, "BR")
	log.Println("IsPostalCodePerCountry:", result)
	v2 := "8904100132sas"
	result = IsPostalCodePerCountry(v2, "MR")
	log.Println("IsPostalCodePerCountry:", result)
}

func TestIsNotPostalCodePerCountry(t *testing.T) {
	v := "89041-001"
	result := IsNotPostalCodePerCountry(v, "BR")
	log.Println("IsNotPostalCodePerCountry:", result)
}

func TestIsObjectId(t *testing.T) {
	result := IsObjectId(primitive.NilObjectID)
	log.Println("IsObjectId:", result)
	result = IsObjectId([]uint8{12, 32, 34, 123})
	log.Println("IsObjectId:", result)
	result = IsObjectId("ashaushas")
	log.Println("IsObjectId:", result)
	result = IsObjectId(primitive.NewObjectID())
	log.Println("IsObjectId:", result)
	result = IsObjectId(primitive.NewObjectID().Hex())
	log.Println("IsObjectId:", result)
}

func TestIsNotObjectId(t *testing.T) {
	v := ""
	result := IsNotObjectId(v)
	log.Println("IsNotObjectId:", result)
}

func TestIsPrimitiveDateTime(t *testing.T) {
	result := IsPrimitiveDateTime(primitive.DateTime(0))
	log.Println("IsPrimitiveDateTime:", result)
}

func TestIsNotPrimitiveDateTime(t *testing.T) {
	result := IsNotPrimitiveDateTime("")
	log.Println("IsNotPrimitiveDateTime:", result)
}

func TestIsBase64(t *testing.T) {
	v := ""
	result := IsBase64(v)
	log.Println("IsBase64:", result)
}

func TestIsNotBase64(t *testing.T) {
	v := ""
	result := IsNotBase64(v)
	log.Println("IsNotBase64:", result)
}

func TestIsBCrypt(t *testing.T) {
	v := ""
	result := IsBCrypt(v)
	log.Println("IsBCrypt:", result)
}

func TestIsNotBCrypt(t *testing.T) {
	v := ""
	result := IsNotBCrypt(v)
	log.Println("IsNotBCrypt:", result)
}

func TestIsBearer(t *testing.T) {
	v := ""
	result := IsBearer(v)
	log.Println("IsBearer:", result)
}

func TestIsNotBearer(t *testing.T) {
	v := ""
	result := IsNotBearer(v)
	log.Println("IsNotBearer:", result)
}

func TestIsPrivateIp(t *testing.T) {
	v := "127.0.0.1"
	result := IsPrivateIp(v)
	log.Println("IsPrivateIp:", result)
}

func TestIsNotPrivateIp(t *testing.T) {
	v := "127.0.0.1"
	result := IsNotPrivateIp(v)
	log.Println("IsNotPrivateIp:", result)
}

func TestIsFullName(t *testing.T) {
	v := ""
	result := IsFullName(v)
	log.Println("IsFullName:", result)
}

func TestIsNotFullName(t *testing.T) {
	v := ""
	result := IsNotFullName(v)
	log.Println("IsNotFullName:", result)
}

func TestIsBirthDate(t *testing.T) {
	v := time.Date(1998, 1, 11, 0, 0, 0, 0, time.Local)
	result := IsBirthDate(v)
	log.Println("IsBirthDate:", result)
}

func TestIsNotBirthDate(t *testing.T) {
	v := time.Date(1998, 1, 11, 0, 0, 0, 0, time.Local)
	result := IsNotBirthDate(v)
	log.Println("IsNotBirthDate:", result)
}

func TestIsIOSDeviceId(t *testing.T) {
	v := ""
	result := IsIOSDeviceId(v)
	log.Println("IsIOSDeviceId:", result)
}

func TestIsNotIOSDeviceId(t *testing.T) {
	v := ""
	result := IsNotIOSDeviceId(v)
	log.Println("IsNotIOSDeviceId:", result)
}

func TestIsAndroidDeviceId(t *testing.T) {
	v := ""
	result := IsAndroidDeviceId(v)
	log.Println("IsAndroidDeviceId:", result)
}

func TestIsNotAndroidDeviceId(t *testing.T) {
	v := ""
	result := IsNotAndroidDeviceId(v)
	log.Println("IsNotAndroidDeviceId:", result)
}

func TestIsMobilePlatform(t *testing.T) {
	v := "ios"
	result := IsMobilePlatform(v)
	log.Println("IsMobilePlatform:", result)
}

func TestIsNotMobilePlatform(t *testing.T) {
	v := "ios"
	result := IsNotMobilePlatform(v)
	log.Println("IsNotMobilePlatform:", result)
}
