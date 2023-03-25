package article_api

import (
	"gbv2/models/res"
	"gbv2/service/es_ser"
	"github.com/gin-gonic/gin"
)

type ArticleDetailRequest struct {
	Title string `json:"title" form:"title"`
}

func (ArticleApi) ArticleDetailByTitleView(c *gin.Context) {
	var cr ArticleDetailRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}
	model, err := es_ser.CommDetailByKeyword(cr.Title)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OKWithData(model, c)
}
