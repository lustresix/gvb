package article_api

import (
	"gbv2/models/res"
	"gbv2/service/es_ser"
	"github.com/gin-gonic/gin"
)

type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var cr ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}
	model, err := es_ser.CommDetail(cr.ID)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OKWithData(model, c)
}
