package server

import (
	"github.com/a180024/golang-api/controllers"
	"github.com/a180024/golang-api/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter(userController controllers.UserController, loginController controllers.LoginController) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger())

	v1 := router.Group("v1")
	{
		authGroup := v1.Group("auth")
		{
			authGroup.POST("/register", userController.Register)
			authGroup.POST("/login", loginController.Login)
		}
		userGroup := v1.Group("user", middlewares.AuthorizeJWT())
		{
			userGroup.GET("/:id", userController.FindOneByID)
		}
	}

	return router
}
