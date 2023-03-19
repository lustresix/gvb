package tag_api

import (
	"gbv2/models"
	"gbv2/models/res"
	"gbv2/service/common"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagListView(c *gin.Context) {
	var info models.PageInfo
	err := c.ShouldBindQuery(&info)
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}
	list, count, _ := common.CommonList(models.TagModel{}, common.Option{
		PageInfo: info,
	})

	res.OKWitList(list, count, c)
}
