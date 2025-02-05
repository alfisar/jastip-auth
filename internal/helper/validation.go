package helper

import (
	"fmt"
	"jastip/domain"
	validator "jastip/internal/validation"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidationDataUser(data domain.User) (err error) {
	err = validation.ValidateStruct(
		&data,
		validation.Field(&data.Email, validator.Required, validator.Email),
		validation.Field(&data.FullName, validator.Required, validator.AlphanumericSimbols),
		validation.Field(&data.NoHP, validator.Required, validator.Numeric),
		validation.Field(&data.Username, validator.Required, validator.AlphanumericSimbols),
		validation.Field(&data.Password, validator.Required, validator.AlphanumericSimbols),
	)
	return
}

func ValidationDataUserVerifyOTP(data domain.UserVerifyOtpRequest) (err error) {
	err = validation.ValidateStruct(
		&data,
		validation.Field(&data.Email, validator.Required, validator.Email),
		validation.Field(&data.NoHP, validator.Required, validator.Numeric),
	)
	return
}

func ValidationDataUserResendOTP(data domain.UserResendOtpRequest) (err error) {
	err = validation.ValidateStruct(
		&data,
		validation.Field(&data.Email, validator.Required, validator.Email),
		validation.Field(&data.NoHP, validator.Required, validator.Numeric),
	)
	return
}

func ValidationLogin(data domain.UserLoginRequest) (err error) {

	_, errs := strconv.Atoi(data.Username)
	if errs != nil {
		err = validation.ValidateStruct(
			&data,
			validation.Field(&data.Username, validator.Email),
			validation.Field(&data.Password, validator.Required, validator.AlphanumericSimbols),
		)
	} else {
		err = validation.ValidateStruct(
			&data,
			validation.Field(&data.Username, validator.Numeric),
			validation.Field(&data.Password, validator.Required, validator.AlphanumericSimbols),
		)
	}

	return
}

func filterRequestBody(data map[string]any, allowedKeys []string) map[string]any {
	filtered := make(map[string]any)
	for _, key := range allowedKeys {
		if val, exists := data[key]; exists {
			filtered[key] = val
		}
	}
	return filtered
}

func ValidationUpdateProfile(data map[string]any) (err error) {
	var (
		rules validation.Rule
	)
	allowedKeys := []string{
		"full_name",
		"email",
		"nohp",
		"password",
	}

	mappingData := filterRequestBody(data, allowedKeys)
	if len(mappingData) == 0 {
		err = fmt.Errorf("data cannot empty")
		return
	}
	for key, v := range mappingData {

		if key == "nohp" {
			rules = validator.Numeric
		} else if key == "email" {
			rules = validator.Email
		} else {
			rules = validator.AlphanumericSimbols
		}
		err = ValidateMappingData(v, key, err, rules)
	}

	return
}

func ValidateMappingData(data any, key string, errs error, rules validation.Rule) (err error) {

	err = validation.Validate(data, rules)
	if err != nil {
		err = fmt.Errorf(key + ": " + err.Error())
		if errs != nil {

			err = fmt.Errorf(errs.Error()[:len(errs.Error())-1] + "; " + err.Error())
		}
	} else {
		err = errs
	}

	return
}
