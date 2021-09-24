package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/a180024/api_template/config"
	"github.com/a180024/api_template/controllers"
	"github.com/a180024/api_template/db"
	"github.com/a180024/api_template/models"
	"github.com/a180024/api_template/server"
	"github.com/a180024/api_template/services"
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
	// Create Controllers
	var userController controllers.UserController = controllers.NewUserController(userService)

	server.Init(userController)
}
