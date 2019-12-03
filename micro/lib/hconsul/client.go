package hconsul

import (
	"log"
	"os"

	"github.com/hashicorp/consul/api"
)

/*

Consul 服务注册客户端
1. 环境变量配置注册中心Consul服务端地址
2. 启动服务时，注册服务
3. 服务停止时，剔除服务

*/

func init() {
	// HTTP代理
	os.Setenv("HTTP_PROXY", "211.152.57.29:39083")

	// 服务注册地址
	if addr := os.Getenv(api.HTTPAddrEnvName); addr == "" {
		os.Setenv(api.HTTPAddrEnvName, "10.18.254.251:8500")
	}
}

func Reg(host string, ip string) {
	log.Print("EXEC............")

	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		log.Panic(err)
	}
	connect := client.Connect()

	log.Print(connect)

	/*intentions, queryMeta, err:=connect.Intentions(&api.QueryOptions{Datacenter: "dc1"})
	if err != nil {
		log.Panic(err)
	}
	log.Print(queryMeta)
	for _,intention := range intentions {
		log.Print(intention.String())
		log.Print(intention.SourceString())
		log.Print(intention.DestinationString())
	}*/

}
