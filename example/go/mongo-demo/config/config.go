// 配置文件
package config

type Config struct {
	// mongodb配置，key为数据库名，value为数据库连接字符串
	Dbs map[string]string `json:"dbs"`
}

// 配置对象全局变量
var config *Config

func init() {
	config = &Config{}
}

// 设置配置
func Set(conf *Config) {
	config = conf
}

// 获取配置信息结构体
func Get() *Config {
	return config
}
