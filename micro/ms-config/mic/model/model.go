package model

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"mic/conf"
)

var (
	database     *mongo.Database
	databaseChan = make(chan *mongo.Database)
)

func Connect() *mongo.Database {
	if database == nil {
		database = <-databaseChan
	}
	return database

}

func init() {
	go func() {
		connectionString := conf.Now.Mongo.Db0
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		options := options.Client().ApplyURI(connectionString)
		options.SetMinPoolSize(3)
		options.SetMaxConnIdleTime(30 * time.Second)

		client, err := mongo.Connect(ctx, options)
		if err != nil {
			log.Fatalf("Error Connect Mongodb Fail...")
		}
		database = client.Database("test")
		log.Printf("Mongodb connect success %s %s", "db0", connectionString)

		databaseChan <- database
	}()
}

/*
// 数据库连接
// 多数据库连接时，新增一个获取函数
var DB = func() *mongo.Database {
	// 等于空时，加锁，解锁后重新判断是否等于空
	if database ==nil{
		databaseMutex.Lock()
		defer databaseMutex.Unlock()
		log.Printf("获取Lock")
	}
	if database != nil {
		log.Printf("res database...........")
		return database
	}
	connectionString := conf.Now.Mongo.Db0
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	options := options.Client().ApplyURI(connectionString)
	options.SetMinPoolSize(3)
	options.SetMaxConnIdleTime(30*time.Second)

	client, err := mongo.Connect(ctx, options)
	if err != nil {
		log.Fatalf("Error Connect Mongodb Fail...")
	}
	database = client.Database("test")
	log.Printf("Mongodb connect success %s %s", "db0", connectionString)
	return database
}
*/
