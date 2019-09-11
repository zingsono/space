package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"mic/config"
)

// 配置信息管理
// 字段：服务名、配置JSON内容、备注、更新时间、创建时间
type MsConfig struct {
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	Remark    string    `json:"remark"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (*MsConfig) Collection() *mongo.Collection {
	return config.Mongo{}.Ms().Collection("ms_config")
}

// 保存数据，不存在新增，存在则更新
func (conf *MsConfig) Save() (int64, error) {
	conf.UpdatedAt = time.Now()
	updateResult, err := conf.Collection().UpdateOne(context.TODO(), bson.D{{"name", conf.Name}}, conf, options.UpdateOptions{}.SetUpsert(true))
	return updateResult.ModifiedCount, err
}
