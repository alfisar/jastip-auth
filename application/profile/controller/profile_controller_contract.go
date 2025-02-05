package controller

import "github.com/gofiber/fiber/v2"

type NewProfileControllerContract interface {
	Get(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}
