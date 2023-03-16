package user_api

import (
	"gbv2/models"
	"gbv2/models/ctype"
	"gbv2/models/res"
	"gbv2/service/common"
	"gbv2/utils/desens"
	"gbv2/utils/jwt"
	"github.com/gin-gonic/gin"
)

func (UserApi) UserListView(c *gin.Context) {
	// 如何判断是管理员
	token := c.Request.Header.Get("token")
	if token == "" {
		res.FailWithMsg("未携带token", c)
		return
	}
	claims, err := jwt.ParseToken(token)
	if err != nil {
		res.FailWithMsg("token错误", c)
		return
	}

	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}
	var users []models.UserModel
	list, count, _ := common.CommonList(models.UserModel{}, common.Option{
		PageInfo: page,
	})
	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			// 管理员
			user.UserName = ""
		}
		user.Tel = desens.DesensitizationTel(user.Tel)
		user.Email = desens.DesensitizationEmail(user.Email)
		// 脱敏
		users = append(users, user)
	}

	res.OKWitList(users, count, c)
}
