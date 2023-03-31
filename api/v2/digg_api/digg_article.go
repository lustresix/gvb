package digg_api

import (
	"gbv2/models"
	"gbv2/models/res"
	"gbv2/service/redis_ser"
	"github.com/gin-gonic/gin"
)

func (DiggApi) DiggArticle(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}

	redis_ser.Digg(cr.ID)
	res.OKWithMsg("点赞成功", c)
}
