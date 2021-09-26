package tests

import (
	"github.com/a180024/golang-api/config"
	"github.com/a180024/golang-api/controllers"
	"github.com/a180024/golang-api/db"
	"github.com/a180024/golang-api/middlewares"
	"github.com/a180024/golang-api/models"
	"github.com/a180024/golang-api/services"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	config.Init("test")
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

	// Setup Router
	router := gin.New()
	gin.SetMode(gin.TestMode)
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
