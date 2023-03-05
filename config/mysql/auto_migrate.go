package mysql

import (
	"gbv2/config/log"
	"gbv2/models"
)

func AutoMigrate() {
	DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
	DB.SetupJoinTable(&models.MenuModel{}, "MenuImages", &models.MenuImageModel{})
	err := DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.CommentModel{},
			&models.ArticleModel{},
			&models.UserModel{},
			&models.MenuModel{},
			&models.ImageModel{},
			&models.TagModel{},
			&models.FadeBackModel{},
			&models.MessageModel{},
			&models.AdvertModel{},
			&models.MenuImageModel{},
			&models.LoginDataModel{},
		)
	if err != nil {
		log.Panicw("数据自动迁移失败", err)
	}
	log.Infow("迁移成功")
}
