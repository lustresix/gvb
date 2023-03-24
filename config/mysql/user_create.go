package mysql

import (
	"fmt"
	"gbv2/config/log"
	"gbv2/models"
	"gbv2/models/ctype"
	"gbv2/utils/pwd"
)

func CreateUser(permissions string) {
	// 创建用户的逻辑
	// 用户名 昵称 密码 确认密码 邮箱
	userName := ""
	nickName := ""
	password := ""
	rePassword := ""
	email := ""
	fmt.Printf("请输入用户名:")
	fmt.Scanln(&userName)
	fmt.Printf("请输入昵称:")
	fmt.Scanln(&nickName)
	fmt.Printf("请输入邮箱:")
	fmt.Scanln(&email)
	fmt.Printf("请输入密码:")
	fmt.Scanln(&password)
	fmt.Printf("请再次输入密码:")
	fmt.Scanln(&rePassword)

	// 判断用户名是否存在
	var userModel models.UserModel
	err := DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		// 存在
		log.Errorw("用户名已存在，请重新输入", "err", err)
		return
	}
	// 校验两次密码
	if password != rePassword {
		log.Errorw("两次密码不一致，请重新输入", "err", err)
		return
	}
	// 对密码进行hash
	hashPwd, err := pwd.HashPwd(password)
	if err != nil {
		return
	}

	role := ctype.PermissionUser
	if permissions == "admin" {
		role = ctype.PermissionAdmin
	}

	// 头像问题
	// 1. 默认头像
	// 2. 随机选择头像
	avatar := "/uploads/avatar/default.jpg"

	// 入库
	err = DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashPwd,
		Email:      email,
		Role:       role,
		Avatar:     avatar,
		IP:         "127.0.0.1",
		Addr:       "内网地址",
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		log.Errorw("err", "err", err)
		return
	}
	log.Infow("用户" + userName + "创建成功!")

}
