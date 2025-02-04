package service

import (
	"context"
	"errors"
	"fmt"
	repoRedis "jastip/application/redis/repository"
	"jastip/application/user/repository"
	"jastip/config"
	"jastip/domain"
	"jastip/internal/consts"
	"jastip/internal/errorhandler"
	"jastip/internal/general"
	"jastip/internal/handler"
	"jastip/internal/helper"
	"log"
	"sync"
)

func validateUser(poolData *config.Config, repo repository.UserContractRepository, data domain.User) (err domain.ErrorData) {
	var wg sync.WaitGroup
	defer handler.PanicError()

	errChan := make(chan domain.ErrorData, 2)
	wg.Add(2)
	go func() {
		defer wg.Done()
		_, err := checkUserExist(poolData, repo, "email", data.Email)
		errChan <- err

	}()

	go func() {
		defer wg.Done()
		_, err := checkUserExist(poolData, repo, "nohp", data.NoHP)
		errChan <- err
	}()

	wg.Wait()

	close(errChan)
	for v := range errChan {
		if v.Code != 0 {
			err = v
			return
		}
	}
	return
}

func checkUserExist(poolData *config.Config, repo repository.UserContractRepository, field string, value string) (result domain.User, err domain.ErrorData) {
	defer handler.PanicError()
	errData := errors.New("")
	where := map[string]any{
		field: value,
	}

	result, errData = repo.Get(poolData.DBSql, where)
	if errData != nil && errData.Error() != "get users error : record not found" {
		return result, errorhandler.ErrGetData(errData)

	}

	if result.Id != 0 {
		message := fmt.Sprintf("Invalid logic on func registration : %s", "email or nohp already exist")
		log.Println(message)

		return result, errorhandler.ErrInvalidLogic(errorhandler.ErrCodeInvalidInput, errorhandler.ErrMsgEmailNoHPUnique, message)
	}

	return
}

func saveUserToDatabase(poolData *config.Config, repo repository.UserContractRepository, data domain.User) (id int, err domain.ErrorData) {
	defer handler.PanicError()

	id, errData := repo.Create(poolData.DBSql, data)
	if errData != nil {
		message := fmt.Sprintf("Error save data to DB on func registration and func SaveUserToDatabase : %s", errData.Error())
		log.Println(message)

		err = errorhandler.ErrInsertData(errData)
		return
	}

	return
}

func generateAndSaveOTP(ctx context.Context, poolData *config.Config, repo repoRedis.RedisRepositoryContract, key string) (otp string, err domain.ErrorData) {
	defer handler.PanicError()
	otps, errData := general.GetRandomOTP(6)
	if errData != nil {
		message := fmt.Sprintf("Error generate otp on func registration and func GenerateAndSaveOTP : %s", errData.Error())
		log.Println(message)
		errorhandler.ErrInternal(errorhandler.ErrCodeGenerate, fmt.Errorf(message))
		return
	}

	errData = repo.Insert(ctx, poolData.DBRedis[consts.RedisOTP], key, otps, consts.RedisOTPExp)
	if errData != nil {
		message := fmt.Sprintf("Error insert data to redis on func registration and func GenerateAndSaveOTP : %s", errData.Error())
		log.Println(message)
		errorhandler.ErrInsertData(fmt.Errorf(message))
		return
	}
	otp = otps
	return
}

func sendEmail(poolData *config.Config, email string, fullName string, otp string) {

	errData := helper.SendEmailOTP(poolData, email, fullName, otp)
	if errData != nil {
		message := fmt.Sprintf("Error send email on func registration and func SendEmail : %s", errData.Error())
		log.Println(message)
		return
	}
}

func validateOtp(ctx context.Context, poolData *config.Config, repo repoRedis.RedisRepositoryContract, key string, otp string) (err domain.ErrorData) {
	defer handler.PanicError()

	otpRedis, errData := repo.Get(ctx, poolData.DBRedis[consts.RedisOTP], key)
	if errData != nil {
		message := fmt.Sprintf("Invalid otp on func verify otp and func ValidateOtp: %s", errData.Error())
		log.Println(message)

		err = errorhandler.ErrGetData(fmt.Errorf(message))
		return
	}

	if otp != otpRedis {
		message := fmt.Sprintf("Invalid otp on func verify otp and ValidateOtp : %s", "Invalid OTP")
		log.Println(message)

		err = errorhandler.ErrInvalidLogic(errorhandler.ErrCodeInvalidLogicBisnis, errorhandler.ErrMsgOTPInvalid, message)
		return
	}

	return
}

func updateDataUser(poolData *config.Config, repo repository.UserContractRepository, key []string, value []any, keyUpdate []string, valueUpdate []any) (err domain.ErrorData) {
	defer handler.PanicError()

	where := map[string]any{}
	for i, v := range key {
		where[v] = value[i]
	}

	updates := map[string]any{}
	for i, v := range keyUpdate {
		updates[v] = valueUpdate[i]
	}

	errData := repo.Update(poolData.DBSql, where, updates)
	if errData != nil {
		message := fmt.Sprintf("Failed update data on func verify otp and func UpdateDataUser : %s", errData.Error())
		log.Println(message)

		err = errorhandler.ErrInternal(errorhandler.ErrCodeUpdate, fmt.Errorf(message))
		return
	}
	return
}
