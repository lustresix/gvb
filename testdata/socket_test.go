package testdata

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"testing"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func TestSocket(t *testing.T) {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	message := "Hello, world!"
	for _, char := range message {
		err = conn.WriteMessage(websocket.TextMessage, []byte(string(char)))
		if err != nil {
			fmt.Println(err)
			return
		}
		// 每个字符发送后等待100毫秒
		time.Sleep(100 * time.Millisecond)
	}
}
