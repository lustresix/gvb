package advert_api

import (
	"gbv2/models"
	"gbv2/models/res"
	"gbv2/service/common"
	"github.com/gin-gonic/gin"
	"strings"
)

// AdvertListView 获取广告
// @Tags 广告管理
// @Summary 获取广告
// @Description 获取广告
// @Param data query models.PageInfo   false "查询参数"
// @Router /api/adverts [get]
// @Produce json
// @Success 200 {Object} res.Response{data=res.ListResponse[models.AdvertModel]}
func (AdvertApi) AdvertListView(c *gin.Context) {
	var info models.PageInfo
	err := c.ShouldBindQuery(&info)
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}
	// 根据 referer 是从哪里发出的获取请求情况
	referer := c.GetHeader("Referer")
	isShow := true
	if strings.Contains(referer, "admin") {
		isShow = false
	}
	list, count, _ := common.CommonList(models.AdvertModel{IsShow: isShow}, common.Option{
		PageInfo: info,
	})
	res.OKWitList(list, count, c)
}
