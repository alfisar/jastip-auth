package service

import (
	"context"
	"jastip/config"
	"jastip/domain"
)

type LoginServiceContract interface {
	Login(ctx context.Context, poolData *config.Config, data domain.UserLoginRequest) (token string, err domain.ErrorData)
}
