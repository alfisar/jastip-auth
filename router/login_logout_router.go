package router

import (
	"jastip/application/login/controller"
	"jastip/internal/handler"
	"jastip/internal/helper"
	"jastip/internal/middlewere"

	"github.com/gofiber/fiber/v2"
)

type loginLogoutRouter struct {
	Controller controller.LoginControllerContract
}

func NewLoginLogoutRouter(Controller controller.LoginControllerContract) *loginLogoutRouter {
	return &loginLogoutRouter{
		Controller: Controller,
	}
}

func (obj *loginLogoutRouter) loginLogoutRouters(v1 fiber.Router) {
	v1.Post("/login", middlewere.Validation(handler.HandlerLogin, helper.ValidationLogin), obj.Controller.Login)

}
