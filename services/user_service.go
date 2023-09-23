package services

import (
	"GoGin-API-Base/api/auth"
	dao "GoGin-API-Base/dao"
	"GoGin-API-Base/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type UserService interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u UserServiceImpl) RegisterUser(c *gin.Context) {
	var request dao.User

	validationError := c.ShouldBindJSON(&request)
	if validationError != nil || request.Username == "" || request.Email == "" || request.Password == "" {
		log.Error("Invalid parameters: ", validationError)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters."})
		return
	}

	_, recordError := u.userRepository.Save(&request)
	if recordError != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": recordError.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User successfully created."})
}

func (u UserServiceImpl) LoginUser(c *gin.Context) {
	var request LoginRequest
	var user dao.User
	validationError := c.ShouldBindJSON(&request)
	if validationError != nil || request.Email == "" || request.Password == "" {
		log.Error("Invalid parameters: ", validationError)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters."})
		return
	}

	user, recordError := u.userRepository.FindUserByEmail(request.Email)
	if recordError != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
		return
	}

	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	expiresIn, tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString, "expires_in": expiresIn})
}

func UserServiceInit(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}
