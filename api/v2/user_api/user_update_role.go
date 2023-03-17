package user_api

import (
	"gbv2/config/log"
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/ctype"
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
)

type UserRole struct {
	Role     ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限参数错误"`
	NickName string     `json:"nick_name"` // 防止用户昵称非法，管理员有能力修改
	UserID   uint       `json:"user_id" binding:"required" msg:"用户id错误"`
}

// UserUpdateRoleView 用户权限变更
func (UserApi) UserUpdateRoleView(c *gin.Context) {
	var cr UserRole
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}
	var user models.UserModel
	err := mysql.DB.Take(&user, cr.UserID).Error
	if err != nil {
		res.FailWithMsg("用户id错误，用户不存在", c)
		return
	}
	err = mysql.DB.Model(&user).Updates(map[string]any{
		"role":      cr.Role,
		"nick_name": cr.NickName,
	}).Error
	if err != nil {
		log.Errorw("err", "err", err)
		res.FailWithMsg("修改权限失败", c)
		return
	}
	res.OKWithMsg("修改权限成功", c)
}
