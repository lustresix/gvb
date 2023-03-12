package common

import (
	"gbv2/config/mysql"
	"gbv2/models"
)

type Option struct {
	models.PageInfo
}

func CommonList[T any](model T, option Option) (list []T, count int64, err error) {
	query := mysql.DB.Where(model)
	count = query.Select("id").Find(&list).RowsAffected
	query = mysql.DB.Where(model)

	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	if option.Sort == "" {
		option.Sort = "created_at desc"
	}

	err = query.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	return list, count, err
}
