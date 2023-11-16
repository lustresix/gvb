package log_api

import (
	"gbv2/models"
	"gbv2/models/res"
	"gbv2/plugin/logstash"
	"gbv2/service/common"
	"github.com/gin-gonic/gin"
)

type LogRequest struct {
	models.PageInfo
	Level logstash.Level `form:"level"`
}

func (LogApi) LogListView(c *gin.Context) {
	var cr LogRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		return
	}
	list, count, _ := common.CommonList(logstash.LogStashModel{Level: cr.Level}, common.Option{
		PageInfo: cr.PageInfo,
	})
	res.OKWitList(list, count, c)
	return
}
