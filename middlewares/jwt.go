package middlewares

import (
	"github.com/gin-gonic/gin"
	"online-judge/common/errCode"
	"online-judge/middlewares/jwt"
	"time"
	"online-judge/controllers"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		code = errCode.SUCCESS
		token := c.GetHeader("Authorization")

		if token == "" {
			code = errCode.TOKEN_MISSING
		} else {
			claims, err := jwt.ParseToken(token)
			if err != nil {
				code = errCode.INVALID_TOKEN
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = errCode.TOKEN_TIMEOUT
			}
		}

		if code != errCode.SUCCESS {
			controllers.Response(c,http.StatusUnauthorized,errCode.UNAUTHORIZED,nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
