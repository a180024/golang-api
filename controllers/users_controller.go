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

// Registration godoc
// @Summary Registers a user
// @Schemes
// @Description Registers a user by saving to DB
// @Accept json
// @Produce json
// @Param Details body dto.UserDto true "Your registration details"
// @Success 200 {object} dto.ResponseDto
// @Router /auth/register [post]
func (userController *userController) Register(c *gin.Context) {
	var user dto.CreateUserDto

	// Bind JSON
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Call userService
	err := userController.userService.CreateUser(user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, &dto.ResponseDto{
		Message: "User successfully registered!",
	})
}

// Find User godoc
// @Summary Finds a user by ID
// @Schemes
// @Security BearerAuth
// @Description Finds a user by querying from DB
// @Param id path string true "User ID"
// @Param authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} dto.UserResponseDto
// @Router /user/{id} [get]
func (userController *userController) FindOneByID(c *gin.Context) {
	var ID dto.UserIdDto

	// Bind URI
	if err := c.ShouldBindUri(&ID); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Call userService
	user, err := userController.userService.FindOneByID(ID.ID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
