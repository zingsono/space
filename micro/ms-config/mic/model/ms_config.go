package model

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 配置信息管理
// 字段：服务名、配置JSON内容、备注、更新时间、创建时间
type MsConfig struct {
	Name      string    `json:"name" bson:"name"`
	Value     string    `json:"value" bson:"value"`
	Remark    string    `json:"remark" bson:"remark"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

func (*MsConfig) Collection() *mongo.Collection {
	return DB().Collection("ms_config")
}

// 保存数据，不存在新增，存在则更新
func (conf *MsConfig) Save() (*mongo.UpdateResult, error) {
	conf.UpdatedAt = time.Now().Local()
	log.Println(conf)
	updateResult, err := conf.Collection().UpdateOne(context.TODO(), bson.M{"name": conf.Name}, bson.M{"$set": conf}, options.Update().SetUpsert(true))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// ObjectID 改为string显示
	if updateResult.UpsertedID != nil {
		updateResult.UpsertedID = updateResult.UpsertedID.(primitive.ObjectID).Hex()
	}
	return updateResult, err
}

func (m *MsConfig) FindOne(name string) (*MsConfig, error) {
	log.Printf("执行FindOne")
	singleResult := m.Collection().FindOne(context.Background(), bson.M{"name": name})
	msConfig := &MsConfig{}
	err := singleResult.Decode(msConfig)
	if err != nil {
		return nil, err
	}
	return msConfig, nil
}
