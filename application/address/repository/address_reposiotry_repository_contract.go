package repository

import (
	"context"

	"github.com/alfisar/jastip-import/domain"

	"gorm.io/gorm"
)

type AddressRepositoryContract interface {
	Insert(ctx context.Context, conn *gorm.DB, data domain.AddressRequest) (err error)
	Save(ctx context.Context, conn *gorm.DB, data domain.AddressRequest) (err error)
	Get(ctx context.Context, conn *gorm.DB, where map[string]any) (result domain.AddressResponse, err error)
	GetAll(ctx context.Context, conn *gorm.DB, where map[string]any) (result []domain.AddressResponse, err error)
	Update(ctx context.Context, conn *gorm.DB, where map[string]any, updates map[string]any) (err error)
}
