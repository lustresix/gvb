package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code""`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func OK(data any, msg string, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}

func OKWithData(data any, c *gin.Context) {
	Result(SUCCESS, data, "success", c)
}

func OKWithMsg(msg string, c *gin.Context) {
	Result(SUCCESS, map[string]any{}, msg, c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(ERROR, data, msg, c)
}

func FailWithMsg(msg string, c *gin.Context) {
	Result(ERROR, map[string]any{}, msg, c)
}

func FailWithCode(code int, c *gin.Context) {
	Result(ERROR, map[string]any{}, GetMsg(code), c)
}
