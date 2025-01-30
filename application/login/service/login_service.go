package service

import (
	"context"
	repoRedis "jastip/application/redis/repository"
	"jastip/application/user/repository"
	"jastip/config"
	"jastip/domain"
	"jastip/internal/handler"
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
	defer handler.PanicError()
	errs := validateUser(poolData, s.repo, data)
	if errs.Code != 0 {
		err = errs
		return
	}

	return
}
