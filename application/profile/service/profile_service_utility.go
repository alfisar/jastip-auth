package service

import (
	"context"
	"encoding/json"
	"jastip/application/address/repository"

	"github.com/alfisar/jastip-import/domain"

	"github.com/alfisar/jastip-import/helpers/errorhandler"

	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
)

func getDataAddress(ctx context.Context, poolData *domain.Config, repo repository.AddressRepositoryContract, keys []string, value []any) (result domain.AddressResponse, err domain.ErrorData) {
	where := map[string]any{}
	for i, v := range keys {
		where[v] = value[i]
	}

	results, errData := repo.Get(poolData.DBSql, where)
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
	jsonData, errData := json.Marshal(data) // Convert map to JSON
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

func inserSaveAddress(ctx context.Context, poolData *domain.Config, userID int, repo repository.AddressRepositoryContract, data map[string]any) (err domain.ErrorData) {
	keys := []string{"user_id"}
	values := []any{userID}
	_, err = getDataAddress(ctx, poolData, repo, keys, values)
	if err.Code != 0 && err.HTTPCode == fasthttp.StatusInternalServerError {

		return
	}

	if err.Code != 0 {
		data["user_id"] = userID
		req, errs := mapToStruct(data)
		if errs.Code != 0 {
			err = errs
			return
		}

		errData := repo.Insert(poolData.DBSql, req)
		if errData != nil {
			err = errorhandler.ErrInsertData(errData)
			return
		}
	} else {
		where := map[string]any{
			"user_id": userID,
		}

		errData := repo.Update(poolData.DBSql, where, data)
		if errData != nil {
			err = errorhandler.ErrUpdateData(errData)
			return
		}
	}
	err = domain.ErrorData{}
	return
}
