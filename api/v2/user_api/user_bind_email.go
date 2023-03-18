package user_api

import (
	"gbv2/config/log"
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/res"
	"gbv2/plugin/email"
	"gbv2/utils/jwt"
	"gbv2/utils/pwd"
	"gbv2/utils/random"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Code     *string `json:"code"`
	Password string  `json:"password"`
}

func (UserApi) UserBindEmailView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	// 用户绑定邮箱， 第一次输入是 邮箱
	// 后台会给这个邮箱发验证码
	var cr BindEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.Errorw("err", "err", err)
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}
	session := sessions.Default(c)
	if cr.Code == nil {
		// 第一次，后台发验证码
		// 生成4位验证码， 将生成的验证码存入session
		code := random.Code(4)
		// 写入session
		session.Set("valid_code", code)
		err = session.Save()
		if err != nil {
			log.Errorw("session错误", "err", err)
			res.FailWithMsg("session错误", c)
			return
		}
		err = email.NewCode().Send(cr.Email, "你的验证码是 "+code)
		if err != nil {
			log.Errorw("验证码发送失败", "err", err)
		}
		res.OKWithMsg("验证码已发送，请查收", c)
		return
	}
	// 第二次，用户输入邮箱，验证码，密码
	code := session.Get("valid_code")
	// 校验验证码
	if code != *cr.Code {
		res.FailWithMsg("验证码错误", c)
		return
	}
	// 修改用户的邮箱
	var user models.UserModel
	err = mysql.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}
	if len(cr.Password) < 4 {
		res.FailWithMsg("密码强度太低", c)
		return
	}
	hashPwd, err := pwd.HashPwd(cr.Password)
	// 第一次的邮箱，和第二次的邮箱也要做一致性校验
	err = mysql.DB.Model(&user).Updates(map[string]any{
		"email":    cr.Email,
		"password": hashPwd,
	}).Error
	if err != nil {
		log.Errorw("绑定邮箱失败", "err", err)
		res.FailWithMsg("绑定邮箱失败", c)
		return
	}
	// 完成绑定
	res.OKWithMsg("邮箱绑定成功", c)
	return
}
