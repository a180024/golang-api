package server

import (
	"github.com/a180024/golang-api/config"
	"github.com/a180024/golang-api/controllers"
)

func Init(userController controllers.UserController) {
	r := NewRouter(userController)
	c := config.GetConfig()
	r.Run(c.GetString("port"))
}
