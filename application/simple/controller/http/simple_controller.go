package http

import (
	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/errorhandler"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/alfisar/jastip-import/proto/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type SimpleController struct{}

func NewSimpleController() *SimpleController {
	return &SimpleController{}
}

func (c *SimpleController) Simple(ctx *fiber.Ctx) error {
	_ = domain.DataPool
	ctx.Status(fasthttp.StatusOK).JSON(domain.ErrorData{
		Status:  "success",
		Code:    0,
		Message: "Welcome to API Justip.in version 1.0, enjoy and chersss :)",
	})
	return nil
}

func (c *SimpleController) HealthyGRPC(ctx *fiber.Ctx) error {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		ctx.Status(fasthttp.StatusInternalServerError).JSON(domain.ErrorData{
			Status:  "error",
			Code:    errorhandler.ErrCodeInternalServer,
			Message: "Cannot connect GRPC",
			Errors:  err.Error(),
		})
	}

	defer conn.Close()
	grpcClient := pb.NewSimpleClient(conn)

	res, err := grpcClient.CheckRunning(ctx.Context(), &emptypb.Empty{})
	if err != nil {
		return ctx.Status(500).SendString("gRPC error: " + err.Error())
	}
	ctx.Status(fasthttp.StatusOK).JSON(domain.ErrorData{
		Status:  "success",
		Code:    0,
		Message: res.Message,
	})
	return nil

}
