package models

import (
	"gbv2/models/ctype"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	NickName   string           `gorm:"size:36" json:"nick_name"`            // 昵称
	UserName   string           `gorm:"size:36" json:"user_name"`            // 用户名
	Password   string           `gorm:"size:128" json:"password"`            // 密码
	Avatar     string           `gorm:"size:256" json:"avatar"`              // 头像id
	Email      string           `gorm:"size:128" json:"email"`               // 邮箱
	Tel        string           `gorm:"size:18" json:"tel"`                  // 手机号
	Addr       string           `gorm:"size:64" json:"addr"`                 // 地址
	Token      string           `gorm:"size:64" json:"token"`                // 唯一id
	IP         string           `gorm:"size:20" json:"ip"`                   // ip地址
	Role       ctype.Role       `gorm:"4;default:1" json:"role"`             // 权限
	SignStatus ctype.SignStatus `gorm:"type=smallint(6)" json:"sing_status"` // 注册来源
}
