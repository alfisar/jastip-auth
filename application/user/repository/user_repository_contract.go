package repository

import (
	"github.com/alfisar/jastip-import/domain"

	"gorm.io/gorm"
)

type UserContractRepository interface {
	Create(conn *gorm.DB, data domain.User) (id int, err error)
	Get(conn *gorm.DB, where map[string]any) (data domain.User, err error)
	Update(conn *gorm.DB, where map[string]any, updates map[string]any) (err error)
}
