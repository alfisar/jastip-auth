package handler

import (
	"context"
	"fmt"
	"jastip/config"
	"log"
	"strconv"
	"time"

	"github.com/alfisar/jastip-import/domain"

	"github.com/alfisar/jastip-import/helpers/consts"
	"github.com/alfisar/jastip-import/helpers/errorhandler"

	repoRedis "jastip/application/redis/repository"
)

func PanicError() (err domain.ErrorData) {

	if r := recover(); r != nil {
		err = errorhandler.ErrInternal(errorhandler.ErrCodePanic, fmt.Errorf(fmt.Sprintf("%s", r)))
	}

	return
}

func setAttempRedis(ctx context.Context, poolData *config.Config, repo repoRedis.RedisRepositoryContract, dbRedis string, key string) (err domain.ErrorData) {
	defer PanicError()

	errData := repo.Incr(ctx, poolData.DBRedis[dbRedis], key)
	if errData != nil {
		message := fmt.Sprintf("Failed incr data on func register and func SetAttempRedis : %s", errData.Error())
		log.Println(message)

		err = errorhandler.ErrInternal(errorhandler.ErrCodeUpdate, fmt.Errorf(message))
		return
	}

	return
}

func setExpAttempRedis(ctx context.Context, poolData *config.Config, repo repoRedis.RedisRepositoryContract, dbRedis string, key string, exp time.Duration) (err domain.ErrorData) {
	defer PanicError()

	errData := repo.Exp(ctx, poolData.DBRedis[dbRedis], key, exp)
	if errData != nil {
		message := fmt.Sprintf("Failed incr data on func register and func SetExpAttempRedis : %s", errData.Error())
		log.Println(message)

		err = errorhandler.ErrInternal(errorhandler.ErrCodeUpdate, fmt.Errorf(message))
		return
	}

	return
}

func getAttempRedis(ctx context.Context, poolData *config.Config, repo repoRedis.RedisRepositoryContract, dbRedis string, key string) (attemp int, err domain.ErrorData) {
	defer PanicError()

	data, errData := repo.Get(ctx, poolData.DBRedis[dbRedis], key)
	if errData != nil {
		if errData.Error() != "get redis error : redis: nil" {
			message := fmt.Sprintf("Failed get data on func register and func GetAttempRedis : %s", errData.Error())
			log.Println(message)

			err = errorhandler.ErrInternal(errorhandler.ErrCodeUpdate, fmt.Errorf(message))
			return
		}
	}

	if data != "" {
		attemp, errData = strconv.Atoi(data)
		if errData != nil {
			message := fmt.Sprintf("Failed parsing data on func register and func GetAttempRedis : %s", errData.Error())
			log.Println(message)

			err = errorhandler.ErrInternal(errorhandler.ErrCodeParsing, fmt.Errorf(message))
			return
		}
	}
	return
}

func AttempRedis(ctx context.Context, poolData *config.Config, repo repoRedis.RedisRepositoryContract, dbRedis string, key string) (block bool, err domain.ErrorData) {
	dataAttemp, errs := getAttempRedis(ctx, poolData, repo, dbRedis, key)
	if errs.Code != 0 {
		err = errs
		return
	}

	err = setAttempRedis(ctx, poolData, repo, dbRedis, key)
	if err.Code != 0 {
		return
	}
	if dataAttemp >= consts.AttempOTP {
		block = true
		err = setExpAttempRedis(ctx, poolData, repo, dbRedis, key, consts.RedisOTPExp)
		if err.Code != 0 {
			return
		}
	}
	return
}
