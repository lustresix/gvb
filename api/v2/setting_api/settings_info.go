package setting_api

import (
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.FailWithCode(res.ErrorUserNotExist, c)
}
