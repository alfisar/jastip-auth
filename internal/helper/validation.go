package helper

import (
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
