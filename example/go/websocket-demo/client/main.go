// websocket client
package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/net/websocket"
)

func main() {

	ws, err := websocket.Dial("ws://127.0.0.1:5809/ws", "", "http://127.0.0.1")
	if err != nil {
		log.Fatalln(err)
	}
	go func() {
		for {
			websocket.Message.Send(ws, "client message send.")
			time.Sleep(5 * time.Second)
		}
	}()

	for {
		var v string
		err := websocket.Message.Receive(ws, &v)
		if err != nil {
			log.Print(err)
			break
		}
		fmt.Println("接收值：" + v)
	}

}
