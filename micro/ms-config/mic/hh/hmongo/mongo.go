// Mongodb Helper
package hmongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongodb客户端集合，Key为连接名（用于多数据源的情况），value为Client
var dbs map[string]*mongo.Database

// 初始化客户端
// key为连接名，value为连接字符串
// 项目启动前调用此方法
func Initialize(connectionMap map[string]string) {
	dbs = make(map[string]*mongo.Database)
	for k, v := range connectionMap {
		var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
		var options = options.Client().ApplyURI(v)
		client, err := mongo.Connect(ctx, options)
		if err != nil {
			log.Fatalf("Error Connect Mongodb Fail...")
		}
		dbs[k] = client.Database(k)
		log.Printf("Mongo load %s => %s", k, v)
	}
	log.Println("初始化Mongodb客户端完成...")
}

// 根据连接名读取缓存数据库客户端
func Db(datasourceName string) *mongo.Database {
	return dbs[datasourceName]
}

// 查询结果集合处理
