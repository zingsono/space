package mconfig

// 配置文件
type Config struct {
}

// 获取配置对象
func Get() *Config {

	// 从配置中心读取配置信息
	return &Config{}
}
