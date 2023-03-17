package routers

import (
	v2 "gbv2/api/v2"
	"gbv2/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var store = cookie.NewStore([]byte("HyvCD89g3VDJ9646BFGEh37GFJ"))

func (router RouterGroup) UserRouter() {
	userApi := v2.ApiGroupApp.UserApi
	router.Use(sessions.Sessions("sessionid", store))
	router.POST("login", userApi.EmailLoginView)
	router.GET("user", middleware.JwtAuth(), userApi.UserListView)
	router.PUT("user_role", middleware.JwtAdmin(), userApi.UserUpdateRoleView)
	router.PUT("user_pwd", middleware.JwtAuth(), userApi.UserUpdatePassword)
	router.POST("logout", middleware.JwtAuth(), userApi.LogoutView)
	router.DELETE("user_remove", middleware.JwtAdmin(), userApi.UserRemoveView)
	router.POST("user_bind_email", middleware.JwtAuth(), userApi.UserBindEmailView)
}
