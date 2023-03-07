package images_api

import (
	"gbv2/models"
	"gbv2/models/res"
	"gbv2/service/common"
	"github.com/gin-gonic/gin"
)

func (ImagesApi) ImageListView(c *gin.Context) {

	var page models.PageInfo
	err := c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}
	list, count, err := common.CommonList(models.ImageModel{}, common.Option{PageInfo: page})
	if err != nil {
		res.FailWithMsg("数据库查询错误", c)
		return
	}
	res.OKWitList(list, count, c)

}
