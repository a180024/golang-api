package controllers

import (
	"net/http"

	"github.com/a180024/golang-api/dto"
	"github.com/a180024/golang-api/services"
	"github.com/gin-gonic/gin"
)

type loginController struct {
	loginService services.LoginService
	jwtService   services.JWTService
}

type LoginController interface {
	Login(c *gin.Context)
}

func NewLoginController(loginService services.LoginService, jwtService services.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

// Login godoc
// @Summary User login
// @Schemes
// @Description Checks if user is authenticated and returns JWT token
// @Accept json
// @Produce json
// @Param Credentials body dto.LoginDto true "Your login credentials"
// @Success 200 {object} map[string]interface{}
// @Router /auth/login [post]
func (loginController *loginController) Login(c *gin.Context) {
	var credentials dto.LoginDto

	// Bind JSON
	err := c.ShouldBind(&credentials)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Call loginService to get password hash
	isAuthenticated, err := loginController.loginService.Login(credentials.UserName, credentials.Password)
	if !isAuthenticated {
		c.JSON(http.StatusBadRequest, &dto.ResponseDto{
			Message: "Login failed",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// GenerateToken
	ss := loginController.jwtService.GenerateToken(credentials.UserName)
	if len(ss) > 0 {
		c.JSON(http.StatusOK, gin.H{"Token": ss})
		return
	}

	c.JSON(http.StatusOK, &dto.ResponseDto{
		Message: "Unable to generate JWT token",
	})
}
