package libmongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongodb连接字符串，应用通过配置读取并赋值
var ConnectionString string

func Db(db string) *mongo.Database {
	var opt = options.Client().ApplyURI(ConnectionString)
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	var client, err = mongo.Connect(ctx, opt)
	Throw(err)
	return client.Database(db)
}
