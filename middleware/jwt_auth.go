package middleware

import (
	"gbv2/models/res"
	"gbv2/utils/jwt"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMsg("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			res.FailWithMsg("token错误", c)
			c.Abort()
			return
		}
		// 登录的用户
		c.Set("claims", claims)
	}
}
