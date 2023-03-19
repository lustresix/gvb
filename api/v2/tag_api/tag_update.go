package tag_api

import (
	"gbv2/config/log"
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagUpdateView(c *gin.Context) {
	var req TagReq
	id := c.Param("id")
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Errorw("err", "err", err, &req)
		res.FailWithMsg("参数上传错误", c)
		return
	}

	//防止重复
	var tag models.TagModel
	err = mysql.DB.Take(&tag, "id = ?", id).Error
	if err != nil {
		res.FailWithMsg("里面没有呀", c)
		return
	}

	err = mysql.DB.Model(&tag).Updates(map[string]any{
		"title": req.Title,
	}).Error

	if err != nil {
		log.Errorw("err", err)
		res.FailWithMsg("修改标签失败", c)
		return
	}
	res.OKWithMsg("修改标签成功", c)
}
