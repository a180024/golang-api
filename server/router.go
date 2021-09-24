package server

import (
	"github.com/a180024/golang-api/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter(userController controllers.UserController) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("v1")
	{
		authGroup := v1.Group("auth")
		{
			authGroup.POST("/register", userController.Register)
		}
	}

	return router
}
