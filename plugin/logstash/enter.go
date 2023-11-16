package logstash

import (
	"gbv2/config/log"
	"gbv2/config/mysql"
	"gbv2/utils/jwt"
	"github.com/gin-gonic/gin"
)

type Log struct {
	ip     string `json:"ip"`
	addr   string `json:"addr"`
	userId uint   `json:"user_id"`
}

func New(ip string, token string) *Log {
	// 解析token
	claims, err := jwt.ParseToken(token)
	var userID uint
	if err == nil {
		userID = claims.UserID
	}

	// 拿到用户id
	return &Log{
		ip:     ip,
		addr:   "内网",
		userId: userID,
	}
}

func NewLogByGin(c *gin.Context) *Log {
	ip := c.ClientIP()
	token := c.Request.Header.Get("token")
	return New(ip, token)
}

func (l Log) Debug(content string) {
	l.send(DebugLevel, content)
}
func (l Log) Info(content string) {
	l.send(InfoLevel, content)
}
func (l Log) Warn(content string) {
	l.send(WarnLevel, content)
}
func (l Log) Error(content string) {
	l.send(ErrorLevel, content)
}

func (l Log) send(level Level, content string) {
	err := mysql.DB.Create(&LogStashModel{
		IP:      l.ip,
		Addr:    l.addr,
		Level:   level,
		Content: content,
		UserID:  l.userId,
	}).Error
	if err != nil {
		log.Errorw(err.Error(), "err", err)
	}
}
