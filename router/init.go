package router

import (
	controllerLogin "jastip/application/login/controller"
	serviceLogin "jastip/application/login/service"
	repoRedis "jastip/application/redis/repository"
	controllerRegister "jastip/application/register/controller"
	"jastip/application/register/service"
	simpleControll "jastip/application/simple/controller"
	repoUser "jastip/application/user/repository"
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
