package v2

import "gbv2/api/v2/setting_api"

type ApiGroup struct {
	SettingsApi setting_api.SettingsApi
}

var ApiGroupApp = new(ApiGroup)
