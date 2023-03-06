package setting_api

import (
	"gbv2/config/log"
	"gbv2/config/site"
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var site site.SiteInfo
	err := c.ShouldBindJSON(&site)
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}
	viper.WatchConfig()

	viper.Set("site_info", site)
	err = viper.WriteConfig()
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}
	res.OKWithData(viper.Get("site_info"), c)
	log.Infow("修改配置成功！")
}
