// mongodb driver doc https://godoc.org/go.mongodb.org/mongo-driver/mongo
package main

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 异常使用 log.Panicf() ，输出日志并抛出
// 外层函数使用recover()  处理
func Throw(err error, args ...interface{}) {
	if err != nil {
		log.Panicf(err.Error(), args)
	}
}

// 连接字符串配置
var dbs = make(map[string]string)

func Initialize(dbMap map[string]string)  {
	dbs = dbMap
}

func Db(db string) *mongo.Database {
	connstr,ok = dbs[db]
	if !ok {
		Throw(errors.New("MongoDB数据库(s%)配置信息不存在"),db)
	}
	var opt = options.Client().ApplyURI(connstr)
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	var client, err = mongo.Connect(ctx, opt)
	Throw(err)
	return client.Database(db)
}
