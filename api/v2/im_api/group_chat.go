package im_api

import (
	"encoding/json"
	"fmt"
	"gbv2/config/log"
	"gbv2/config/mysql"
	"gbv2/models"
	"gbv2/models/ctype"
	"gbv2/models/res"
	"github.com/DanPlayer/randomname"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
	"time"
)

type ChatUser struct {
	Conn     *websocket.Conn
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
}

var ConnGroupMap = map[string]ChatUser{}

const (
	TextMsg    ctype.MsgType = 1
	ImageMsg   ctype.MsgType = 2
	SystemMsg  ctype.MsgType = 3
	InRoomMsg  ctype.MsgType = 4
	OutRoomMsg ctype.MsgType = 5
)

type GroupRequest struct {
	Content string        `json:"content"`  // 聊天的内容
	MsgType ctype.MsgType `json:"msg_type"` // 聊天类型
}
type GroupResponse struct {
	NickName string        `json:"nick_name"` // 前端自己生成
	Avatar   string        `json:"avatar"`    // 头像
	MsgType  ctype.MsgType `json:"msg_type"`  // 聊天类型
	Content  string        `json:"content"`   // 聊天的内容
	Date     time.Time     `json:"date"`      // 消息的时间
}

func (IMApi) ChatGroupView(c *gin.Context) {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// 鉴权 true表示放行，false表示拦截
			return true
		},
	}
	// 将http升级至websocket
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}
	addr := conn.RemoteAddr().String()
	nickName := randomname.GenerateName()
	nickNameFirst := string([]rune(nickName)[0])
	avatar := fmt.Sprintf("uploads/chat_avatar/%s.png", nickNameFirst)

	chatUser := ChatUser{
		Conn:     conn,
		NickName: nickName,
		Avatar:   avatar,
	}
	ConnGroupMap[addr] = chatUser
	// 需要去生成昵称，根据昵称首字关联头像地址
	// 昵称关联 addr
	log.Infow("%s %s 连接成功", addr, chatUser.NickName)
	for {
		// 消息类型，消息，错误
		_, p, err := conn.ReadMessage()
		if err != nil {
			// 用户断开聊天
			SendGroupMsg(conn, GroupResponse{
				Content: fmt.Sprintf("%s 离开聊天室", chatUser.NickName),
				Date:    time.Now(),
			})
			break
		}
		// 进行参数绑定
		var request GroupRequest
		err = json.Unmarshal(p, &request)
		if err != nil {
			// 参数绑定失败
			SendMsg(addr, GroupResponse{
				NickName: chatUser.NickName,
				Avatar:   chatUser.Avatar,
				MsgType:  SystemMsg,
				Content:  "参数绑定失败",
			})
			continue
		}
		// 判断类型
		switch request.MsgType {
		case TextMsg:
			if strings.TrimSpace(request.Content) == "" {
				SendMsg(addr, GroupResponse{
					NickName: chatUser.NickName,
					Avatar:   chatUser.Avatar,
					MsgType:  SystemMsg,
					Content:  "消息不能为空",
				})
				continue
			}
			SendGroupMsg(conn, GroupResponse{
				NickName: chatUser.NickName,
				Avatar:   chatUser.Avatar,
				Content:  request.Content,
				MsgType:  TextMsg,
				Date:     time.Now(),
			})
		case InRoomMsg:
			SendGroupMsg(conn, GroupResponse{
				NickName: chatUser.NickName,
				Avatar:   chatUser.Avatar,
				Content:  fmt.Sprintf("%s 进入聊天室", chatUser.NickName),
				Date:     time.Now(),
			})
		default:
			SendMsg(addr, GroupResponse{
				NickName: chatUser.NickName,
				Avatar:   chatUser.Avatar,
				MsgType:  SystemMsg,
				Content:  "消息类型错误",
			})
		}

	}
	defer conn.Close()
	delete(ConnGroupMap, addr)
}

// SendGroupMsg 群聊功能
func SendGroupMsg(conn *websocket.Conn, response GroupResponse) {
	byteData, _ := json.Marshal(response)
	_addr := conn.RemoteAddr().String()
	ip, addr := getIPAndAddr(_addr)

	mysql.DB.Create(&models.ChatModel{
		NickName: response.NickName,
		Avatar:   response.Avatar,
		Content:  response.Content,
		IP:       ip,
		Addr:     addr,
		IsGroup:  true,
		MsgType:  response.MsgType,
	})
	for _, chatUser := range ConnGroupMap {
		chatUser.Conn.WriteMessage(websocket.TextMessage, byteData)
	}
}

// SendMsg 给某个用户发消息
func SendMsg(_addr string, response GroupResponse) {
	byteData, _ := json.Marshal(response)
	chatUser := ConnGroupMap[_addr]
	ip, addr := getIPAndAddr(_addr)
	mysql.DB.Create(&models.ChatModel{
		NickName: response.NickName,
		Avatar:   response.Avatar,
		Content:  response.Content,
		IP:       ip,
		Addr:     addr,
		IsGroup:  false,
		MsgType:  response.MsgType,
	})
	chatUser.Conn.WriteMessage(websocket.TextMessage, byteData)
}

func getIPAndAddr(_addr string) (ip string, addr string) {
	addrList := strings.Split(_addr, ":")
	addr = "内网"
	return addrList[0], addr
}
