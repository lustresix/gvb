package tag_api

import (
	"gbv2/config/log"
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

type TagReq struct {
	Title string `json:"title" binding:"required" msg:"请输入标题" structs:"title"`
}

func (TagApi) TagCreateView(c *gin.Context) {
	var req TagReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Errorw("err", "err", err, &req)
		res.FailWithMsg("参数上传错误", c)
		return
	}

	//防止重复
	var tag models.TagModel
	err = mysql.DB.Take(&tag, "title = ?", req.Title).Error
	if err == nil {
		res.FailWithMsg("已经在里面了", c)
		return
	}

	maps := structs.Map(&req)
	err = mysql.DB.Create(maps).Error

	if err != nil {
		log.Errorw("err", "err", err)
		res.FailWithMsg("添加标签失败", c)
		return
	}
	res.OKWithMsg("添加标签成功", c)
}
