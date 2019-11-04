// Go WebSocket Server
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	fmt.Println("*********************************************************************************************")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("WebSocket Server "))
	})
	http.Handle("/ws", websocket.Handler(func(ws *websocket.Conn) {
		log.Println(ws.RemoteAddr().String())
		for {
			var v string
			err := websocket.Message.Receive(ws, &v)
			if err != nil {
				log.Print(err)
				break
			}
			log.Printf("接收到消息：%s", v)

			go func() {
				// 处理接收的报文，做响应
				websocket.Message.Send(ws, v)
			}()

		}
	}))
	ListenAndServe()
}

// 生成32位随机字符串
func Id32() string {
	s := time.Now().String() + fmt.Sprintf("%d", rand.Int()) + fmt.Sprintf("%d", rand.Int())
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

// 启动HTTP服务监听
func ListenAndServe() {
	port := 5809
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	log.Printf("Start server %s", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
