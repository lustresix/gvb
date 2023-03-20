package message_api

import (
	"gbv2/models"
	"gbv2/models/res"
	"gbv2/service/common"
	"github.com/gin-gonic/gin"
)

func (MessageApi) MessageAllListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
	}
	list, count, _ := common.CommonList(models.MessageModel{}, common.Option{
		PageInfo: cr,
	})
	res.OKWitList(list, count, c)
}
