package article_api

import (
	"gbv2/models"
	"gbv2/models/res"
	"gbv2/service/es_ser"
	"gbv2/service/redis_ser"
	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}
	redis_ser.Look(cr.ID)
	model, err := es_ser.CommDetail(cr.ID)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OKWithData(model, c)
}
