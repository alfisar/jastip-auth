package service

import (
	"errors"
	"fmt"
	"jastip/application/user/repository"
	"jastip/config"
	"jastip/domain"
	"jastip/internal/errorhandler"
	"jastip/internal/handler"
	"jastip/internal/helper"
	"log"
)

func validateUser(poolData *config.Config, repo repository.UserContractRepository, data domain.UserLoginRequest) (err domain.ErrorData) {
	defer handler.PanicError()
	result, errs := getUserByEmailOrHP(poolData, repo, data)
	if errs.Code != 0 {
		err = errs
		return
	}

	errData := helper.Verify(result.Password, data.Password)
	if errData != nil {
		message := fmt.Sprintf("Invalid verify pass on func validate user : %s", errData.Error())
		log.Println(message)

		return errorhandler.ErrInvalidLogic(errorhandler.ErrCodeInvalidInput, errorhandler.ErrMsgLoginFailed, errData.Error())
	}

	return
}
func getUserByEmailOrHP(poolData *config.Config, repo repository.UserContractRepository, data domain.UserLoginRequest) (result domain.User, err domain.ErrorData) {
	defer handler.PanicError()
	errData := errors.New("")

	where := map[string]any{}
	if data.Email != "" {
		where = map[string]any{
			"email": data.Email,
		}
	} else if data.NoHP != "" {
		where = map[string]any{
			"nohp": data.NoHP,
		}
	}

	result, errData = repo.Get(poolData.DBSql, where)
	if errData != nil && errData.Error() != "get users error : record not found" {
		return result, errorhandler.ErrGetData(errData)

	}

	return
}
