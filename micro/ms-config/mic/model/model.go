package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"mic/conf"
)

var database *mongo.Database

// 数据库连接
// 多数据库连接时，新增一个获取函数
var DB = func() *mongo.Database {
	if database != nil {
		return database
	}
	connectionString := conf.Now().Mongo.Db0
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	var options = options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		log.Fatalf("Error Connect Mongodb Fail...")
	}
	database = client.Database("test")
	log.Printf("Mongo load %s => %s", k, v)
	return database
}
