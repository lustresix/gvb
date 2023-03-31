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
	var cr es_ser.Option
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}

	list, count, err := es_ser.CommList(es_ser.Option{
		PageInfo: cr.PageInfo,
		Fields:   []string{"string", "content"},
		Tag:      cr.Tag,
	})
	if err != nil {
		log.Errorw("err", "err", err)
		res.OKWithMsg("查询失败", c)
		return
	}
	//  如果list为空
	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.ArticleModel, 0)
		res.OKWitList(list, int64(count), c)
		return
	}

	res.OKWitList(filter.Omit("list", list), int64(count), c)
}
