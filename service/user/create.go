package user

import (
	"errors"
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/ctype"
	"gbv2/utils/pwd"
)

const Avatar = "/uploads/avatar/default.png"

func CreateUser(userName, nickName, password string, role ctype.Role, email string, ip string) error {
	// 判断用户名是否存在
	var userModel models.UserModel
	err := mysql.DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		return errors.New("用户名已存在")
	}
	// 对密码进行hash
	hashPwd, _ := pwd.HashPwd(password)

	// 头像问题
	// 1. 默认头像
	// 2. 随机选择头像

	// 入库
	err = mysql.DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashPwd,
		Email:      email,
		Role:       role,
		Avatar:     Avatar,
		IP:         ip,
		Addr:       "内网地址",
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
