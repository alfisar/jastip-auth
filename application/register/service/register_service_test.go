package service

import (
	"context"
	"jastip/config"
	"jastip/domain"
	"testing"

	repoRedis "jastip/application/redis/repository"
	"jastip/application/user/repository"

	"github.com/stretchr/testify/require"
)

var (
	poolData   *config.Config
	repoRediss = repoRedis.NewRedisRepository()
	repo       = repository.NewUserRpository()
	serv       = NewRegisterService(repo, repoRediss)
	ctx        = context.Background()
	otp        string
)

func TestRegister(t *testing.T) {
	config.Init()
	poolData = config.DataPool.Get().(*config.Config)
	data := domain.User{
		FullName: "Alfisar Rachman",
		Password: "$2a$10$47.eeIVUSlxJ7jBj1tScn.tF2VyVGUj.luuamN8oAg.VWjO7RY7U2",
		Email:    "alfisarrachman@gmail.com",
		NoHP:     "081291276666",
		Role:     1,
		Status:   0,
		Username: "alfisar",
	}
	_, err := serv.Register(ctx, poolData, data)
	require.Equal(t, domain.ErrorData{}, err)

}

func TestResendOTP(t *testing.T) {

	poolData = config.DataPool.Get().(*config.Config)

	err := serv.ResendOtp(ctx, poolData, "alfisarrachman@gmail.com", "081291276666", "alfisar rachman")

	require.Equal(t, domain.ErrorData{}, err)
}

func TestVerifyOTP(t *testing.T) {

	poolData = config.DataPool.Get().(*config.Config)

	err := serv.VerifyOTP(ctx, poolData, "alfisarrachman@gmail.com", "081291276666", otp)
	require.Equal(t, domain.ErrorData{}, err)

}
