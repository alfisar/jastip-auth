package router

import (
	"jastip/application/register/controller"
	"jastip/internal/handler"
	"jastip/internal/helper"
	"jastip/internal/middlewere"

	"github.com/gofiber/fiber/v2"
)

type regisRouter struct {
	Controller controller.RegisterControllerContract
}

func NewRegisRouter(Controller controller.RegisterControllerContract) *regisRouter {
	return &regisRouter{
		Controller: Controller,
	}
}

func (obj *regisRouter) regisRouters(v1 fiber.Router) {
	v1.Post("/registration", middlewere.Validation(handler.HandlerRegistration, helper.ValidationDataUser), obj.Controller.Register)

	v1.Post("/verify-otp", middlewere.Validation(handler.HandlerVerify, helper.ValidationDataUserVerifyOTP), obj.Controller.VerifyOTP)

	v1.Post("/resend-otp", middlewere.Validation(handler.HandlerResend, helper.ValidationDataUserResendOTP), obj.Controller.ResendOTP)

}
