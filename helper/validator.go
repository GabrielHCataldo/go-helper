package helper

import (
	"github.com/go-playground/validator/v10"
)

var customValidate *validator.Validate

// Validate persist custom validator.Validate
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
		_ = customValidate.RegisterValidation("mongodb", validateMongoDb)
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
	return IsFullName(fl.Field().String())
}

func validateBcrypt(fl validator.FieldLevel) bool {
	return IsBCrypt(fl.Field().Interface())
}

func validateBearer(fl validator.FieldLevel) bool {
	return IsBearer(fl.Field().Interface())
}

func validateBeforeNow(fl validator.FieldLevel) bool {
	return IsBeforeNow(fl.Field().Interface())
}

func validateBeforeToday(fl validator.FieldLevel) bool {
	return IsBeforeDateToday(fl.Field().Interface())
}

func validateAfterNow(fl validator.FieldLevel) bool {
	return IsAfterNow(fl.Field().Interface())
}

func validateAfterToday(fl validator.FieldLevel) bool {
	return IsAfterDateToday(fl.Field().Interface())
}

func validateToday(fl validator.FieldLevel) bool {
	return IsToday(fl.Field().Interface())
}

func validateNow(fl validator.FieldLevel) bool {
	return IsNow(fl.Field().Interface())
}

func validateFullNow(fl validator.FieldLevel) bool {
	return IsFullNow(fl.Field().Interface())
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

func validateMongoDb(fl validator.FieldLevel) bool {
	if IsSlice(fl.Field().Interface()) {
		for i := 0; i < fl.Field().Len(); i++ {
			fieldValueSlice := fl.Field().Index(i).Interface()
			if IsNotObjectId(fieldValueSlice) {
				return false
			}
		}
		return true
	}
	return IsObjectId(fl.Field().Interface())
}

func validateEnum(fl validator.FieldLevel) bool {
	if IsSlice(fl.Field().Interface()) {
		for i := 0; i < fl.Field().Len(); i++ {
			fieldValueSlice, ok := fl.Field().Index(i).Interface().(BaseEnum)
			if !ok || !fieldValueSlice.IsEnumValid() {
				return false
			}
		}
		return true
	}
	value, ok := fl.Field().Interface().(BaseEnum)
	return ok && value.IsEnumValid()
}
