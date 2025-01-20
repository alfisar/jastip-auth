package router

import (
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
