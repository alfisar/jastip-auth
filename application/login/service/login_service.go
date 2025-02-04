package service

import (
	"context"
	repoRedis "jastip/application/redis/repository"
	"jastip/application/user/repository"
	"jastip/config"
	"jastip/domain"
	"jastip/internal/consts"
	"jastip/internal/handler"
	"strconv"
)

type loginService struct {
	repo      repository.UserContractRepository
	repoRedis repoRedis.RedisRepositoryContract
}

func NewLoginService(repo repository.UserContractRepository,
	repoRedis repoRedis.RedisRepositoryContract) *loginService {
	return &loginService{
		repo:      repo,
		repoRedis: repoRedis,
	}
}
func (s *loginService) Login(ctx context.Context, poolData *config.Config, data domain.UserLoginRequest) (token string, err domain.ErrorData) {

	defer func() {
		errs := handler.PanicError()
		if errs.Code != 0 {
			err = errs
		}

		if err.Code == 0 {
			s.repoRedis.Delete(ctx, poolData.DBRedis[consts.RedisToken], "Attemp_"+data.Username)
		}
	}()
	id, errs := validateUser(ctx, poolData, s.repo, s.repoRedis, data)
	if errs.Code != 0 {
		err = errs
		return
	}

	token, err = getToken(ctx, poolData.DBRedis[consts.RedisToken], "TOKEN_"+strconv.Itoa(id), id, s.repoRedis)

	return
}
