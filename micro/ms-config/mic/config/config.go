package config

type Config struct {
	Server Server `json:"server"`

	// Mongodb连接配置，key=数据库名 value=连接URL
	Mongos map[string]string
}

type Server struct {
	Port int32 `json:"port"`
}

var _config = new(Config)

// 默认配置初始化
func init() {
	_config.Server.Port = 50508

}

func Get() *Config {
	return _config
}

// 通过配置中心读取配置
func NewConfig() *Config {
	mongos := make(map[string]string)
	mongos["test"] = "mongodb://127.0.0.1:27017/test"
	_config.Mongos = mongos
	return _config
}
