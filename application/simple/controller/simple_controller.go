package controller

import (
	"jastip/config"
	"jastip/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type SimpleController struct{}

func NewSimpleController() *SimpleController {
	return &SimpleController{}
}

func (c *SimpleController) Simple(ctx *fiber.Ctx) error {
	_ = config.DataPool
	ctx.Status(fasthttp.StatusOK).JSON(domain.ErrorData{
		Status:  "success",
		Code:    0,
		Message: "Welcome to API Justip.in version 1.0, enjoy and chersss :)",
	})
	return nil
}
