package tcp

import (
	"context"
	"fmt"
	"jastip/application/profile/service"

	"github.com/alfisar/jastip-import/domain"
	authpb "github.com/alfisar/jastip-import/proto/auth"
)

type ProfileGrpcController struct {
	authpb.UnimplementedProfileServer
	serv service.ProfileServiceContract
}

func NewProfileGrpcController(serv service.ProfileServiceContract) *ProfileGrpcController {
	return &ProfileGrpcController{
		serv: serv,
	}
}

func (c *ProfileGrpcController) InitPoolData() *domain.Config {
	poolData := domain.DataPool.Get().(*domain.Config)
	return poolData
}

func (c *ProfileGrpcController) AddressByID(ctx context.Context, data *authpb.RequestAddressByID) (result *authpb.ResponseAddressByID, err error) {
	poolData := c.InitPoolData()
	resultAddr, errs := c.serv.GetAddress(ctx, poolData, int(data.AdressID), int(data.UserID))
	if errs.Code != 0 {
		err = fmt.Errorf(errs.Message)
		return
	}

	result = &authpb.ResponseAddressByID{
		Id:            int32(resultAddr.Id),
		Province:      resultAddr.Province,
		Street:        resultAddr.Street,
		City:          resultAddr.City,
		District:      resultAddr.District,
		SubDistrict:   resultAddr.SUbDistrict,
		PostalCode:    resultAddr.PostalCode,
		ReceiverName:  resultAddr.ReceiverName,
		ReceiverPhone: resultAddr.ReceiverPhone,
	}

	return
}
