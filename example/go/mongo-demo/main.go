package main

import (
	"context"
	"log"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	log.Println("BEG------------------------------")

	Db0()

	log.Println("END------------------------------")
}

func Connect(connectionString string) *mongo.Database {
	dbName := (strings.Split((strings.Split(connectionString, "/"))[3], "?"))[0]
	if dbName == "" {
		log.Fatalf("Errror Mongodb connectionString %s", connectionString)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// TODO 设置 SetMinPoolSize(3) 会无法连接
	options := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		log.Fatalf("Error Connect Mongodb Fail...")
	}
	database := client.Database(dbName)
	log.Printf("Mongodb connect %s %s", "db0", connectionString)
	return database
}

// 数据库连接全局缓存
var db0 *mongo.Database

func init() {
	connectionString := "mongodb://test:test@121.40.83.200:37017/test?authSource=admin&authMechanism=SCRAM-SHA-1&maxPoolSize=100"
	// connectionString := conf.Now.Mongo.Db0
	db0 = Connect(connectionString)
	c, e := db0.ListCollectionNames(context.TODO(), bson.D{})
	if e != nil {
		log.Fatal(e)
	}
	log.Print(c)
}

// 获取数据库连接
func Db0() *mongo.Database {
	if db0 != nil {
		return db0
	}
	connectionString := "mongodb://test:test@121.40.83.200:37017/test?authSource=admin&authMechanism=SCRAM-SHA-1&maxPoolSize=100"
	// connectionString := conf.Now.Mongo.Db0
	db0 = Connect(connectionString)
	return db0
}
