package hws

import (
	"log"

	"golang.org/x/net/websocket"
)

type M struct {
	// 消息唯一编号
	Tid string `json:"tid"`
	// 服务名，根据服务名查询转发目标
	Name string `json:"name"`
	// 消息内容
	Msg string `json:"msg"`
}

type MReq struct {
	M
}

type MRes struct {
	M
}

// 连接信息，key为连接URL字符串
var wsmap = make(ConcurrentMap, 32)

// 消息ID对应响应消息通道
var reschanmap = make(ConcurrentMap, 64)

// 发送消息，接收响应
func (m *MReq) Invoke() MRes {
	reschanmap.Put(m.Tid, make(chan MRes))
	m.Send()
	msg := <-reschanmap.Get(m.Tid).(chan MRes)
	reschanmap.Del(m.Tid)
	return msg
}

func (m *MReq) Send() {
	err := websocket.JSON.Send(m.Connect(), m)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *MReq) Connect() *websocket.Conn {
	// TODO ws连接URL，根据name获取
	url := m.Name

	// 根据url缓存ws连接，如果连接可用直接返回
	if wsmap.Get(url) != nil {
		return wsmap.Get(url).(*websocket.Conn)
	}
	ws, err := websocket.Dial(url, "", "http://gateway")
	if err != nil {
		log.Fatal(err)
	}
	// 使用新线程订阅消息
	go subscribe(ws, url)
	return ws
}

func subscribe(ws *websocket.Conn, url string) {
	wsmap.Put(url, ws)
	for {
		var res MRes
		err := websocket.JSON.Receive(ws, &res)
		if err != nil {
			wsmap.Del(url)
			log.Fatal(err)
		}
		reschanmap.Get(res.Tid).(chan MRes) <- res
	}
}

// 网关调用接口方法
func WsGraphqlGateway(ws *websocket.Conn) {
	log.Printf("RequestURI %s", ws.Request().RequestURI)
	for {
		// 接收JSON格式Hws结构体消息
		var req *MReq
		err := websocket.JSON.Receive(ws, &req)
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			err = websocket.JSON.Send(ws, req.Invoke())
			if err != nil {
				log.Fatal(err)
			}
		}()
	}
}
