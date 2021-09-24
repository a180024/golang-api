package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/a180024/nft_api/config"
	"github.com/a180024/nft_api/db"
	"github.com/a180024/nft_api/models"
	"github.com/a180024/nft_api/server"
	"github.com/a180024/nft_api/service"
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
	var userService service.UserService = service.NewUserService(userRepository)

	server.Init()
}
