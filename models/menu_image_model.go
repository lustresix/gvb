package models

type MenuImageModel struct {
	MenuID     uint       `json:"menu_id"`
	MenuModel  MenuModel  `gorm:"foreignKey:MenuID"`
	ImageID    uint       `json:"image_id"`
	ImageModel ImageModel `gorm:"foreignKey:ImageID"`
	Sort       int        `gorm:"size:10" json:"sort"`
}
