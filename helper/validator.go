package helper

import (
	"github.com/go-playground/validator/v10"
)

var customValidate *validator.Validate

func Validate() *validator.Validate {
	if customValidate == nil {
		customValidate = validator.New()
		_ = customValidate.RegisterValidation("enum", validateEnum)
		_ = customValidate.RegisterValidation("phone_us", validatePhoneUs)
		_ = customValidate.RegisterValidation("phone_br", validatePhoneBr)
		_ = customValidate.RegisterValidation("full_name", validateFullName)
		_ = customValidate.RegisterValidation("bcrypt", validateBcrypt)
		_ = customValidate.RegisterValidation("postal_code", validatePostalCode)
		_ = customValidate.RegisterValidation("bearer", validateBearer)
		_ = customValidate.RegisterValidation("before_now", validateBeforeNow)
		_ = customValidate.RegisterValidation("before_today", validateBeforeToday)
		_ = customValidate.RegisterValidation("after_now", validateAfterNow)
		_ = customValidate.RegisterValidation("after_today", validateAfterToday)
		_ = customValidate.RegisterValidation("today", validateToday)
		_ = customValidate.RegisterValidation("now", validateNow)
		_ = customValidate.RegisterValidation("full_now", validateFullNow)
		_ = customValidate.RegisterValidation("cpf", validateCpf)
		_ = customValidate.RegisterValidation("cnpj", validateCnpj)
		_ = customValidate.RegisterValidation("cpfcnpj", validateCpfCnpj)
	}
	return customValidate
}

func validatePhoneBr(fl validator.FieldLevel) bool {
	return IsPhoneNumber(fl.Field().String(), "BR")
}

func validatePhoneUs(fl validator.FieldLevel) bool {
	return IsPhoneNumber(fl.Field().String(), "US")
}

func validateFullName(fl validator.FieldLevel) bool {
	return ValidateFullName(fl.Field().String())
}

func validateBcrypt(fl validator.FieldLevel) bool {
	return IsBCrypt(fl.Field().Interface())
}

func validateBearer(fl validator.FieldLevel) bool {
	return IsBearer(fl.Field().Interface())
}

func validateBeforeNow(fl validator.FieldLevel) bool {
	timeValidate, err := ConvertToTime(fl.Field().Interface())
	return err == nil && IsBeforeNow(timeValidate)
}

func validateBeforeToday(fl validator.FieldLevel) bool {
	timeValidate, err := ConvertToTime(fl.Field().Interface())
	return err == nil && IsBeforeDateToday(timeValidate)
}

func validateAfterNow(fl validator.FieldLevel) bool {
	timeValidate, err := ConvertToTime(fl.Field().Interface())
	return err == nil && IsAfterNow(timeValidate)
}

func validateAfterToday(fl validator.FieldLevel) bool {
	timeValidate, err := ConvertToTime(fl.Field().Interface())
	return err == nil && IsAfterDateToday(timeValidate)
}

func validateToday(fl validator.FieldLevel) bool {
	timeValidate, err := ConvertToTime(fl.Field().Interface())
	return err == nil && IsToday(timeValidate)
}

func validateNow(fl validator.FieldLevel) bool {
	timeValidate, err := ConvertToTime(fl.Field().Interface())
	return err == nil && IsNow(timeValidate)
}

func validateFullNow(fl validator.FieldLevel) bool {
	timeValidate, err := ConvertToTime(fl.Field().Interface())
	return err == nil && IsFullNow(timeValidate)
}

func validateCpf(fl validator.FieldLevel) bool {
	return IsCpf(fl.Field().Interface())
}

func validateCnpj(fl validator.FieldLevel) bool {
	return IsCnpj(fl.Field().Interface())
}

func validateCpfCnpj(fl validator.FieldLevel) bool {
	a := fl.Field().Interface()
	return IsCpf(a) || IsCnpj(a)
}

func validatePostalCode(fl validator.FieldLevel) bool {
	return IsPostalCode(fl.Field().String())
}

func validateEnum(fl validator.FieldLevel) bool {
	value := fl.Field().Interface().(BaseEnum)
	return value.IsEnumValid()
}
