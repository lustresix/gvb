package routers

import (
	v2 "gbv2/api/v2"
	"gbv2/middleware"
)

func (router RouterGroup) UserRouter() {
	userApi := v2.ApiGroupApp.UserApi
	router.POST("login", userApi.EmailLoginView)
	router.GET("user", middleware.JwtAuth(), userApi.UserListView)
	router.PUT("user_role", middleware.JwtAuth(), userApi.UserUpdateRoleView)
	router.PUT("user_pwd", userApi.UserUpdatePassword)
}
