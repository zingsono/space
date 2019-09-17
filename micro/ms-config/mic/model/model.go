package model

import (
	"go.mongodb.org/mongo-driver/mongo"

	"mic/helper/hmongo"
)

// 数据库连接
// 多数据库连接时，新增一个获取函数
var DB = func() *mongo.Database {
	return hmongo.Db("test")
}
