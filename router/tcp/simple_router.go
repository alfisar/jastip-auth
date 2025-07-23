package router

import (
	authpb "github.com/alfisar/jastip-import/proto/auth"

	"google.golang.org/grpc"
)

func simpleRoute(s *grpc.Server) {
	controll := SimpleInit()
	authpb.RegisterSimpleServer(s, controll)
}
