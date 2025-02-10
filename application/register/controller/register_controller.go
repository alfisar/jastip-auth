package controller

import (
	"jastip/application/register/service"
	"jastip/config"

	"github.com/alfisar/jastip-import/domain"

	"github.com/alfisar/jastip-import/helpers/consts"
	"github.com/alfisar/jastip-import/helpers/response"

	"github.com/gofiber/fiber/v2"
)

type registerController struct {
	serv service.RegisterServiceContract
}

func NewRegisterController(serv service.RegisterServiceContract) *registerController {
	return &registerController{
		serv: serv,
	}
}

func (c *registerController) InitPoolData() *config.Config {
	poolData := config.DataPool.Get().(*config.Config)
	return poolData
}

func (c *registerController) Register(ctx *fiber.Ctx) error {
	err := domain.ErrorData{}

	poolData := c.InitPoolData()
	request := ctx.Locals("validatedData").(domain.User)

	request.Role = 1
	result, err := c.serv.Register(ctx.Context(), poolData, request)
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	respData := domain.UserResponse{
		Id:       result.Id,
		FullName: result.FullName,
		Username: result.Username,
		Status:   result.Status,
	}
	resp := response.ResponseSuccess(respData, consts.SuccessCreatedData)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}

func (c *registerController) ResendOTP(ctx *fiber.Ctx) error {
	err := domain.ErrorData{}

	poolData := c.InitPoolData()
	request := ctx.Locals("validatedData").(domain.UserResendOtpRequest)

	err = c.serv.ResendOtp(ctx.Context(), poolData, request.Email, request.NoHP, request.FullName)
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccess(nil, consts.SuccessCreatedData)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}

func (c *registerController) VerifyOTP(ctx *fiber.Ctx) error {
	err := domain.ErrorData{}

	poolData := c.InitPoolData()
	request := ctx.Locals("validatedData").(domain.UserVerifyOtpRequest)

	err = c.serv.VerifyOTP(ctx.Context(), poolData, request.Email, request.NoHP, request.Otp)
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccess(nil, consts.SuccessVerifyData)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}
