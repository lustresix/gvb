package advert_api

import (
	"gbv2/config/log"
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
)

type AdvertReq struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题"`
	Href   string `json:"href" binding:"required,url" msg:"请输入连接"`
	Images string `json:"images" binding:"required,url" msg:"请输入图片"`
	IsShow bool   `json:"is_show" binding:"required" msg:"请输入是否展示"`
}

func (AdvertApi) AdvertCreateView(c *gin.Context) {
	var req AdvertReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Errorw("err", "err", err, &req)
		res.FailWithMsg("参数上传错误", c)
		return
	}

	//防止重复
	var advert models.AdvertModel
	err = mysql.DB.Take(&advert, "title = ?", req.Title).Error
	if err == nil {
		res.FailWithMsg("已经在里面了", c)
		return
	}

	err = mysql.DB.Create(&models.AdvertModel{
		Title:  req.Title,
		Href:   req.Href,
		Images: req.Images,
		IsShow: req.IsShow,
	}).Error

	if err != nil {
		log.Errorw("err", err)
		res.FailWithMsg("添加广告失败", c)
		return
	}
	res.OKWithMsg("添加广告成功", c)
}
