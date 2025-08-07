package router

import (
	controller "jastip/application/profile/controller/http"

	"github.com/alfisar/jastip-import/helpers/middlewere"

	"github.com/alfisar/jastip-import/helpers/handler"
	"github.com/alfisar/jastip-import/helpers/helper"

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

	v1.Get("/profile/address", middleweres.Authenticate, obj.Controller.GetAllAddress)
	v1.Post("/profile/address", middleweres.Authenticate, middlewere.Validation(handler.HandlerpostAddress, helper.ValidationAddress), obj.Controller.SaveAddress)

	v1.Get("/profile/address/grpc/:id", middleweres.Authenticate, obj.Controller.GetAddrGrpc)
	v1.Get("/profile/address/:id", middleweres.Authenticate, obj.Controller.GetAddress)

	v1.Patch("/profile/address/:id", middleweres.Authenticate, middlewere.Validation(handler.HandlerpostAddress, helper.ValidationAddress), obj.Controller.UpdateAddress)

}
