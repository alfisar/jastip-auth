package controller

import "github.com/gofiber/fiber/v2"

type RegisterControllerContract interface {
	Register(ctx *fiber.Ctx) error
	ResendOTP(ctx *fiber.Ctx) error
	VerifyOTP(ctx *fiber.Ctx) error
}
