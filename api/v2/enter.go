package v2

import (
	"gbv2/api/v2/images_api"
	"gbv2/api/v2/setting_api"
)

type ApiGroup struct {
	SettingsApi setting_api.SettingsApi
	ImagesApi   images_api.ImagesApi
}

var ApiGroupApp = new(ApiGroup)
