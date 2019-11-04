package main

import (
	"log"
	"net/http"
)

/**
服务说明
服务通讯方式：WebSocket


启动参数：
--discover.uris=http://127.0.0.1:50100,http://127.0.0.1:50101,
*/
func main() {

	log.Fatal(http.ListenAndServe("0.0.0.0:50800", nil))

}
