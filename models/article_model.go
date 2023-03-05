package models

import (
	"gbv2/models/ctype"
	"gorm.io/gorm"
)

type ArticleModel struct {
	gorm.Model
	Title         string         `gorm:"size:32" json:"title"`
	Abstract      string         `json:"abstract"`
	Content       string         `json:"content"`
	LookCount     int            `json:"look_count"`
	CommentCount  int            `json:"comment_count"`
	DiggCount     int            `json:"digg_count"`
	CollectsCount int            `json:"collects_count"`
	TagModel      []TagModel     `gorm:"many2many:article_tag_models" json:"tag_models"`
	CommentModel  []CommentModel `gorm:"foreignKey:ArticleID" json:"-"`
	UserModel     UserModel      `gorm:"foreignKey:UserID" json:"-"`
	UserID        uint           `json:"user_id"`
	Category      string         `gorm:"size:20" json:"category"`
	Source        string         `json:"source"`
	Link          string         `json:"link"`
	Cover         ImageModel     `gorm:"foreignKey:CoverID" json:"-"`
	CoverID       uint           `json:"cover_id"`
	Nickname      string         `gorm:"size:42" json:"nickname"`
	CoverPath     string         `json:"cover_path"`
	Tags          ctype.Array    `gorm:"type:string;size64" json:"tags"`
}
