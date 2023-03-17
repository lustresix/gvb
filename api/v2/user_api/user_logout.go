package user_api

import (
	"fmt"
	"gbv2/config/log"
	"gbv2/config/redis"
	"gbv2/models/res"
	"gbv2/utils/jwt"
	"github.com/gin-gonic/gin"
	"time"
)

func (UserApi) LogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	token := c.Request.Header.Get("token")

	at := claims.ExpiresAt
	now := time.Now()

	diff := at.Time.Sub(now)

	err := redis.RDB.Set(fmt.Sprintf("logout_%s", token), "", diff).Err()

	if err != nil {
		log.Errorw("注销失败", "err", err)
		res.FailWithMsg("注销失败", c)
	}

	res.OKWithMsg("注销成功", c)
}
