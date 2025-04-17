package server

import (
	"surgicalsteel-shop/controller"
	"surgicalsteel-shop/infrastructure"
	"surgicalsteel-shop/repository"
	"surgicalsteel-shop/service"
)

func Init() {
	infrastructure.Init()
	repository.Init()
	service.Init()
	controller.Init()

}
