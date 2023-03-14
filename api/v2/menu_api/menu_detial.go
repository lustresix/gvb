package menu_api

import (
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuDetailView(c *gin.Context) {
	// 先查菜单
	id := c.Param("id")
	var menuModel models.MenuModel
	err := mysql.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMsg("菜单不存在", c)
		return
	}
	// 查连接表
	var menuBanners []models.MenuImageModel
	mysql.DB.Preload("ImageModel").Order("sort desc").Find(&menuBanners, "menu_id = ?", id)
	var images = make([]Image, 0)
	for _, image := range menuBanners {
		if menuModel.ID != image.MenuID {
			continue
		}
		images = append(images, Image{
			ID:   image.ImageID,
			Path: image.ImageModel.Path,
		})
	}
	menuResponse := MenuResponse{
		MenuModel: menuModel,
		Images:    images,
	}
	res.OKWithData(menuResponse, c)
	return
}
