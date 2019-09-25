// 配置管理
package hconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// 服务配置信息
var config = make(map[string]interface{})

// 解析JSON为Map
func Parse(jsonBytes []byte) map[string]interface{} {
	confMap := make(map[string]interface{})
	err := json.Unmarshal(jsonBytes, confMap)
	if err != nil {
		log.Fatalln(err)
	}
	config = conf
	return config
}

func Get(name string) interface{} {
	return config[name]
}

// 从json字符串读取配置信息
func JsonStringConfig(jsonString string) *HConfig {
	return Parse([]byte(jsonString))
}

// 从json文件读取配置信息
func JsonFileConfig(filename string) *HConfig {
	bytes, err = ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	return Parse(bytes)
}

// 从Graphql服务读取配置信息
func GraphConfig(name []string) *HConfig {
	ql := fmt.Sprintf(
		`query ConfigValue 
		{
			value(name:"%s")	
		}	
		`,
		name)
	resp, err := http.Post("/graphql", "text/plain", strings.NewReader(ql))
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("GraphConfig Server Response code %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return Parse(body)
}

// 从http服务读取配置信息

// 从yaml读取配置信息

// 从toml读取配置信息

// 从git读取配置信息
