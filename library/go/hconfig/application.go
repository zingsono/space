package hconfig

// 微服务配置
type Application struct {
	Name   string  `json:"name"`
	Server *Server `json:"server"`
}

// Http Server
type Server struct {
	Port int `json:"port"`
}
