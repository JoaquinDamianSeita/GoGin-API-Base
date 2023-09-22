package config

import (
	"GoGin-API-Base/api/handlers"
	"GoGin-API-Base/repository"
	"GoGin-API-Base/services"
)

type Initialization struct {
	userRepo  repository.UserRepository
	userSvc   services.UserService
	UserHdler handlers.UserHandler
}

func NewInitialization(userRepo repository.UserRepository,
	userService services.UserService,
	UserHdler handlers.UserHandler) *Initialization {
	return &Initialization{
		userRepo:  userRepo,
		userSvc:   userService,
		UserHdler: UserHdler,
	}
}
