package middleware

import (
	"gbv2/config/redis"
	"gbv2/models/res"
	"gbv2/utils"
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

		// 如果在redis中说明注销过
		keys := redis.RDB.Keys("logout_*").Val()
		list := utils.InList("logout_"+token, keys)
		if list {
			res.FailWithMsg("token失效", c)
			c.Abort()
			return
		}

		// 登录的用户
		c.Set("claims", claims)
	}
}
