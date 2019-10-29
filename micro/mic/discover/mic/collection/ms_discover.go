package collection

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// 表名：ms_discover
// 字段：实例名（服务唯一标识）、服务名、主机、port、过期时间、更新时间、
//      name   host  port  expires  updatedAt
// 状态（0=准备 1=正常 9=停止）
/**
通信方式：WebSocket
创建连接自动注册且定时发送心跳包、断开连接自动剔除

*/
type MsDiscover struct {
	Instance  string     `json:"instance"`
	Name      string     `json:"name"`
	Host      string     `json:"host"`
	Port      string     `json:"port"`
	Expires   time.Timer `json:"expires"`
	UpdatedAt time.Timer `json:"updatedAt"`
}

func (*MsDiscover) Collection() *mongo.Collection {
	return Db0().Collection("ms_discover")
}

func (m *MsDiscover) Save() {
	return m.Collection().InsertOne(context.TODO(), m)
}
