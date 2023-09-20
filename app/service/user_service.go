package service

import (
	"GoGin-API-Base/app/domain/dao"
	"GoGin-API-Base/app/domain/dto"
	"GoGin-API-Base/app/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type UserService interface {
	RegisterUser(c *gin.Context)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func (u UserServiceImpl) RegisterUser(c *gin.Context) {
	var request dao.User
	err := c.ShouldBindJSON(&request)
	log.Println(err, request.Username == "", request.Email == "", request.Password == "")
	if err != nil || request.Username == "" || request.Email == "" || request.Password == "" {
		log.Error("Invalid parameters: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.ApiErrorResponse{
			Error: "Invalid parameters.",
		})
		return
	}

	u.userRepository.Save(&request)

	c.JSON(http.StatusOK, dto.ApiMessageResponse{Message: "User successfully created."})
}

func UserServiceInit(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}
