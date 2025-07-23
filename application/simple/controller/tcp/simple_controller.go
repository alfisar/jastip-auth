package tcp

import (
	"context"

	authpb "github.com/alfisar/jastip-import/proto/auth"

	"google.golang.org/protobuf/types/known/emptypb"
)

type SimpleGrpcController struct {
	authpb.UnimplementedSimpleServer
}

func NewSimpleGrpcController() *SimpleGrpcController {
	return &SimpleGrpcController{}
}

func (c *SimpleGrpcController) CheckRunning(ctx context.Context, _ *emptypb.Empty) (*authpb.Health, error) {
	return &authpb.Health{
		Message: "Welcome to gRPC Auth Jastip.in version 1.0, enjoy and chersss :)",
	}, nil
}
