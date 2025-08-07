package repository

import (
	"context"

	"github.com/alfisar/jastip-import/domain"

	"gorm.io/gorm"
)

type UserContractRepository interface {
	Create(ctx context.Context, conn *gorm.DB, data domain.User) (id int, err error)
	Get(ctx context.Context, conn *gorm.DB, where map[string]any) (data domain.User, err error)
	Update(ctx context.Context, conn *gorm.DB, where map[string]any, updates map[string]any) (err error)
}
