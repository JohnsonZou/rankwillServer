package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rankwillServer/common"
	"rankwillServer/model"
	"rankwillServer/response"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		log.Println(tokenString)
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			response.Response(c, http.StatusUnauthorized, 401, nil, "Unauthorized token")
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.Response(c, http.StatusUnauthorized, 401, nil, "Unauthorized token")
			c.Abort()
			return
		}
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)
		if user.ID == 0 {
			response.Response(c, http.StatusUnauthorized, 401, nil, "Unauthorized token")
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
