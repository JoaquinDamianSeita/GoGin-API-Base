package api

import (
	"GoGin-API-Base/api/middleware"
	"GoGin-API-Base/config"

	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.ErrorHandler())

	api := router.Group("/api")
	{
		user := api.Group("/users")
		user.POST("", init.UserHdler.RegisterUser)
	}

	return router
}