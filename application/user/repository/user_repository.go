package repository

import (
	"context"
	"fmt"

	"github.com/alfisar/jastip-import/domain"

	"github.com/alfisar/jastip-import/helpers/errorhandler"

	"gorm.io/gorm"
)

type userRepository struct{}

func NewUserRpository() *userRepository {
	return &userRepository{}
}

func (r *userRepository) Create(ctx context.Context, conn *gorm.DB, data domain.User) (id int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	err = conn.WithContext(ctx).Debug().Table("users").Create(&data).Error
	if err != nil {
		err = fmt.Errorf("create users error : %w", err)
		return
	}

	id = data.Id
	return
}

func (r *userRepository) Get(ctx context.Context, conn *gorm.DB, where map[string]any) (data domain.User, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	err = conn.WithContext(ctx).Debug().Table("users").Where(where).First(&data).Error
	if err != nil {
		err = fmt.Errorf("get users error : %w", err)
		return
	}

	return
}

func (r *userRepository) Update(ctx context.Context, conn *gorm.DB, where map[string]any, updates map[string]any) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}
		return

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	data := conn.WithContext(ctx).Debug().Table("users").Where(where).Updates(updates)

	if data.Error != nil {
		err = fmt.Errorf("Update user error : %w", err)

	} else if data.RowsAffected == 0 {
		err = fmt.Errorf("Update Failed : No Rows Affected")
	}
	return
}
