package config

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
}

// Mongodb客户端集合，Key为连接名（用于多数据源的情况），value为Client
var dbs map[string]*mongo.Database

// Mongodb Client
func init() {
	dbs = make(map[string]*mongo.Database)
	// TODO 从命令行参数读取mongodb连接字符串
	connectionString := "mongodb://test:test@121.40.83.200:37017/test?authSource=admin&authMechanism=SCRAM-SHA-1"
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatalf("连接Mongodb失败...")
	}
	dbs["ms"] = client.Database("test")
}

func (mongo *Mongo) Ms() *mongo.Database {
	return dbs["ms"]
}
