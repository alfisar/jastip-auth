package service

import (
	"context"
	"jastip/application/user/repository"
	"jastip/config"
	"jastip/domain"
	"jastip/internal/errorhandler"
	"jastip/internal/helper"

	"gorm.io/gorm"
)

type profileService struct {
	repo repository.UserContractRepository
}

func NewProfileService(repo repository.UserContractRepository) *profileService {
	return &profileService{
		repo: repo,
	}
}

func (s *profileService) Get(ctx context.Context, poolData *config.Config, userID int) (result domain.ProfileResponse, err domain.ErrorData) {
	where := map[string]any{
		"id": userID,
	}

	dataUser, errData := s.repo.Get(poolData.DBSql, where)
	if errData != nil && errData.Error() != "get users error : "+gorm.ErrRecordNotFound.Error() {
		err = errorhandler.ErrRecordNotFound()
		return
	}
	result = domain.ProfileResponse{
		Id:       dataUser.Id,
		FullName: dataUser.FullName,
		Username: dataUser.Username,
		Email:    dataUser.Email,
		NoHP:     dataUser.NoHP,
		Role:     dataUser.Role,
		Status:   dataUser.Status,
	}

	return
}

func (s *profileService) Update(ctx context.Context, poolData *config.Config, userId int, data map[string]any) (err domain.ErrorData) {
	if _, exists := data["password"]; exists {
		data["password"], err = helper.GeneratePass(data["password"].(string))
		if err.Code != 0 {
			return
		}
	}
	where := map[string]any{
		"id": userId,
	}

	errData := s.repo.Update(poolData.DBSql, where, data)
	if errData != nil {
		err = errorhandler.ErrUpdateData(errData)
		return
	}

	return
}
