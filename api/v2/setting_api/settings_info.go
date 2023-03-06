package setting_api

import (
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	viper.WatchConfig()
	res.OKWithData(viper.Get("site_info"), c)
}
