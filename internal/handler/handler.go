package handler

import (
	"jastip/domain"

	"github.com/gofiber/fiber/v2"
)

func HandlerRegistration(c *fiber.Ctx) (domain.User, error) {
	request := domain.User{}
	errData := c.BodyParser(&request)
	if errData != nil {
		return request, errData
	}

	return request, nil
}

func HandlerVerify(c *fiber.Ctx) (domain.UserVerifyOtpRequest, error) {
	request := domain.UserVerifyOtpRequest{}
	errData := c.BodyParser(&request)
	if errData != nil {
		return request, errData
	}

	return request, nil
}

func HandlerResend(c *fiber.Ctx) (domain.UserResendOtpRequest, error) {
	request := domain.UserResendOtpRequest{}
	errData := c.BodyParser(&request)
	if errData != nil {
		return request, errData
	}

	return request, nil
}
