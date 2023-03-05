package models

import "gorm.io/gorm"

type CommentModel struct {
	gorm.Model
	SubComments        []*CommentModel `gorm:"foreignkey" json:"sub_comments"`
	ParentCommentModel *CommentModel   `gorm:"foreignkey" json:"parent_comment_model"`
	ParentCommentID    *uint           `json:"parent_comment_id"`
	Content            string          `gorm:"size:256" json:"content"`
	DiggCount          int             `gorm:"size:8;default:0" json:"digg_count"`
	CommentCount       int             `gorm:"size:8;default:0" json:"comment_count"`
	Article            ArticleModel    `gorm:"foreignKey:ArticleID " json:"-"`
	ArticleID          uint            `json:"article_id"`
	User               UserModel       `json:"user"`
	UserID             uint            `json:"user_id"`
}
