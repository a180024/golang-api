package users

import (
	"github.com/a180024/nft_api/dto/users"
	services "github.com/a180024/nft_api/services/users"
	"github.com/a180024/nft_api/utils/errors"
	"github.com/gin-gonic/gin"
)

/* Register User */
func Register(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(err.Status, err)
	}

	services.CreateUser(user)
}
