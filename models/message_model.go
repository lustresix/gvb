package models

import "gorm.io/gorm"

type MessageModel struct {
	gorm.Model
	SendUserID       uint      `gorm:"primaryKey" json:"send_user_id"`
	SendUserModel    UserModel `gorm:"foreignKey:SendUserID" json:"-"`
	SendUserNickName string    `gorm:"size:42" json:"send_user_nick_name"`
	SendUserAvatar   string    `json:"send_user_avatar"`

	RevUserID       uint      `gorm:"primaryKey" json:"rev_user_id"`
	RevUserModel    UserModel `gorm:"foreignKey:RevUserID" json:"-"`
	RevUserNickName string    `gorm:"size:42" json:"rev_user_nick_name"`
	RevUserAvatar   string    `json:"rev_user_avatar"`

	IsRead  bool   `gorm:"default:false" json:"is_read"`
	Content string `json:"content"`
}
