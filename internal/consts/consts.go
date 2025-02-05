package consts

import "time"

const (
	FailedValidation                = 1001
	RegexAlphanumericSimbols string = `^[a-zA-Z]?[a-zA-Z0-9\-\.\,\~\@\!\#\%\&\^\*\$(\)\/\s]+$`
	RegexAlphanumericPetik   string = `^[a-zA-Z0-9'\s]*[a-zA-Z][a-zA-Z0-9'\s]*$`

	AlphanumericSimbols      string = "Kolom harus dengan huruf, angka dan spasi."
	AlphanumericPetik        string = "Kolom hanya bisa dengan huruf, angka, spasi, dan petik serta harus minimal 1 huruf"
	AlphanumericSimbolsLogin string = "Password harus dengan huruf, angka, simbols, dan minimal 12 karakter"
	Digit                    string = "Kolom harus dengan angka."
	IsEmail                  string = "kolom harus sesuai kaidah email"
	MaxMinChar17             string = "Panjang data harus 17 character."
	RequiredField            string = "Kolom harus di isi"

	// config const redis db
	RedisToken string = "token"
	RedisOTP   string = "otp"

	// config response message
	SuccessRegister    string = "successfully register"
	SuccessCreatedData string = "successfully created data"
	SuccessGetData     string = "successfully get data"
	SuccessUpdateData  string = "successfully updated data"
	SuccessVerifyData  string = "successfully verify data"

	SuccessLogin  string = "successfully login"
	SuccessLogout string = "successfully logout"

	// config const name db
	TaskReminder string = "task_reminder"

	// config expired time
	RedisOTPExp time.Duration = 5 * time.Minute
	TokenExp    time.Duration = 15 * time.Minute

	// configuration flag
	Attemp    string = "Attemp_"
	AttempOTP int    = 5
)
