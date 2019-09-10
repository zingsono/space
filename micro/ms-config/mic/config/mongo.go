package config

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongodb结构体
type Mongo struct {
	Url string
}

// 全局Mongodb客户端
var MongoClient *mongo.Client

// Mongodb Client
func init() {
	connectionString := "mongodb://test:test@121.40.83.200:37017/test?authSource=admin&authMechanism=SCRAM-SHA-1"
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatalf("连接Mongodb失败...")
	}
	MongoClient = client
}
