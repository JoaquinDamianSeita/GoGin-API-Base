package handlers

import (
	"GoGin-API-Base/services"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	RegisterUser(c *gin.Context)
}

type UserHandlerImpl struct {
	svc services.UserService
}

func (u UserHandlerImpl) RegisterUser(c *gin.Context) {
	u.svc.RegisterUser(c)
}

func UserHandlerInit(userService services.UserService) *UserHandlerImpl {
	return &UserHandlerImpl{
		svc: userService,
	}
}
