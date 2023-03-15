package v2

import (
	"gbv2/api/v2/advert_api"
	"gbv2/api/v2/images_api"
	"gbv2/api/v2/menu_api"
	"gbv2/api/v2/setting_api"
	"gbv2/api/v2/user_api"
)

type ApiGroup struct {
	SettingsApi setting_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
	MenuApi     menu_api.MenuApi
	UserApi     user_api.UserApi
}

var ApiGroupApp = new(ApiGroup)
