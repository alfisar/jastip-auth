package controller

import (
	"jastip/application/profile/service"
	"strconv"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/consts"
	"github.com/alfisar/jastip-import/helpers/errorhandler"
	"github.com/alfisar/jastip-import/helpers/response"

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

func (c *profileController) InitPoolData() *domain.Config {
	poolData := domain.DataPool.Get().(*domain.Config)
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
	id := ctx.Params("id")
	if id == "" {
		err := errorhandler.ErrValidation(nil)
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	dataId, errs := strconv.Atoi(id)
	if errs != nil {
		err := errorhandler.ErrValidation(errs)
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}
	result, err := c.serv.GetAddress(ctx.Context(), poolData, dataId, int(userID))
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccess(result, consts.SuccessGetData)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}

func (c *profileController) GetAllAddress(ctx *fiber.Ctx) error {
	poollData := c.InitPoolData()
	userID := ctx.Locals("data").(float64)

	result, err := c.serv.GetAllAddress(ctx.Context(), poollData, int(userID))
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
