// 根据服务名调用服务
package hfeign

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"
)

type Feign struct {
	// 消息唯一编号
	Tid string `json:"tid"`
	// 服务名，根据服务名查询转发目标
	Name string `json:"name"`
	// 消息内容
	Msg string `json:"msg"`
}

// 调用服务
func (f *Feign) Call() {
	// 根据服务名获取连接
	url := f.Name
	request, err := http.NewRequestWithContext(context.TODO(), "POST", url, bytes.NewBufferString(""))
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Add("Content-Type", "application/json")
	defer request.Body.Close()

	client := &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)

	http.Post()
}
