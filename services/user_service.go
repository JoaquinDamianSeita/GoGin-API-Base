package services

import (
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

	// TODO: Add email uniq validation and error response
	_, recordError := u.userRepository.Save(&request)
	if recordError != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": recordError})
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": recordError})
		return
	}

	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// TODO: Add token generation for user
	// tokenString, err:= auth.GenerateJWT(user.Email, user.Username)
	// if err != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	context.Abort()
	// 	return
	// }
	// context.JSON(http.StatusOK, gin.H{"token": tokenString})

	c.JSON(http.StatusOK, gin.H{"token": "dasdasdasdas", "expires_in": 123145})
}

func UserServiceInit(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}
