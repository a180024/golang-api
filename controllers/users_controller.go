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
	FindOneByID(c *gin.Context)
}

func NewUserController(userService services.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (userController *userController) Register(c *gin.Context) {
	var user dto.UserDto

	// Bind JSON
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// Call userService
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

func (userController *userController) FindOneByID(c *gin.Context) {
	var ID dto.UserIdDto

	// Bind URI
	if err := c.ShouldBindUri(&ID); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// Call userService
	user, err := userController.userService.FindOneByID(ID.ID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
