package models

import "gorm.io/gorm"

type AdvertModel struct {
	gorm.Model
	Title  string `gorm:"size:32" json:"title"`
	Href   string `json:"href"`
	Images string `json:"images"`
	IsShow bool   `json:"is_show"`
}
