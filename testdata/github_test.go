package testdata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

type nun struct {
	Category string `json:"category"`
	Period   string `json:"period"`
	Lang     string `json:"lang"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	Cursor   string `json:"cursor"`
}

type NewData struct {
	Id            string `json:"id"`
	Url           string `json:"url"`
	Username      string `json:"username"`
	Reponame      string `json:"reponame"`
	Description   string `json:"description"`
	Lang          string `json:"lang"`
	LangColor     string `json:"langColor"`
	DetailPageUrl string `json:"detailPageUrl"`
	StarCount     int    `json:"starCount"`
	ForkCount     int    `json:"forkCount"`
	Owner         struct {
		Username string `json:"username"`
		Avatar   string `json:"avatar"`
		Url      string `json:"url"`
	} `json:"owner"`
	Translation struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"translation"`
}

type NewResponse struct {
	Code int       `json:"code"`
	Data []NewData `json:"data"`
}

func TestGithub(T *testing.T) {
	n := nun{
		Category: "trending",
		Period:   "week",
		Lang:     "go",
		Offset:   0,
		Limit:    30,
		Cursor:   "0",
	}
	resp, _ := Post("https://e.juejin.cn/resources/github", n, nil, 2*time.Second)
	byteData, _ := io.ReadAll(resp.Body)
	var response NewResponse
	_ = json.Unmarshal(byteData, &response)
	fmt.Println(response)
}
func Post(url string, data any, headers map[string]interface{}, timeout time.Duration) (resp *http.Response, err error) {
	reqParam, _ := json.Marshal(data)
	reqBody := strings.NewReader(string(reqParam))
	httpReq, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return
	}
	httpReq.Header.Add("Content-Type", "application/json")
	for key, val := range headers {
		switch v := val.(type) {
		case string:
			httpReq.Header.Add(key, v)
		case int:
			httpReq.Header.Add(key, strconv.Itoa(v))
		}
	}
	client := http.Client{
		Timeout: timeout,
	}
	httpResp, err := client.Do(httpReq)
	return httpResp, err
}
