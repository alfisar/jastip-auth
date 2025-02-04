package controller

import (
	"jastip/application/login/service"
	"jastip/config"
	"jastip/domain"
	"jastip/internal/consts"
	"jastip/internal/response"

	"github.com/gofiber/fiber/v2"
)

type loginController struct {
	serv service.LoginServiceContract
}

func NewLoginController(serv service.LoginServiceContract) *loginController {
	return &loginController{
		serv: serv,
	}
}

func (c *loginController) InitPoolData() *config.Config {
	poolData := config.DataPool.Get().(*config.Config)
	return poolData
}

func (c *loginController) Login(ctx *fiber.Ctx) error {
	var (
		err   domain.ErrorData
		token string
	)

	poolData := c.InitPoolData()
	request := ctx.Locals("validatedData").(domain.UserLoginRequest)

	token, err = c.serv.Login(ctx.Context(), poolData, request)
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccessWithToken(nil, consts.SuccessLogin, token)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}
