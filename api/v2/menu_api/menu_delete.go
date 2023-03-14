package menu_api

import (
	"fmt"
	"gbv2/config/log"
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (MenuApi) MenuRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}

	var menuList []models.MenuModel
	count := mysql.DB.Find(&menuList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("菜单不存在", c)
		return
	}

	// 事务
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		err = mysql.DB.Model(&menuList).Association("MenuImages").Clear()
		if err != nil {
			log.Errorw("err", "err", err)
			return err
		}
		err = mysql.DB.Delete(&menuList).Error
		if err != nil {
			log.Errorw("err", "err", err)
			return err
		}
		return nil
	})
	if err != nil {
		log.Errorw("err", "err", err)
		res.FailWithMsg("删除菜单失败", c)
		return
	}
	res.OKWithMsg(fmt.Sprintf("共删除 %d 个菜单", count), c)

}
