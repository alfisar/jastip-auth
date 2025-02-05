package service

import (
	"context"
	"jastip/config"
	"jastip/domain"
)

type ProfileServiceContract interface {
	Get(ctx context.Context, poolData *config.Config, userID int) (result domain.ProfileResponse, err domain.ErrorData)
	Update(ctx context.Context, poolData *config.Config, userId int, data map[string]any) (err domain.ErrorData)
}
