package models

import "time"

type CollectModel struct {
	UserID       uint         `gorm:"primaryKay"`
	UserModel    UserModel    `gorm:"foreignKey:UserID"'`
	ArticleID    uint         `gorm:"primaryKey"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID"`
	CreatedAt    time.Time
}
