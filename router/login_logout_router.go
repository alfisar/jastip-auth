package router

import (
	"jastip/application/loginlogout/controller"

	"jastip/internal/middlewere"

	"github.com/alfisar/jastip-import/helpers/handler"
	"github.com/alfisar/jastip-import/helpers/helper"

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
	middleweres := setMiddleware()
	v1.Post("/login", middlewere.Validation(handler.HandlerLogin, helper.ValidationLogin), obj.Controller.Login)

	v1.Post("/logout", middleweres.Authenticate, obj.Controller.Logout)

}
