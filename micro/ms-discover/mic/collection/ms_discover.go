package collection

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// 表名：ms_discover
// 字段：服务名、实例名、IP、port、过期时间、更新时间、
//      name   ip  port  expires  updatedAt
// 状态（0=准备 1=正常 9=停止）
type MsDiscover struct {
	Name      string     `json:"name"`
	Instance  string     `json:"instance"`
	Ip        string     `json:"ip"`
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
