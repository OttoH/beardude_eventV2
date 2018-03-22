package main

import (
	"strings"
	"time"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"

	"qpet-engine/dao"
)

func GetAuthMiddleware(Dao *dao.DAO) (j *jwt.GinJWTMiddleware) {
	return &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("beardude secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(user string, pwd string, c *gin.Context) (string, bool) {
			if user == "beardude" && pwd == "lalala" {
				return user, true
			}

			if res := Dao.ValidateRacer(user, pwd); res == true {
				return user, true
			}

			return user, false
		},
		Authorizator: func(user string, c *gin.Context) bool {
			url := c.Request.URL.String()

			// console need admin role
			if len(url) > 0 && strings.Index(url, "console") != -1 && user != "beardude" {
				return false
			}

			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}
}
