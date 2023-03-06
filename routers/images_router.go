package routers

import v2 "gbv2/api/v2"

func (router RouterGroup) ImagesRouter() {
	app := v2.ApiGroupApp.ImagesApi
	router.POST("images", app.ImageUploadView)
}
