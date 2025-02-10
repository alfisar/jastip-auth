package service

import (
	"context"
	"fmt"
	repoRedis "jastip/application/redis/repository"
	"jastip/application/user/repository"
	"jastip/config"
	"log"

	"github.com/alfisar/jastip-import/domain"

	"jastip/internal/handler"

	"github.com/alfisar/jastip-import/helpers/consts"
	"github.com/alfisar/jastip-import/helpers/errorhandler"
	"github.com/alfisar/jastip-import/helpers/helper"
	"github.com/alfisar/jastip-import/helpers/jwthandler"

	"github.com/go-redis/redis/v8"
)

func validateUser(ctx context.Context, poolData *config.Config, repo repository.UserContractRepository, reposiRedis repoRedis.RedisRepositoryContract, data domain.UserLoginRequest) (idUser int, err domain.ErrorData) {
	defer handler.PanicError()
	result := domain.User{}

	block, errs := handler.AttempRedis(ctx, poolData, reposiRedis, consts.RedisToken, "Attemp_"+data.Username)
	if errs.Code != 0 {
		err = errs
		return
	} else if block {
		err = errorhandler.ErrBlocking()
		return
	}

	result, errs = getUser(poolData, repo, []string{"email"}, []any{data.Username})
	if errs.Code != 0 {
		err = errs
		return
	}
	if result.Id == 0 {
		result, errs = getUser(poolData, repo, []string{"nohp"}, []any{data.Username})
		if errs.Code != 0 {
			err = errs
			return
		}
		if result.Id == 0 {
			err = errorhandler.ErrLogin(nil)
			return
		}
	}
	errData := helper.Verify(result.Password, data.Password)
	if errData != nil {
		message := fmt.Sprintf("Invalid verify pass on func validate user : %s", errData.Error())
		log.Println(message)

		err = errorhandler.ErrInvalidLogic(errorhandler.ErrCodeInvalidInput, errorhandler.ErrMsgLoginFailed, errData.Error())
		return
	}

	idUser = result.Id

	return
}

func getUser(poolData *config.Config, repo repository.UserContractRepository, key []string, value []any) (result domain.User, err domain.ErrorData) {
	var (
		errData error
	)

	defer handler.PanicError()

	where := map[string]any{}
	for i, v := range key {
		where[v] = value[i]
	}

	result, errData = repo.Get(poolData.DBSql, where)
	if errData != nil && errData.Error() != "get users error : record not found" {
		return result, errorhandler.ErrGetData(errData)

	}

	return
}

func getToken(ctx context.Context, poolData *config.Config, conn *redis.Client, key string, id int, repoRedis repoRedis.RedisRepositoryContract) (token string, err domain.ErrorData) {

	jwts := jwthandler.GetJwt()
	tokenData, errData := jwts.GetToken(consts.TokenExp, id)
	if errData != nil {
		message := fmt.Sprintf("failed generate token : %s", errData.Error())
		log.Println(message)

		err = errorhandler.ErrInvalidLogic(errorhandler.ErrCodeGenerateToken, errorhandler.ErrMsgFailedGenerateToken, errData.Error())
		return
	}

	tokenData, errData = helper.EncryptAES256CBC(poolData.Hash.Key, []byte(tokenData))
	if errData != nil {
		message := fmt.Sprintf("failed generate token on func getToken: %s", errData.Error())
		log.Println(message)

		err = errorhandler.ErrInvalidLogic(errorhandler.ErrCodeGenerateToken, errorhandler.ErrMsgFailedGenerateToken, errData.Error())
		return
	}

	token = tokenData
	errData = repoRedis.Insert(ctx, conn, key, token, consts.TokenExp)
	if errData != nil {
		message := fmt.Sprintf("failed save token : %s", errData.Error())
		log.Println(message)
	}
	return
}
