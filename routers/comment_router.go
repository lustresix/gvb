package routers

import (
	v2 "gbv2/api/v2"
	"gbv2/middleware"
)

func (router RouterGroup) CommentRouter() {
	commentApi := v2.ApiGroupApp.CommentApi
	router.POST("comment", middleware.JwtAuth(), commentApi.CommentCreateView)

}
