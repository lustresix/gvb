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

type EmailLoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

func (UserApi) EmailLoginView(c *gin.Context) {
	var cr EmailLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}

	var userModel models.UserModel
	err = mysql.DB.Take(&userModel, "user_name = ? or email = ?", cr.UserName, cr.UserName).Error
	if err != nil {
		// 没找到
		log.Warnw("用户名不存在")
		res.FailWithMsg("用户名或密码错误", c)
		return
	}
	// 校验密码
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		log.Warnw("用户名密码错误")
		res.FailWithMsg("用户名或密码错误", c)
		return
	}
	// 登录成功，生成token
	token, err := jwt.GetToken(jwt.JwtPayLoad{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
	})
	if err != nil {
		log.Errorw("err", "err", err)
		res.FailWithMsg("token生成失败", c)
		return
	}
	res.OKWithData(token, c)

}
