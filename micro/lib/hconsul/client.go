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

// Consul Client
func GetClient() *api.Client {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		log.Panic(err)
	}
	return client
}

func Register(datacenter, node, address string) {
	log.Print("EXEC Register............")

	if datacenter == "" {
		datacenter = "dc1"
	}
	catalog := GetClient().Catalog()
	_, err := catalog.Register(&api.CatalogRegistration{
		ID:              "",
		Node:            node,
		Address:         address,
		TaggedAddresses: nil,
		NodeMeta:        nil,
		Datacenter:      datacenter,
		Service: &api.AgentService{
			Kind:              "",
			ID:                "s1",
			Service:           "svf",
			Tags:              nil,
			Meta:              nil,
			Port:              8090,
			Address:           "10.18.0.11",
			TaggedAddresses:   nil,
			Weights:           api.AgentWeights{},
			EnableTagOverride: false,
			CreateIndex:       0,
			ModifyIndex:       0,
			ContentHash:       "",
			Proxy:             nil,
			Connect:           nil,
		},
		Check:          nil,
		Checks:         nil,
		SkipNodeUpdate: false,
	}, nil)
	_, err = catalog.Register(&api.CatalogRegistration{
		ID:              "",
		Node:            node,
		Address:         address,
		TaggedAddresses: nil,
		NodeMeta:        nil,
		Datacenter:      datacenter,
		Service: &api.AgentService{
			Kind:              "",
			ID:                "s2",
			Service:           "svf",
			Tags:              nil,
			Meta:              nil,
			Port:              8090,
			Address:           "10.18.0.12",
			TaggedAddresses:   nil,
			Weights:           api.AgentWeights{},
			EnableTagOverride: false,
			CreateIndex:       0,
			ModifyIndex:       0,
			ContentHash:       "",
			Proxy:             nil,
			Connect:           nil,
		},
		Check:          nil,
		Checks:         nil,
		SkipNodeUpdate: false,
	}, nil)
	if err != nil {
		log.Panic(err)
	}

}
