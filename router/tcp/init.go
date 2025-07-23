package router

import (
	repoAddress "jastip/application/address/repository"
	profileControll "jastip/application/profile/controller/tcp"
	serviceProfile "jastip/application/profile/service"
	simpleControll "jastip/application/simple/controller/tcp"
	repoUser "jastip/application/user/repository"
)

func SimpleInit() *simpleControll.SimpleGrpcController {
	return simpleControll.NewSimpleGrpcController()
}

func ProfileInit() *profileGrpcRouter {
	repo := repoUser.NewUserRpository()
	repoAddr := repoAddress.NewAddressRepository()

	serv := serviceProfile.NewProfileService(repo, repoAddr)
	controlProfile := profileControll.NewProfileGrpcController(serv)
	return NewProfileGrpcRouter(*controlProfile)
}
