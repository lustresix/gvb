package article_api

import (
	"gbv2/config/log"
	"gbv2/models"
	"gbv2/models/res"
	"gbv2/service/es_ser"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}
	list, count, err := es_ser.CommList(cr.Key, cr.Page, cr.Limit)
	if err != nil {
		log.Errorw("err", "err", err)
		res.OKWithMsg("查询失败", c)
		return
	}

	res.OKWitList(filter.Omit("list", list), int64(count), c)
}
