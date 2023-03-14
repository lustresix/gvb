package advert_api

import (
	"gbv2/config/log"
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
)

// AdvertUpdateView 修改广告
// @Tags 广告管理
// @Summary 修改广告
// @Description 修改广告
// @Param data body AdvertReq   true "查询参数"
// @Router /api/adverts/:id [put]
// @Produce json
// @Success 200 {Object} res.Response{data=string}
func (AdvertApi) AdvertUpdateView(c *gin.Context) {
	var req AdvertReq
	id := c.Param("id")
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Errorw("err", "err", err, &req)
		res.FailWithMsg("参数上传错误", c)
		return
	}

	//防止重复
	var advert models.AdvertModel
	err = mysql.DB.Take(&advert, "id = ?", id).Error
	if err != nil {
		res.FailWithMsg("里面没有呀", c)
		return
	}

	err = mysql.DB.Model(&advert).Updates(map[string]any{
		"title":   req.Title,
		"href":    req.Href,
		"images":  req.Images,
		"is_show": req.IsShow,
	}).Error

	if err != nil {
		log.Errorw("err", err)
		res.FailWithMsg("修改广告失败", c)
		return
	}
	res.OKWithMsg("修改广告成功", c)
}
