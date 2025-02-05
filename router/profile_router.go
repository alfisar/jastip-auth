package router

import (
	"jastip/application/profile/controller"
	"jastip/internal/handler"
	"jastip/internal/helper"
	"jastip/internal/middlewere"

	"github.com/gofiber/fiber/v2"
)

type profileRouter struct {
	Controller controller.NewProfileControllerContract
}

func NewProfileRouter(Controller controller.NewProfileControllerContract) *profileRouter {
	return &profileRouter{
		Controller: Controller,
	}
}

func (obj *profileRouter) profileRouters(v1 fiber.Router) {
	middleweres := setMiddleware()
	v1.Get("/profile", middleweres.Authenticate, obj.Controller.Get)

	v1.Patch("/profile", middleweres.Authenticate, middlewere.Validation(handler.HandlerUpdateProfile, helper.ValidationUpdateProfile), obj.Controller.Update)

}
