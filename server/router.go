package server

import (
	"github.com/a180024/golang-api/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter(userController controllers.UserController, loginController controllers.LoginController) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("v1")
	{
		authGroup := v1.Group("auth")
		{
			authGroup.POST("/register", userController.Register)
			authGroup.POST("/login", loginController.Login)
		}
		userGroup := v1.Group("user")
		{
			userGroup.GET("/:id", userController.FindOneByID)
		}
	}

	return router
}
