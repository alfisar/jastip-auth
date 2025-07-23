package http

import "github.com/gofiber/fiber/v2"

type NewProfileControllerContract interface {
	Get(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error

	GetAddrGrpc(ctx *fiber.Ctx) error
	GetAddress(ctx *fiber.Ctx) error
	GetAllAddress(ctx *fiber.Ctx) error
	SaveAddress(ctx *fiber.Ctx) error
}
