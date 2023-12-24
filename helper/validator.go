package helper

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"reflect"
	"strings"
	"time"
)

var customValidate *validator.Validate

func Validate() *validator.Validate {
	if customValidate != nil {
		return customValidate
	}
	customValidate = validator.New()
	_ = customValidate.RegisterValidation("enum", validateEnum)
	_ = customValidate.RegisterValidation("phone_us", validatePhoneUs)
	_ = customValidate.RegisterValidation("phone_br", validatePhoneBr)
	_ = customValidate.RegisterValidation("full_name", validateFullName)
	_ = customValidate.RegisterValidation("bcrypt", validateBcrypt)
	_ = customValidate.RegisterValidation("postal_code", validatePostalCode)
	_ = customValidate.RegisterValidation("bearer", validateBearer)
	_ = customValidate.RegisterValidation("before_now", validateBirthDate) // todo -> desenvolver
	_ = customValidate.RegisterValidation("after_now", validateBirthDate)  // todo -> desenvolver
	_ = customValidate.RegisterValidation("date_now", validateBirthDate)   // todo -> desenvolver
	_ = customValidate.RegisterValidation("cpfcnpj", validateCpfCnpj)
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
	cost, err := bcrypt.Cost([]byte(fl.Field().String()))
	if err != nil {
		return false
	}
	return cost == bcrypt.DefaultCost
}

func validateBearer(fl validator.FieldLevel) bool {
	const bearer = "Bearer "
	splitAuthorization := strings.Split(fl.Field().String(), bearer)
	if len(splitAuthorization) != 2 {
		return false
	}
	return splitAuthorization[0] == bearer
}

func validateBirthDate(fl validator.FieldLevel) bool {
	var timeValidate time.Time
	if fl.Field().Kind() == reflect.String {
		t, err := time.Parse(time.RFC3339, fl.Field().String())
		if err != nil {
			return false
		}
		timeValidate = t
	} else {
		datetime, ok := fl.Field().Interface().(primitive.DateTime)
		if !ok {
			return false
		}
		timeValidate = datetime.Time()
	}
	return ValidateBirthDate(timeValidate)
}

func validateCpfCnpj(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	return IsCpf(s) || IsCnpj(s)
}

func validatePostalCode(fl validator.FieldLevel) bool {
	return IsPostalCode(fl.Field().String())
}

func validateEnum(fl validator.FieldLevel) bool {
	value := fl.Field().Interface().(BaseEnum)
	return value.IsEnumValid()
}
