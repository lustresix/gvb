package models

import "gorm.io/gorm"

type LoginDataModel struct {
	gorm.Model
	UserID    uint      `json:"user_id"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
	IP        string    `gorm:"size:20" json:"ip"`
	Nickname  string    `gorm:"size:42" json:"nick_name"`
	Token     string    `gorm:"size:256" json:"token"`
	Device    string    `gorm:"size:256" json:"device"`
	Addr      string    `gorm:"size:64" json:"addr"`
}
