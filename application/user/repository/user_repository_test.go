package repository

import (
	"context"
	"testing"

	"github.com/alfisar/jastip-import/config"

	"github.com/alfisar/jastip-import/domain"

	"github.com/stretchr/testify/require"
)

var (
	poolData *domain.Config
	repo     = NewUserRpository()
	ctx      = context.Background()
)

func TestInsert(t *testing.T) {
	config.Init()
	poolData = domain.DataPool.Get().(*domain.Config)
	data := domain.User{
		FullName: "Alfisar Rachman",
		Password: "$2a$10$47.eeIVUSlxJ7jBj1tScn.tF2VyVGUj.luuamN8oAg.VWjO7RY7U2",
		Email:    "alfisarrachman@gmail.com",
		NoHP:     "081291276666",
		Role:     1,
		Status:   0,
		Username: "alfisar",
	}
	_, err := repo.Create(ctx, poolData.DBSql, data)
	require.Nil(t, err)

}

func TestGet(t *testing.T) {
	where := map[string]interface{}{
		"email": "alfisarrachman@gmail.com",
	}
	_, err := repo.Get(ctx, poolData.DBSql, where)
	require.Nil(t, err)

}
func TestGetFaileUser(t *testing.T) {
	where := map[string]interface{}{
		"email": "alfisarrachma1n@gmail.com",
	}
	_, err := repo.Get(ctx, poolData.DBSql, where)
	require.NotNil(t, err)

}
func TestInsertFailed(t *testing.T) {
	data := domain.User{
		FullName: "Alfisar Rachman",
		Password: "$2a$10$47.eeIVUSlxJ7jBj1tScn.tF2VyVGUj.luuamN8oAg.VWjO7RY7U2",
		Email:    "alfisarrachman@gmail.com",
		NoHP:     "081291276666",
		Role:     1,
		Status:   0,
		Username: "alfisar",
	}
	poolData.DBSql = nil
	_, err := repo.Create(ctx, poolData.DBSql, data)
	require.NotNil(t, err)

}
func TestGetFailed(t *testing.T) {
	where := map[string]interface{}{
		"email": "alfisarrachman1@gmail.com",
	}
	poolData.DBSql = nil
	_, err := repo.Get(ctx, poolData.DBSql, where)
	require.NotNil(t, err)

}
