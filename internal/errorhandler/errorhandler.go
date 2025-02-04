package errorhandler

import (
	"jastip/domain"

	"github.com/valyala/fasthttp"
)

const (
	// code series 1xx for validation data

	// Error code for invalid input ex : email format not valid, type data not valid
	ErrCodeInvalidInput int = 101

	// Error code for data required ex: phone is required but phone is empty
	ErrCodeRequired int = 102

	// code series 2xx for error DB

	// Error code for invalid connection DB
	ErrCodeConnection int = 201

	// Error code for data is empty or not found
	ErrCodeDataNotFound int = 202

	// Error code for save data
	ErrCodeInsert int = 203

	// Error code for update data
	ErrCodeUpdate int = 204

	// Error code for delete data
	ErrCodeDelete int = 205

	// Error code for get data
	ErrCodeGet int = 206

	// Error code for data is empty or not found
	ErrCodeBlocked int = 207

	// code series 3xx for error auth

	// Error code for authenticate is not valid ex: token not valid
	ErrCodeInvalidAuth int = 301

	// Error code for authenticate is expired ex: token is expired
	ErrCodeExpSession int = 302

	// Error code for forbiddes akses auth ex: user not have akses
	ErrCodeForbiddenAccess int = 303

	// code series 4xx for error bisnis

	// Error code for invalid logic bisnis ex: user duplicate
	ErrCodeInvalidLogicBisnis int = 401

	// code series 5xx for error internal ex: error hashing

	// Error code for hashing data
	ErrCodeHashing int = 501

	// Error code for generate token
	ErrCodeGenerateToken int = 502

	// Error code for parsing data
	ErrCodeParsing int = 503

	// Error code for rabbitmq
	ErrCodeRabbitMQ int = 504

	// Error code for generating data
	ErrCodeGenerate int = 505

	// Error code for send email data
	ErrCodeSendEmail int = 506

	// Error code for panic
	ErrCodePanic int = 507

	// Error message for data exist
	ErrMsgDataExist string = "Data already exist"

	// Error message connection empty
	ErrMsgConnEmpty string = "connection is nil"

	// Error message for data Login
	ErrMsgLoginRequired string = "Username / Password Cannot Empty"

	// Error message for data Login not match
	ErrMsgLoginFailed string = "Pastikan Email / No Hp / Password Benar"

	// Error message for invalid token
	ErrMsgTokenInvalid string = "Token Tidak Valid"

	// Error message for data Login not match
	ErrMsgEmailNoHPUnique string = "Email / No HP sudah terdaftar"

	// Error message for data OTP not match
	ErrMsgOTPInvalid string = "Invalid OTP"

	// Error message for failed generate token
	ErrMsgFailedGenerateToken string = "Gagal mendapatkan token"
)

func ErrValidation(err error) (result domain.ErrorData) {
	result.Status = "error"
	result.Code = ErrCodeInvalidInput
	result.HTTPCode = fasthttp.StatusBadRequest
	result.Message = "Invalid data input"
	result.Errors = err.Error()
	return
}

func ErrLogin(err error) (result domain.ErrorData) {
	result.Status = "error"
	result.Code = ErrCodeInvalidInput
	result.HTTPCode = fasthttp.StatusBadRequest
	result.Message = ErrMsgLoginFailed
	if err != nil {
		result.Errors = err.Error()
	}
	return
}

func ErrRecordNotFound() (result domain.ErrorData) {
	result = domain.ErrorData{
		Status:   "error",
		HTTPCode: fasthttp.StatusBadRequest,
		Code:     ErrCodeDataNotFound,
		Message:  "Data not found",
	}

	return
}

func ErrBlocking() (result domain.ErrorData) {
	result = domain.ErrorData{
		Status:   "error",
		HTTPCode: fasthttp.StatusBadRequest,
		Code:     ErrCodeBlocked,
		Message:  "User blocked, please wait many times and try again",
	}

	return
}

func ErrGetData(err error) (result domain.ErrorData) {
	result = domain.ErrorData{
		Status:   "error",
		Code:     ErrCodeGet,
		HTTPCode: fasthttp.StatusBadRequest,
		Message:  "Failed get data",
		Errors:   err.Error(),
	}

	return
}

func ErrInsertData(err error) (result domain.ErrorData) {
	result = domain.ErrorData{
		Status:   "error",
		Code:     ErrCodeInsert,
		HTTPCode: fasthttp.StatusBadRequest,
		Message:  "Failed insert data",
		Errors:   err.Error(),
	}

	return
}

func ErrInvalidLogic(code int, message string, errorData string) (result domain.ErrorData) {
	result = domain.ErrorData{
		Status:   "error",
		Code:     code,
		Message:  message,
		HTTPCode: fasthttp.StatusBadRequest,
		Errors:   errorData,
	}

	return
}

func ErrHashing(err error) (result domain.ErrorData) {
	result = domain.ErrorData{
		Status:   "error",
		Code:     ErrCodeHashing,
		HTTPCode: fasthttp.StatusBadRequest,
		Message:  "Internal errors",
		Errors:   err.Error(),
	}

	return
}

func ErrInternal(code int, err error) (result domain.ErrorData) {
	result = domain.ErrorData{
		Status:   "error",
		Code:     code,
		HTTPCode: fasthttp.StatusInternalServerError,
		Message:  "Internal errors",
		Errors:   err.Error(),
	}

	return
}

func ErrMiddleware(code int, msg string, err error) (result domain.ErrorData) {
	result = domain.ErrorData{
		Status:   "error",
		Code:     code,
		HTTPCode: fasthttp.StatusBadRequest,
		Message:  msg,
		Errors:   err.Error(),
	}

	return
}
