package models

import (
	"gbv2/models/ctype"
	"gorm.io/gorm"
)

type MenuModel struct {
	gorm.Model
	Title        string       `gorm:"size:32" json:"title"`
	Path         string       `gorm:"size:32" json:"path"`
	Slogan       string       `gorm:"size:64" json:"slogan"`
	Abstract     ctype.Array  `gorm:"type:string" json:"abstract"`
	AbstractTime int          `json:"abstract_time"`
	MenuImages   []ImageModel `gorm:"many2many:menu_image_models;joinForeignKey:MenuID;JoinReferences:ImageID" json:"menu_images"`
	MenuTime     int          `json:"menu_time"`
	Sort         int          `gorm:"size:10" json:"sort"`
}
