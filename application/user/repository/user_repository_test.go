package repository

import (
	"jastip/config"
	"jastip/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	poolData *config.Config
	repo     = NewUserRpository()
)

func TestInsert(t *testing.T) {
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
	_, err := repo.Create(poolData.DBSql, data)
	require.Nil(t, err)

}

func TestGet(t *testing.T) {
	where := map[string]interface{}{
		"email": "alfisarrachman@gmail.com",
	}
	_, err := repo.Get(poolData.DBSql, where)
	require.Nil(t, err)

}
func TestGetFaileUser(t *testing.T) {
	where := map[string]interface{}{
		"email": "alfisarrachma1n@gmail.com",
	}
	_, err := repo.Get(poolData.DBSql, where)
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
	_, err := repo.Create(poolData.DBSql, data)
	require.NotNil(t, err)

}
func TestGetFailed(t *testing.T) {
	where := map[string]interface{}{
		"email": "alfisarrachman1@gmail.com",
	}
	poolData.DBSql = nil
	_, err := repo.Get(poolData.DBSql, where)
	require.NotNil(t, err)

}
