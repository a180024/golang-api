package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/a180024/golang-api/config"
	"github.com/a180024/golang-api/controllers"
	"github.com/a180024/golang-api/db"
	"github.com/a180024/golang-api/models"
	"github.com/a180024/golang-api/server"
	"github.com/a180024/golang-api/services"
)

func main() {
	env := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*env)
	db := db.Init()

	// Create Repositories
	var userRepository models.UserRepository = models.NewUserRepository(db)
	// Create Services
	var userService services.UserService = services.NewUserService(userRepository)
	var jwtService services.JWTService = services.NewJWTService()
	var loginService services.LoginService = services.NewLoginService(userRepository)
	// Create Controllers
	var userController controllers.UserController = controllers.NewUserController(userService)
	var loginController controllers.LoginController = controllers.NewLoginController(loginService, jwtService)

	server.Init(userController, loginController)
}
