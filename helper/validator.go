package helper

import (
	"github.com/go-playground/validator/v10"
)

var customValidate *validator.Validate

// Validate persist custom validator.Validate
func Validate() *validator.Validate {
	if customValidate != nil {
		return customValidate
	}

	customValidate = validator.New()
	_ = customValidate.RegisterValidation("http_method", validateHttpMethod)
	_ = customValidate.RegisterValidation("url_path", validateUrlPath)
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
	_ = customValidate.RegisterValidation("duration", validateDuration)
	_ = customValidate.RegisterValidation("byte_unit", validateByteUnit)
	_ = customValidate.RegisterValidation("megabyte_unit", validateMegabyteUnit)
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
	return IsObjectIdType(fl.Field().Interface())
}

func validateDuration(fl validator.FieldLevel) bool {
	return IsTimeDurationType(fl.Field().Interface())
}

func validateByteUnit(fl validator.FieldLevel) bool {
	return IsByteUnit(fl.Field().Interface())
}

func validateMegabyteUnit(fl validator.FieldLevel) bool {
	return IsMegabyteUnit(fl.Field().Interface())
}

func validateEnum(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(BaseEnum)
	return ok && value.IsEnumValid()
}

func validateUrlPath(fl validator.FieldLevel) bool {
	return IsUrlPath(fl.Field().Interface())
}

func validateHttpMethod(fl validator.FieldLevel) bool {
	return IsHttpMethod(fl.Field().Interface())
}
