package menu_api

import (
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
)

type Image struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

type MenuResponse struct {
	models.MenuModel
	Images []Image `json:"images"`
}

func (MenuApi) MenuListView(c *gin.Context) {
	// 先查菜单
	var menuList []models.MenuModel
	var menuIDList []uint
	mysql.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList)
	// 查连接表
	var menuImages []models.MenuImageModel
	mysql.DB.Preload("ImageModel").Order("sort desc").Find(&menuImages, "menu_id in ?", menuIDList)
	var menus []MenuResponse
	for _, model := range menuList {
		// model就是一个菜单
		images := []Image{}
		for _, image := range menuImages {
			if model.ID != image.MenuID {
				continue
			}
			images = append(images, Image{
				ID:   image.ImageID,
				Path: image.ImageModel.Path,
				Name: image.ImageModel.Name,
			})
		}
		menus = append(menus, MenuResponse{
			MenuModel: model,
			Images:    images,
		})
	}
	res.OKWithData(menus, c)
	return
}
