package repository

import (
	"jastip/domain"

	"gorm.io/gorm"
)

type AddressRepositoryContract interface {
	Insert(conn *gorm.DB, data domain.AddressRequest) (err error)
	Save(conn *gorm.DB, data domain.AddressRequest) (err error)
	Get(conn *gorm.DB, where map[string]any) (result domain.AddressResponse, err error)
	Update(conn *gorm.DB, where map[string]any, updates map[string]any) (err error)
}
