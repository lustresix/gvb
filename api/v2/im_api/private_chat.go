package im_api

import (
	"fmt"
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"time"
)

type IMIDRequest struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// CommentRemoveView 一对一聊天
func (IMApi) CommentRemoveView(c *gin.Context) {
	var cr IMIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}
	// 升级为 websocket
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	fmt.Println(err)
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}
	for {
		// 消息类型，消息，错误
		_, p, err := conn.ReadMessage()
		if err != nil {
			// 用户断开聊天
			break
		}
		fmt.Println(string(p))
		// 发送消息
		for i := 0; i < 5; i++ {
			_ = conn.WriteMessage(websocket.TextMessage, []byte("xxx"))
			time.Sleep(1 * time.Second)
		}
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)
}
