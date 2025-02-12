package service

import (
	"context"

	"github.com/alfisar/jastip-import/domain"
)

type RegisterServiceContract interface {
	Register(ctx context.Context, poolData *domain.Config, data domain.User) (result domain.User, err domain.ErrorData)
	VerifyOTP(ctx context.Context, poolData *domain.Config, email string, nohp string, otp string) (err domain.ErrorData)
	ResendOtp(ctx context.Context, poolData *domain.Config, email string, nohp string, fullName string) (err domain.ErrorData)
}
