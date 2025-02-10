package service

import (
	"context"
	"jastip/config"

	"github.com/alfisar/jastip-import/domain"
)

type LoginServiceContract interface {
	Login(ctx context.Context, poolData *config.Config, data domain.UserLoginRequest) (token string, err domain.ErrorData)

	Logout(ctx context.Context, poolData *config.Config, userID int) (err domain.ErrorData)
}
