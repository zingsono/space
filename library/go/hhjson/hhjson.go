// JSON
package hhjson

import (
	"encoding/json"
	"log"
)

// ToString 结构体对象转JSON字符串
// v 是struct或者map
func ToJson(v interface{}) string {
	bytes,e := json.Marshal(v)
	if e != nil {
		log.Panicf("ToJson Error %s",e.Error())
	}
	return string(bytes)
}


// ToMap JSON转Map
// data是json字符串，返回值是转换结果
func ToMap(data string) map[string]interface{} {
	var rs map[string]interface{}
	e := json.Unmarshal([]byte(data),rs)
	if e != nil {
		log.Panicf("ToStruct Error %s",e.Error())
	}
	return rs
}

// ToStruct
// v 是struct或者map
func ToStruct(data string,v interface{}) {
	e := json.Unmarshal([]byte(data),v)
	if e != nil {
		log.Panicf("ToStruct Error %s",e.Error())
	}
}
