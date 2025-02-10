package service

import (
	"context"
	"jastip/config"

	"github.com/alfisar/jastip-import/domain"
)

type RegisterServiceContract interface {
	Register(ctx context.Context, poolData *config.Config, data domain.User) (result domain.User, err domain.ErrorData)
	VerifyOTP(ctx context.Context, poolData *config.Config, email string, nohp string, otp string) (err domain.ErrorData)
	ResendOtp(ctx context.Context, poolData *config.Config, email string, nohp string, fullName string) (err domain.ErrorData)
}
