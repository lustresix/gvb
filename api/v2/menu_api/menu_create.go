package menu_api

import (
	"gbv2/config/log"
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/ctype"
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}
type MenuReq struct {
	Title         string      `json:"title" binding:"required" msg:"请输入菜单名称" structs:"title"`
	Path          string      `json:"path" binding:"required" msg:"请输入菜单英文名" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      ctype.Array `json:"abstract" structs:"abstract"`
	AbstractTime  int         `json:"abstract_time" structs:"abstract_time"`
	MenuTime      int         `json:"menu_time" structs:"menu_time"`
	Sort          int         `json:"sort" binding:"required" msg:"请输入菜单序号" structs:"sort"`
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`
}

func (MenuApi) MenuCreatView(c *gin.Context) {
	var cr MenuReq
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.Errorw("err", "err", err)
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}
	var menuModel models.MenuModel
	// 重复值判断
	count := mysql.DB.Find(&menuModel, "menu_title = ? or path = ?", cr.Title, cr.Path).RowsAffected
	if count > 0 {
		res.FailWithMsg("已经在里面了", c)
		return
	}
	// 创建menu数据入库
	menuModel = models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		MenuTime:     cr.MenuTime,
		Sort:         cr.Sort,
	}

	err = mysql.DB.Create(&menuModel).Error

	if err != nil {
		log.Errorw("菜单添加失败", "err", err)
		res.FailWithMsg("菜单添加失败", c)
		return
	}
	if len(cr.ImageSortList) == 0 {
		res.OKWithMsg("菜单添加成功", c)
		return
	}

	var menuBannerList []models.MenuImageModel

	for _, sort := range cr.ImageSortList {
		// 这里也得判断image_id是否真正有这张图片
		menuBannerList = append(menuBannerList, models.MenuImageModel{
			MenuID:  menuModel.ID,
			ImageID: sort.ImageID,
			Sort:    sort.Sort,
		})
	}
	// 给第三张表入库
	err = mysql.DB.Create(&menuBannerList).Error
	if err != nil {
		log.Errorw("菜单图片关联失败", "err", err)
		res.FailWithMsg("菜单图片关联失败", c)
		return
	}
	res.OKWithMsg("菜单添加成功", c)
}
