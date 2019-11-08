package hws

import (
	"log"
	"sync"

	"golang.org/x/net/websocket"
)

type message struct {
	// 消息唯一编号
	Tid string `json:"tid"`
	// 服务名，根据服务名查询转发目标
	Name string `json:"name"`
	// 消息内容
	Msg string `json:"msg"`
}

type MReq struct {
	message
}

type MRes struct {
	message
}

var (
	// 连接信息
	wsmap sync.Map

	// 消息通道Map
	msgchan sync.Map
)

func GraphqlGateway(ws *websocket.Conn) {
	for {
		var mReq MReq
		err := websocket.JSON.Receive(ws, &mReq)
		if err != nil {
			log.Fatalln(err)
		}
		go func() {
			// TODO mReq验证签名、解密等操作

			// 初始化消息通道
			msgchan.Store(mReq.Tid, make(chan MRes))

			// 调用其它服务
			err := websocket.JSON.Send(connect(mReq.Name), mReq)
			if err != nil {
				log.Fatal(err)
			}

			// 订阅消息通道，回复消息
			v, ok := msgchan.Load(mReq.Tid)
			if !ok {
				log.Fatal("msgchan ok=false mReq.Tid=" + mReq.Tid)
			}
			mRes := <-v.(chan MRes)
			// TODO 消息加密加签名

			// 回复消息
			websocket.JSON.Send(ws, mRes)
		}()
	}
}

func connect(name string) *websocket.Conn {
	// TODO ws连接URL，根据name获取
	url := name

	// 根据url缓存ws连接，如果连接可用直接返回
	v, ok := wsmap.Load(url)
	if ok {
		return v.(*websocket.Conn)
	}
	ws, err := websocket.Dial(url, "", "http://gateway")
	if err != nil {
		log.Fatal(err)
	}
	wsmap.Store(url, ws)

	// 使用新线程订阅消息
	go func(ws *websocket.Conn, url string) {
		for {
			var res MRes
			err := websocket.JSON.Receive(ws, &res)
			if err != nil {
				wsmap.Delete(url)
				log.Fatal(err)
			}
			v, ok := msgchan.Load(res.Tid)
			if ok {
				v.(chan MRes) <- res
			}
		}
	}(ws, url)

	return ws
}
