package service

import (
	"context"
	"encoding/json"
	"jastip/application/address/repository"

	"github.com/alfisar/jastip-import/domain"

	"github.com/alfisar/jastip-import/helpers/errorhandler"

	"gorm.io/gorm"
)

func getAllDataAddress(ctx context.Context, poolData *domain.Config, repo repository.AddressRepositoryContract, keys []string, value []any) (result []domain.AddressResponse, err domain.ErrorData) {
	where := map[string]any{}
	for i, v := range keys {
		where[v] = value[i]
	}

	results, errData := repo.GetAll(ctx, poolData.DBSql, where)
	if errData != nil {
		if errData.Error() == "get all address error : "+gorm.ErrRecordNotFound.Error() {
			err = errorhandler.ErrRecordNotFound()
		} else {
			err = errorhandler.ErrInternal(errorhandler.ErrCodeInternalServer, errData)
		}
		return
	}

	if len(results) == 0 {
		err = errorhandler.ErrRecordNotFound()
		return
	}

	result = results
	return
}

func getDataAddress(ctx context.Context, poolData *domain.Config, repo repository.AddressRepositoryContract, keys []string, value []any) (result domain.AddressResponse, err domain.ErrorData) {
	where := map[string]any{}
	for i, v := range keys {
		where[v] = value[i]
	}

	results, errData := repo.Get(ctx, poolData.DBSql, where)
	if errData != nil {
		if errData.Error() == "get address error : "+gorm.ErrRecordNotFound.Error() {
			err = errorhandler.ErrRecordNotFound()
		} else {
			err = errorhandler.ErrInternal(errorhandler.ErrCodeInternalServer, errData)
		}
		return
	}

	result = results
	return
}

func mapToStruct(data map[string]any) (address domain.AddressRequest, err domain.ErrorData) {
	jsonData, errData := json.Marshal(data)
	if errData != nil {
		err = errorhandler.ErrInternal(errorhandler.ErrCodeParsing, errData)
		return
	}

	errData = json.Unmarshal(jsonData, &address)
	if errData != nil {
		err = errorhandler.ErrInternal(errorhandler.ErrCodeParsing, errData)
		return
	}
	return
}
