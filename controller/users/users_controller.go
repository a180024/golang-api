package controller

import (
	"net/http"

	"github.com/a180024/nft_api/dto/users"
	"github.com/a180024/nft_api/services"
	"github.com/a180024/nft_api/utils/errors"
	"github.com/gin-gonic/gin"
)

type userController struct {
	userService services.UserService
}

type UserController interface {
	Register(c *gin.Context)
}

func NewUserController(userService services.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (userController *userController) Register(c *gin.Context) {
	var user users.UserDto

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(err.Status, err)
	}

	userController.userService.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{"status": "Successfully registered!"})
}
