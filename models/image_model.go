package models

import (
	"gbv2/models/ctype"
	"gorm.io/gorm"
	"os"
)

type ImageModel struct {
	gorm.Model
	Path      string          `json:"path"`
	Hash      string          `json:"hash"`
	Name      string          `gorm:"size:38" json:"name"`
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"`
}

func (i *ImageModel) BeforeDelete(tx *gorm.DB) error {
	if i.ImageType == ctype.Local {
		err := os.Remove(i.Path + "/" + i.Name)
		return err
	}
	return nil
}
