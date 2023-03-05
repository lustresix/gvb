package models

import "gorm.io/gorm"

type TagModel struct {
	gorm.Model
	Title    string         `gorm:"size:16" json:"title"`
	Articles []ArticleModel `gorm:"many2many;article_tag" json:"-"`
}
