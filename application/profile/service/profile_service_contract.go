package service

import (
	"context"

	"github.com/alfisar/jastip-import/domain"
)

type ProfileServiceContract interface {
	Get(ctx context.Context, poolData *domain.Config, userID int) (result domain.ProfileResponse, err domain.ErrorData)
	Update(ctx context.Context, poolData *domain.Config, userId int, data map[string]any) (err domain.ErrorData)
	GetAddress(ctx context.Context, poolData *domain.Config, userId int) (result domain.AddressResponse, err domain.ErrorData)
	SaveAddress(ctx context.Context, poolData *domain.Config, userID int, data map[string]any) (err domain.ErrorData)
}
