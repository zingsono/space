package hconfig

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 测试Json读取配置文件
func TestJsonConfig(t *testing.T) {
	s := `
		{
		  "application": {
			"name": "config"
		  },
		  "server": {
			"port": 50501,
			"host": "0.0.0.0"
		  },
		  "mongo": {
			  "test":"mongodb://test:test@121.40.83.200:37017/test?authSource=admin&authMechanism=SCRAM-SHA-1"
		  }	
		}
	`
	hConf := JsonConfig(s)
	bytes, err := json.Marshal(hConf)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println("JSON：", string(bytes))
}
