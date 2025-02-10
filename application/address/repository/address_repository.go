package repository

import (
	"fmt"

	"github.com/alfisar/jastip-import/domain"

	"github.com/alfisar/jastip-import/helpers/errorhandler"

	"gorm.io/gorm"
)

type addressRepository struct{}

func NewAddressRepository() *addressRepository {
	return &addressRepository{}
}

func (r *addressRepository) Insert(conn *gorm.DB, data domain.AddressRequest) (err error) {
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

	err = conn.Debug().Table("address").Create(&data).Error
	if err != nil {
		err = fmt.Errorf("insert address error : %w", err)
		return
	}

	return
}

func (r *addressRepository) Save(conn *gorm.DB, data domain.AddressRequest) (err error) {
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

	err = conn.Debug().Table("address").Save(&data).Error
	if err != nil {
		err = fmt.Errorf("save address error : %w", err)
		return
	}

	return
}

func (r *addressRepository) Get(conn *gorm.DB, where map[string]any) (result domain.AddressResponse, err error) {
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

	err = conn.Debug().Table("address").Where(where).First(&result).Error
	if err != nil {
		err = fmt.Errorf("get address error : %w", err)
		return
	}

	return
}

func (r *addressRepository) Update(conn *gorm.DB, where map[string]any, updates map[string]any) (err error) {
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

	data := conn.Debug().Table("address").Where(where).Updates(updates)

	if data.Error != nil {
		err = fmt.Errorf("Update user error : %w", err)

	} else if data.RowsAffected == 0 {
		err = fmt.Errorf("Update Failed : No Rows Affected")
	}
	return
}
