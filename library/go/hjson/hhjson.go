// JSON
package hjson

import (
	"encoding/json"
)

// ToString 结构体对象转JSON字符串
// v 是struct或者map或者数组
func ToJson(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if e != nil {
		return "", err
	}
	return string(bytes), nil
}

// ToMap JSON转Map
// data是json字符串，返回值是转换结果
func ToMap(v string) (map[string]interface{}, error) {
	var rs map[string]interface{}
	e := json.Unmarshal([]byte(v), rs)
	if e != nil {
		return nil, e
	}
	return rs, nil
}

// ToArray JSON转数组
func ToArray() {

}

// ToStruct JSON转结构体
func ToStruct(data string, obj interface{}) error {
	e := json.Unmarshal([]byte(data), obj)
	return e
}
