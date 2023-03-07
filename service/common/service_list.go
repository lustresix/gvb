package common

import (
	"gbv2/config/mysql"
	"gbv2/models"
)

type Option struct {
	models.PageInfo
}

func CommonList[T any](model T, option Option) (list []T, count int64, err error) {
	count = mysql.DB.Select("id").Find(&list).RowsAffected
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	if option.Sort == "" {
		option.Sort = "created_at desc"
	}
	err = mysql.DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	return list, count, err
}
