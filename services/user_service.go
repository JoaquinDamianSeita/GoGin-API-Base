package services

import (
	api_responses "GoGin-API-Base/api_responses"
	dao "GoGin-API-Base/dao"
	"GoGin-API-Base/repository"
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
	if err != nil || request.Username == "" || request.Email == "" || request.Password == "" {
		log.Error("Invalid parameters: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, api_responses.ApiErrorResponse{
			Error: "Invalid parameters.",
		})
		return
	}

	u.userRepository.Save(&request)

	c.JSON(http.StatusOK, api_responses.ApiMessageResponse{Message: "User successfully created."})
}

func UserServiceInit(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}
