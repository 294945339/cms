package framework

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt"
	"strings"

	userService "cms/service/user"
)

func getAuthMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "UserRealm",
		Key:        []byte("hjTSdjskiTOWJWRsjfskfjlkowqj8j23LQj"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(username string, password string, c *gin.Context) (string, bool) {
			user := userService.GetUserByName(username)
			if strings.EqualFold(user.IsActive, "Yes") && user.DepartmentId == 1 && user.RoleId == 100 && strings.Compare(password, user.Password) > -1 {
				return username, true
			}
			return username, false
		},
		Authorizator: func(username string, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, message)
		},
		TokenLookup: "header:Authorization",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}

	return authMiddleware
}
