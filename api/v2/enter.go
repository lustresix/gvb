package v2

import (
	"gbv2/api/v2/advert_api"
	"gbv2/api/v2/article_api"
	"gbv2/api/v2/comment_api"
	"gbv2/api/v2/digg_api"
	"gbv2/api/v2/gpt_api"
	"gbv2/api/v2/images_api"
	"gbv2/api/v2/menu_api"
	"gbv2/api/v2/message_api"
	"gbv2/api/v2/setting_api"
	"gbv2/api/v2/tag_api"
	"gbv2/api/v2/user_api"
)

type ApiGroup struct {
	SettingsApi setting_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
	MenuApi     menu_api.MenuApi
	UserApi     user_api.UserApi
	TagApi      tag_api.TagApi
	MessageApi  message_api.MessageApi
	ArticleApi  article_api.ArticleApi
	GPTApi      gpt_api.GptApi
	DiggApi     digg_api.DiggApi
	CommentApi  comment_api.CommentApi
}

var ApiGroupApp = new(ApiGroup)
