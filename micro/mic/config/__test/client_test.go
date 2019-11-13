package main

import (
	"encoding/json"
	"log"
	"testing"

	"golang.org/x/net/websocket"
)

func Panic(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// ws:127.0.0.01:5800ws/config  配置服务测试
func TestConfigWsClient(t *testing.T) {

	ws, err := websocket.Dial("ws://127.0.0.01:5800/ws/config?name=gateway", "", "http://127.0.0.1")
	Panic(err)

	err = websocket.Message.Send(ws, `["default","mongo"]`)
	Panic(err)
	for {
		var value map[string]interface{}
		err = websocket.JSON.Receive(ws, &value)
		Panic(err)
		b, err := json.Marshal(value)
		Panic(err)
		log.Print(string(b))

	}
}
