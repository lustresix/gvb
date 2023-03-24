package models

import "gorm.io/gorm"

type CommentModel struct {
	gorm.Model
	SubComments        []*CommentModel `gorm:"foreignkey:ParentCommentID" json:"sub_comments"`
	ParentCommentModel *CommentModel   `gorm:"foreignkey:ParentCommentID" json:"parent_comment_model"`
	ParentCommentID    *uint           `json:"parent_comment_id"`
	Content            string          `gorm:"size:256" json:"content"`
	DiggCount          int             `gorm:"size:8;default:0" json:"digg_count"`
	CommentCount       int             `gorm:"size:8;default:0" json:"comment_count"`
	ArticleESID        string          `json:"article_es_id"`
	User               UserModel       `json:"user"`
	UserID             uint            `json:"user_id"`
}
