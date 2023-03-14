package menu_api

import (
	"gbv2/config/log"
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr MenuReq
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMsg("err", c)
		return
	}
	id := c.Param("id")

	// 先把之前的Image清空
	var menuModel models.MenuModel
	err = mysql.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMsg("菜单不存在", c)
		return
	}
	mysql.DB.Model(&menuModel).Association("MenuImages").Clear()
	// 如果选择了Image，那就添加
	if len(cr.ImageSortList) > 0 {
		// 操作第三张表
		var imageList []models.MenuImageModel
		for _, sort := range cr.ImageSortList {
			imageList = append(imageList, models.MenuImageModel{
				MenuID:  menuModel.ID,
				ImageID: sort.ImageID,
				Sort:    sort.Sort,
			})
		}
		err = mysql.DB.Create(&imageList).Error
		if err != nil {
			log.Errorw("err", "err", err)
			res.FailWithMsg("创建菜单图片失败", c)
			return
		}
	}

	// 普通更新
	maps := structs.Map(&cr)
	err = mysql.DB.Model(&menuModel).Updates(maps).Error

	if err != nil {
		log.Errorw("err", "err", err)
		res.FailWithMsg("修改菜单失败", c)
		return
	}

	res.OKWithMsg("修改菜单成功", c)

}
