package router

import (
	controllerLogin "jastip/application/loginlogout/controller"
	serviceLogin "jastip/application/loginlogout/service"

	controllerProfile "jastip/application/profile/controller"
	serviceProfile "jastip/application/profile/service"

	repoAddress "jastip/application/address/repository"
	repoRedis "jastip/application/redis/repository"
	controllerRegister "jastip/application/register/controller"
	"jastip/application/register/service"
	simpleControll "jastip/application/simple/controller"
	repoUser "jastip/application/user/repository"
	"jastip/internal/jwthandler"
	"jastip/internal/middlewere"
	"os"
)

func SimpleInit() *simpleControll.SimpleController {
	return simpleControll.NewSimpleController()
}

func RegisterInit() *regisRouter {
	repo := repoUser.NewUserRpository()
	repoRedis := repoRedis.NewRedisRepository()
	serv := service.NewRegisterService(repo, repoRedis)
	controlRegis := controllerRegister.NewRegisterController(serv)
	return NewRegisRouter(controlRegis)
}

func LoginLogoutInit() *loginLogoutRouter {
	repo := repoUser.NewUserRpository()
	repoRedis := repoRedis.NewRedisRepository()
	serv := serviceLogin.NewLoginService(repo, repoRedis)
	controlRegis := controllerLogin.NewLoginController(serv)
	return NewLoginLogoutRouter(controlRegis)
}

func setMiddleware() *middlewere.AuthenticateMiddleware {
	jwtData := jwthandler.GetJwt()
	jwtData.Secret = os.Getenv("JWT_SECRET")
	middleWR := middlewere.NewAuthenticateMiddleware(jwtData)
	return middleWR
}

func ProfileInit() *profileRouter {
	repo := repoUser.NewUserRpository()
	repoAddr := repoAddress.NewAddressRepository()

	serv := serviceProfile.NewProfileService(repo, repoAddr)
	controlProfile := controllerProfile.NewProfileController(serv)
	return NewProfileRouter(controlProfile)
}
