package server

import (
	"github.com/a180024/golang-api/config"
	"github.com/a180024/golang-api/controllers"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func Init(userController controllers.UserController, loginController controllers.LoginController) {
	r := NewRouter(userController, loginController)
	c := config.GetConfig()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(c.GetString("port"))
}
