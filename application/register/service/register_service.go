package service

import (
	"context"
	"fmt"
	repoRedis "jastip/application/redis/repository"
	"jastip/application/user/repository"
	"log"

	"github.com/alfisar/jastip-import/domain"

	"github.com/alfisar/jastip-import/helpers/handler"

	"github.com/alfisar/jastip-import/helpers/consts"
	"github.com/alfisar/jastip-import/helpers/errorhandler"
	"github.com/alfisar/jastip-import/helpers/helper"
)

type registerService struct {
	repo      repository.UserContractRepository
	repoRedis repoRedis.RedisRepositoryContract
}

func NewRegisterService(repo repository.UserContractRepository, repoRedis repoRedis.RedisRepositoryContract) *registerService {
	return &registerService{
		repo:      repo,
		repoRedis: repoRedis,
	}
}

func (s *registerService) Register(ctx context.Context, poolData *domain.Config, data domain.User) (result domain.User, err domain.ErrorData) {

	defer handler.PanicError()

	// Validasi apakah sudah ada user dengan email dan nomor hp yang sama
	err = validateUser(poolData, s.repo, data)
	if err.Code != 0 {
		return
	}

	data.Password, err = helper.GeneratePass(data.Password)
	if err.Code != 0 {
		return
	}

	// save ke dalam database
	id, errs := saveUserToDatabase(poolData, s.repo, data)
	if errs.Code != 0 {
		err = errs
		return
	}

	// Generate dan Save OTP
	otp, errs := generateAndSaveOTP(ctx, poolData, s.repoRedis, data.Email+data.NoHP)
	if errs.Code != 0 {
		err = errs
		return
	}

	// set data attemp redis
	block, errs := handler.AttempRedis(ctx, poolData, s.repoRedis, consts.RedisOTP, consts.Attemp+data.Email+data.NoHP)
	if errs.Code != 0 {
		err = errs
		return
	}

	if block {
		err = errorhandler.ErrBlocking()
		return
	}

	data.Id = id
	result = data
	// Send Email ke user
	go sendEmail(poolData, data.Email, data.FullName, otp)

	return
}

func (s *registerService) VerifyOTP(ctx context.Context, poolData *domain.Config, email string, nohp string, otp string) (err domain.ErrorData) {
	defer func() {
		if r := recover(); r != nil {
			errData := fmt.Errorf(fmt.Sprintf("%s", r))
			err = errorhandler.ErrInsertData(errData)
		}

		if err.Code == 0 {
			s.repoRedis.Delete(ctx, poolData.DBRedis[consts.RedisOTP], "Attemp_"+email+nohp)
		}
	}()

	// validasi otp apakah sama atau tidak
	err = validateOtp(ctx, poolData, s.repoRedis, email+nohp, otp)
	if err.Code != 0 {
		return
	}

	// update status user
	key := []string{
		"email", "nohp",
	}
	value := []any{
		email, nohp,
	}
	keyUpdate := []string{
		"status",
	}
	valueUpdate := []any{
		1,
	}

	err = updateDataUser(poolData, s.repo, key, value, keyUpdate, valueUpdate)
	if err.Code != 0 {
		return
	}
	return
}

func (s *registerService) ResendOtp(ctx context.Context, poolData *domain.Config, email string, nohp string, fullName string) (err domain.ErrorData) {
	defer handler.PanicError()

	block, errs := handler.AttempRedis(ctx, poolData, s.repoRedis, consts.RedisOTP, consts.Attemp+email+nohp)
	if errs.Code != 0 {
		err = errs
		return
	}

	if block {
		err = errorhandler.ErrBlocking()
		return
	}

	result, errs := checkUserExist(poolData, s.repo, "email", email)
	if errs.Code != 0 && result.Id == 0 {
		err = errs
		return
	}

	if result.Status != 0 {
		message := fmt.Sprintf("Invalid logic on func registration : %s", "User Already Active")
		log.Println(message)

		return errorhandler.ErrInvalidLogic(errorhandler.ErrCodeInvalidInput, errorhandler.ErrMsgDataExist, message)
	}

	otp, errs := generateAndSaveOTP(ctx, poolData, s.repoRedis, email+nohp)
	if errs.Code != 0 {
		err = errs
		return
	}

	go sendEmail(poolData, email, fullName, otp)
	return
}
