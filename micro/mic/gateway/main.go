package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("***************************  Gateway ***************************************")
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
		Port: 8610,
	}
}

func StartServer() {
	// 初始化配置
	config = NowConfig()

	// 启动服务
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", config.Port), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)

		// 接口网关，提供接口加密解密、验证签名、token验证
		switch r.RequestURI {
		// Graphql接口网关
		case "/":
			w.Write([]byte("/v1/gateway---------"))
		// 请求转发
		default:
			w.Write([]byte(r.RequestURI))
		}
	}))
	log.Fatal(err)
}
