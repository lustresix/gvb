package advert_api

import (
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
	"strconv"
)

// AdvertRemoveView 删除广告
// @Tags 广告管理
// @Summary 删除广告
// @Description 删除广告
// @Param data body models.RemoveRequest  true "删除参数"
// @Router /api/adverts [delete]
// @Produce json
// @Success 200 {Object} res.Response{"msg":"响应"}
func (AdvertApi) AdvertRemoveView(c *gin.Context) {
	var re models.RemoveRequest
	err := c.ShouldBindJSON(&re)
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}

	var advertList []models.AdvertModel
	count := mysql.DB.Find(&advertList, re.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("广告不存在", c)
		return
	}
	mysql.DB.Delete(&advertList)
	total := strconv.Itoa(int(count))
	res.OKWithMsg("共删除"+total+"广告", c)
}
