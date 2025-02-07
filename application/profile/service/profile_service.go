package service

import (
	"context"
	repositoryAddress "jastip/application/address/repository"
	"jastip/application/user/repository"
	"jastip/config"
	"jastip/domain"
	"jastip/internal/errorhandler"
	"jastip/internal/helper"

	"gorm.io/gorm"
)

type profileService struct {
	repo        repository.UserContractRepository
	repoAddress repositoryAddress.AddressRepositoryContract
}

func NewProfileService(repo repository.UserContractRepository, repoAddress repositoryAddress.AddressRepositoryContract) *profileService {
	return &profileService{
		repo:        repo,
		repoAddress: repoAddress,
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

func (s *profileService) GetAddress(ctx context.Context, poolData *config.Config, userId int) (result domain.AddressResponse, err domain.ErrorData) {
	keys := []string{"user_id"}
	values := []any{userId}
	result, err = getDataAddress(ctx, poolData, s.repoAddress, keys, values)
	if err.Code != 0 {
		return
	}

	return
}

func (s *profileService) SaveAddress(ctx context.Context, poolData *config.Config, userID int, data map[string]any) (err domain.ErrorData) {
	err = inserSaveAddress(ctx, poolData, userID, s.repoAddress, data)
	return
}
