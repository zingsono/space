// 当前服务配置文件内容
package conf

type Config struct {
	Application *Application `json:"application"`
	Server      *Server      `json:"server"`
	Mongo       *Mongo       `json:"mongo"`
}

type Application struct {
	Name string `json:"name"`
}

type Server struct {
	Port int `json:"port"`
}

// 对应Mongo服务
type Mongo struct {
	Db0 string `json:"db0"`
}

// 默认配置内容
var Now = &Config{
	Application: &Application{Name: "config"},
	Server:      &Server{Port: 10508},
	Mongo:       &Mongo{Db0: "mongodb://test:test@121.40.83.200:37017/test?authSource=admin&authMechanism=SCRAM-SHA-1"},
}

// 初始化配置内容
func init() {
	// 读配置文件

	// 读配置中心

	// 读数据库
}
