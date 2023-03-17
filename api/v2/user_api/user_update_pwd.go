package user_api

import (
	"gbv2/config/log"
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/res"
	"gbv2/utils/jwt"
	"gbv2/utils/pwd"
	"github.com/gin-gonic/gin"
)

type UpdatePasswordRequest struct {
	OldPwd string `json:"old_pwd"` // 旧密码
	Pwd    string `json:"pwd"`     // 新密码
}

// UserUpdatePassword 修改登录人的id
func (UserApi) UserUpdatePassword(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var cr UpdatePasswordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}
	var user models.UserModel
	err := mysql.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}
	// 判断密码是否一致
	if !pwd.CheckPwd(user.Password, cr.OldPwd) {
		res.FailWithMsg("密码错误", c)
		return
	}
	hashPwd, _ := pwd.HashPwd(cr.Pwd)
	err = mysql.DB.Model(&user).Update("password", hashPwd).Error
	if err != nil {
		log.Errorw("err", "err", err)
		res.FailWithMsg("密码修改失败", c)
		return
	}
	res.OKWithMsg("密码修改成功", c)
	return
}
