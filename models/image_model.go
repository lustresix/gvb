package models

import "gorm.io/gorm"

type ImageModel struct {
	gorm.Model
	Path string `json:"path"`
	Hash string `json:"hash"`
	Name string `gorm:"size:38" json:"name"`
}
