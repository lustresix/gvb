package routers

import (
	v2 "gbv2/api/v2"
)

func (router RouterGroup) AdvertRouter() {
	AdvertApi := v2.ApiGroupApp.AdvertApi
	router.POST("adverts", AdvertApi.AdvertCreateView)
}
