package server

import (
	"github.com/a180024/golang-api/config"
	"github.com/a180024/golang-api/controllers"
)

func Init(userController controllers.UserController, loginController controllers.LoginController) {
	r := NewRouter(userController, loginController)
	c := config.GetConfig()
	r.Run(c.GetString("port"))
}
