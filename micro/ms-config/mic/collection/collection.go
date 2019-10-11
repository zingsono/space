package collection

import (
	"context"
	"log"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"mic/conf"
)

func Connect(connectionString string) *mongo.Database {
	dbName := (strings.Split((strings.Split(connectionString, "/"))[3], "?"))[0]
	if dbName == "" {
		log.Fatalf("Errror Mongodb connectionString %s", connectionString)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	options := options.Client().ApplyURI(connectionString)
	options.SetMinPoolSize(3)
	options.SetMaxConnIdleTime(30 * time.Second)

	client, err := mongo.Connect(ctx, options)
	if err != nil {
		log.Fatalf("Error Connect Mongodb Fail...")
	}
	database := client.Database(dbName)
	log.Printf("Mongodb connect success %s %s", "db0", connectionString)
	return database
}

// 数据库连接全局缓存
var db0 *mongo.Database

// 获取数据库连接
func Db0() *mongo.Database {
	if db0 != nil {
		return db0
	}
	// mongodb://test:test@121.40.83.200:37017/test?authSource=admin&authMechanism=SCRAM-SHA-1
	connectionString := conf.Now.Mongo.Db0
	db0 = Connect(connectionString)
	return db0
}
