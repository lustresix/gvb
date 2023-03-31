package models

import "time"

type UserCollectModel struct {
	UserID    uint      `gorm:"primaryKey"`
	UserModel UserModel `gorm:"foreignKey:UserID"'`
	ArticleID string    `gorm:"primaryKey"`
	CreatedAt time.Time
}
