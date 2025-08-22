package http

import (
	"jastip/application/profile/service"
	"strconv"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/consts"
	"github.com/alfisar/jastip-import/helpers/errorhandler"
	"github.com/alfisar/jastip-import/helpers/response"
	"github.com/valyala/fasthttp"
	"google.golang.org/grpc"

	pb "github.com/alfisar/jastip-import/proto/auth"
	"github.com/gofiber/fiber/v2"
)

type profileController struct {
	serv service.ProfileServiceContract
}

func NewProfileController(serv service.ProfileServiceContract) *profileController {
	return &profileController{
		serv: serv,
	}
}

func (c *profileController) InitPoolData() *domain.Config {
	poolData := domain.DataPool.Get().(*domain.Config)
	return poolData
}

func (c *profileController) Get(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()
	userID := ctx.Locals("data").(float64)
	result, err := c.serv.Get(ctx.Context(), poolData, int(userID))
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccess(result, consts.SuccessGetData)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}

func (c *profileController) Update(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()
	userID := ctx.Locals("data").(float64)
	request := ctx.Locals("validatedData").(map[string]any)
	err := c.serv.Update(ctx.Context(), poolData, int(userID), request)
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccess(nil, consts.SuccessUpdateData)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}

func (c *profileController) GetAddress(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()
	userID := ctx.Locals("data").(float64)
	id := ctx.Params("id")
	if id == "" {
		err := errorhandler.ErrValidation(nil)
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	dataId, errs := strconv.Atoi(id)
	if errs != nil {
		err := errorhandler.ErrValidation(errs)
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}
	result, err := c.serv.GetAddress(ctx.Context(), poolData, dataId, int(userID))
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccess(result, consts.SuccessGetData)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}

func (c *profileController) GetAllAddress(ctx *fiber.Ctx) error {
	poollData := c.InitPoolData()
	userID := ctx.Locals("data").(float64)

	result, err := c.serv.GetAllAddress(ctx.Context(), poollData, int(userID))
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccess(result, consts.SuccessGetData)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}

func (c *profileController) SaveAddress(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()
	userID := ctx.Locals("data").(float64)
	request := ctx.Locals("validatedData").(map[string]any)
	err := c.serv.SaveAddress(ctx.Context(), poolData, int(userID), request)
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccess(nil, consts.SuccessUpdateData)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}

func (c *profileController) UpdateAddress(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()
	userID := ctx.Locals("data").(float64)
	request := ctx.Locals("validatedData").(map[string]any)

	id := ctx.Params("id")
	if id == "" {
		err := errorhandler.ErrValidation(nil)
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	dataId, errs := strconv.Atoi(id)
	if errs != nil {
		err := errorhandler.ErrValidation(errs)
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	err := c.serv.UpdateAddress(ctx.Context(), poolData, dataId, int(userID), request)
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccess(nil, consts.SuccessUpdateData)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}

// Test Grpc
func (c *profileController) GetAddrGrpc(ctx *fiber.Ctx) error {
	userID := ctx.Locals("data").(float64)
	id := ctx.Params("id")

	dataId, errs := strconv.Atoi(id)
	if errs != nil {
		err := errorhandler.ErrValidation(errs)
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

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
	grpcClient := pb.NewProfileClient(conn)
	data := pb.RequestAddressByID{
		UserID:   int32(userID),
		AdressID: int32(dataId),
	}

	res, err := grpcClient.AddressByID(ctx.Context(), &data)
	if err != nil {
		return ctx.Status(500).SendString("gRPC error: " + err.Error())
	}

	result := domain.AddressResponse{
		Id:          int(res.Id),
		Province:    res.Province,
		Street:      res.Street,
		City:        res.City,
		District:    res.District,
		SUbDistrict: res.SubDistrict,
		PostalCode:  res.PostalCode,
	}

	resp := response.ResponseSuccess(result, consts.SuccessGetData)
	response.WriteResponse(ctx, resp, domain.ErrorData{}, fiber.StatusOK)
	return nil

}
