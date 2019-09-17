// 配置读取客户端
// 为项目提供配置信息读取
// 配置读取位置可实现为配置文件、配置中心、数据库表等
package hconfig

import (
	"mic/helper/hmongo"
)

// HTTP 服务相关配置
type Server struct {
	Port int32
}

var server *Server

func init() {
	server = &Server{Port: 50501}
}

func GetServer() *Server {
	return server
}

// 微服务配置信息
type Application struct {
	// 服务名，用于服务标识，微服务通过服务名调用服务接口
	Name string
}

var application *Application

func init() {
	application = &Application{Name: "config"}
}

func GetApplication() *Application {
	return application
}

type Mongo struct {
	// Mongodb 连接字符串
	ConnectionString map[string]string
}

var mongo *Mongo

func init() {
	conns := make(map[string]string)
	conns["test"] = "mongodb://test:test@121.40.83.200:37017/test?authSource=admin&authMechanism=SCRAM-SHA-1"

	mongo = &Mongo{ConnectionString: conns}

	// 初始化Mongo Client
	hmongo.Initialize(mongo.ConnectionString)
}
