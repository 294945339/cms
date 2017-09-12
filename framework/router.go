package framework

import (
	"github.com/gin-gonic/gin"
)

func LoadHandler(router *gin.Engine) {

	authMiddleware := getAuthMiddleware()

	router.POST("/login", authMiddleware.LoginHandler)
	router.POST("/refresh", authMiddleware.RefreshHandler)


}
