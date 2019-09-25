// 对应配置文件config.json
package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

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

var nowConfig *Config

// 读取当前配置
func Now() *Config {
	return NowFile("")
}

func NowFile(filename string) *Config {
	if filename == "" {
		filename = "./config.json"
	}
	if nowConfig == nil {
		nowConfig = loadConfig(filename)
	}
	return nowConfig
}

// 加载配置
func loadConfig(filename string) *Config {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	config := &Config{}
	err = json.Unmarshal(data, config)
	if err != nil {
		log.Fatalf("解析配置文件'%s'出错 %s", filename, err.Error())
	}
	return config
}
