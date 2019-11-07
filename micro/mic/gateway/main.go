package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func mainBak() {
	fmt.Print("***************************  Gateway ***************************************")

	// 加载配置信息

	// 启动HTTP服务
	//router.RunHttpServer()
}

func main() {
	StartServer()
}

// 全局变量定义
var (
	config *Config
)

type Config struct {
	Port                  int
	MongoConnectionString string
}

// 设置配置对象
func NowConfig() *Config {
	return &Config{
		Port:                  8610,
		MongoConnectionString: "mongodb://unionlive:unionlive@proxy.unionlive.com:27017/unionlive?authSource=admin&authMechanism=SCRAM-SHA-1",
	}
}

func StartServer() {
	fmt.Println("Start server ......")
	// 初始化配置
	config = NowConfig()

	// 默认首页
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("/"))
	})

	hws := websocket.Codec{
		Marshal: func(v interface{}) (data []byte, payloadType byte, err error) {
			rid := "12312312"
			return
		},
		Unmarshal: func(data []byte, payloadType byte, v interface{}) (err error) {
			rid := "12312312"
			return
		},
	}

	client, _ := websocket.Dial("", "", "")

	// 消息订阅
	http.Handle("/ws/graphql", websocket.Handler(func(ws *websocket.Conn) {
		log.Printf("RequestURI %s", ws.Request().RequestURI)
		for {
			var v string
			err := websocket.Message.Receive(ws, &v)
			if err != nil {
				log.Fatal(err)
			}
			// 成功订阅到消息，启动新线程处理
			go func(value string, conn *websocket.Conn) {

				// 作为客户端发送消息
				hws.Send(client, value)
				var resValue string
				hws.Receive(client, &resValue)

				// 响应来源请求
				websocket.Message.Send(conn, resValue)

			}(v, ws)
		}
	}))

	// 启动服务
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", config.Port), nil))
}
