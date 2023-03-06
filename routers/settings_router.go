package routers

import v2 "gbv2/api/v2"

func (router RouterGroup) SettingRouter() {
	settingsApi := v2.ApiGroupApp.SettingsApi
	router.GET("settings", settingsApi.SettingsInfoView)
	router.PUT("settings", settingsApi.SettingsInfoUpdateView)
}
