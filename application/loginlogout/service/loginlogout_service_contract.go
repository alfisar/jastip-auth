package service

import (
	"context"

	"github.com/alfisar/jastip-import/domain"
)

type LoginServiceContract interface {
	Login(ctx context.Context, poolData *domain.Config, data domain.UserLoginRequest) (token string, err domain.ErrorData)

	Logout(ctx context.Context, poolData *domain.Config, userID int) (err domain.ErrorData)
}
