package controllers

import (
	"log"
	"net/http"

	"github.com/a180024/golang-api/dto"
	"github.com/a180024/golang-api/services"
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
	var user dto.UserDto

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err := userController.userService.CreateUser(user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, &dto.ResponseDto{
		Message: "User successfully registered!",
	})
}
