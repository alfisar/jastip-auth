package controller

import (
	"jastip/application/profile/service"
	"jastip/config"
	"jastip/internal/consts"
	"jastip/internal/response"

	"github.com/gofiber/fiber/v2"
)

type profileController struct {
	serv service.ProfileServiceContract
}

func NewProfileController(serv service.ProfileServiceContract) *profileController {
	return &profileController{
		serv: serv,
	}
}

func (c *profileController) InitPoolData() *config.Config {
	poolData := config.DataPool.Get().(*config.Config)
	return poolData
}

func (c *profileController) Get(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()
	userID := ctx.Locals("data").(float64)
	result, err := c.serv.Get(ctx.Context(), poolData, int(userID))
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccess(result, consts.SuccessGetData)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}

func (c *profileController) Update(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()
	userID := ctx.Locals("data").(float64)
	request := ctx.Locals("validatedData").(map[string]any)
	err := c.serv.Update(ctx.Context(), poolData, int(userID), request)
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccess(nil, consts.SuccessUpdateData)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}

func (c *profileController) GetAddress(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()
	userID := ctx.Locals("data").(float64)

	result, err := c.serv.GetAddress(ctx.Context(), poolData, int(userID))
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccess(result, consts.SuccessGetData)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}

func (c *profileController) SaveAddress(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()
	userID := ctx.Locals("data").(float64)
	request := ctx.Locals("validatedData").(map[string]any)
	err := c.serv.SaveAddress(ctx.Context(), poolData, int(userID), request)
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccess(nil, consts.SuccessUpdateData)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}
