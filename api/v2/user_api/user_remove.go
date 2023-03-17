package user_api

import (
	"fmt"
	"gbv2/config/log"
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/res"
	"gbv2/utils/jwt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (UserApi) UserRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	// 不能删除自己
	for i, s := range cr.IDList {
		if claims.UserID == s {
			cr.IDList[i] = 0
		}
	}

	var userList []models.UserModel
	count := mysql.DB.Find(&userList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("用户不存在", c)
		return
	}

	// 事务
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		// TODO:删除用户，消息表，评论表，用户收藏的文章，用户发布的文章

		err = mysql.DB.Delete(&userList).Error
		if err != nil {
			log.Errorw("删除事物失败", "err", err)
			return err
		}
		return nil
	})
	if err != nil {
		log.Errorw("删除失败", "err", err)
		res.FailWithMsg("删除用户失败", c)
		return
	}
	res.OKWithMsg(fmt.Sprintf("共删除 %d 个用户", count), c)
}
