package controller

import "github.com/gofiber/fiber/v2"

type LoginControllerContract interface {
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
}
