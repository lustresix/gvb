package message_api

import (
	"gbv2/config/log"
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
)

type MessageRequest struct {
	SendUserID uint   `json:"send_user_id" binding:"required"` // 发送人id
	RevUserID  uint   `json:"rev_user_id" binding:"required"`  // 接收人id
	Content    string `json:"content" binding:"required"`      // 消息内容
}

// MessageCreateView 发布消息
func (MessageApi) MessageCreateView(c *gin.Context) {
	// 当前用户发布消息
	// SendUserID 就是当前登录人的id
	var cr MessageRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	var senUser, recvUser models.UserModel

	err = mysql.DB.Take(&senUser, cr.SendUserID).Error
	if err != nil {
		res.FailWithMsg("发送人不存在", c)
		return
	}
	err = mysql.DB.Take(&recvUser, cr.RevUserID).Error
	if err != nil {
		res.FailWithMsg("接收人不存在", c)
		return
	}

	err = mysql.DB.Create(&models.MessageModel{
		SendUserID:       cr.SendUserID,
		SendUserNickName: senUser.NickName,
		SendUserAvatar:   senUser.Avatar,
		RevUserID:        cr.RevUserID,
		RevUserNickName:  recvUser.NickName,
		RevUserAvatar:    recvUser.Avatar,
		IsRead:           false,
		Content:          cr.Content,
	}).Error
	if err != nil {
		log.Errorw("err", "err", err)
		res.FailWithMsg("消息发送失败", c)
		return
	}
	res.OKWithMsg("消息发送成功", c)
	return
}
