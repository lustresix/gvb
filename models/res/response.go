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

type ListResponse[T any] struct {
	Count int64 `json:"count"`
	List  T     `json:"list"`
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
func OKWitList(list any, count int64, c *gin.Context) {
	OKWithData(ListResponse[any]{
		List:  list,
		Count: count,
	}, c)
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
