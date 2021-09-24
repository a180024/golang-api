package server

import (
	"github.com/a180024/nft_api/controller/users"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("v1")
	{
		authGroup := v1.Group("auth")
		{
			authGroup.POST("/register", users.Register)
		}
	}

	return router
}
