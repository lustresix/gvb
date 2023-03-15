package routers

import v2 "gbv2/api/v2"

func (router RouterGroup) UserRouter() {
	userApi := v2.ApiGroupApp.UserApi
	router.POST("login", userApi.EmailLoginView)

}
