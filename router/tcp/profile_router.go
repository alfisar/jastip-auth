package router

import (
	controller "jastip/application/profile/controller/tcp"

	authpb "github.com/alfisar/jastip-import/proto/auth"

	"google.golang.org/grpc"
)

type profileGrpcRouter struct {
	Controller controller.ProfileGrpcController
}

func NewProfileGrpcRouter(Controller controller.ProfileGrpcController) *profileGrpcRouter {
	return &profileGrpcRouter{
		Controller: Controller,
	}
}

func (obj *profileGrpcRouter) profileGrpcRouters(s *grpc.Server) {
	authpb.RegisterProfileServer(s, &obj.Controller)
}
