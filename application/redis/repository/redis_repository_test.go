package repository

import (
	"context"
	"jastip/config"
	"jastip/internal/consts"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	poolData *config.Config
	repo     = NewRedisRepository()
	ctx      = context.Background()
)

func TestInsert(t *testing.T) {
	config.Init()
	poolData = config.DataPool.Get().(*config.Config)

	err := repo.Insert(ctx, poolData.DBRedis[consts.RedisToken], "coba", "nilai", 0)
	require.Nil(t, err)

}

func TestGet(t *testing.T) {
	_, err := repo.Get(ctx, poolData.DBRedis[consts.RedisToken], "coba")
	require.Nil(t, err)

}

func TestDelete(t *testing.T) {

	err := repo.Delete(ctx, poolData.DBRedis[consts.RedisToken], "coba")
	require.Nil(t, err)

}

func TestGetFailed(t *testing.T) {
	_, err := repo.Get(ctx, poolData.DBRedis[consts.RedisToken], "coba")
	require.NotNil(t, err)

}

func TestInsertConnNil(t *testing.T) {
	poolData.DBRedis[consts.RedisToken] = nil
	err := repo.Insert(ctx, poolData.DBRedis[consts.RedisToken], "coba", "nilai", 0)
	require.NotNil(t, err)

}

func TestGetConnNil(t *testing.T) {
	poolData.DBRedis[consts.RedisToken] = nil
	_, err := repo.Get(ctx, poolData.DBRedis[consts.RedisToken], "coba")
	require.NotNil(t, err)

}

func TestDeleteConnNil(t *testing.T) {
	poolData.DBRedis[consts.RedisToken] = nil
	err := repo.Delete(ctx, poolData.DBRedis[consts.RedisToken], "coba")
	require.NotNil(t, err)

}
