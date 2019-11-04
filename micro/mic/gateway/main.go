package main

import (
	"fmt"

	"gateway/router"
)

func main() {
	fmt.Print("***************************  Gateway ***************************************")

	// 加载配置信息

	// 启动HTTP服务
	router.RunHttpServer()
}
