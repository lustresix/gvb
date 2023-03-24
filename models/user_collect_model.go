package models

import "time"

type UserCollectModel struct {
	UserID    uint      `gorm:"primaryKey"`
	UserModel UserModel `gorm:"foreignKey:UserID"'`
	ArticleID uint      `gorm:"primaryKey"`
	CreatedAt time.Time
}
