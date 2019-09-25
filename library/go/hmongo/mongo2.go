// Mongodb Helper
package hmongo

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var databases = make(map[string]*mongo.Database)
var mongoChan = make(chan *MongoConfig)

// MongoDB连接名与连接字符串
type MongoConfig struct {
	Name string
	Uri  string
}

func init() {
	go func() {
		// 循环接收通道数据，实现动态更新配置
		for mConf := range mongoChan {
			db, err := newDatabase(mConf.Uri)
			if err != nil {
				log.Printf("Error Connect Mongodb Fail name=‘%s’ Uri=%s Msg: %s", mConf.Name, mConf.Uri, err.Error())
				continue
			}
			SetDb(mConf.Name, db)
			log.Printf("Connect Mongodb success name=‘%s’ Uri=%s", mConf.Name, mConf.Uri)
		}
	}()
}

func newDatabase(uri string) (*mongo.Database, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var options = options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return nil, err
	}
	return client.Database("test"), nil
}

// 设置Mongodb连接配置,配置更新是，重新设置，可动态更新配置
func SetConnectUri(name string, uri string) {
	conf := &MongoConfig{
		Name: name,
		Uri:  uri,
	}
	mongoChan <- conf
}

func SetDb(name string, database *mongo.Database) {
	databases[name] = database
}

// 读取数据库连接
func GetDb(name string) (*mongo.Database, error) {
	db := databases[name]
	if db == nil {
		return nil, errors.New("未读取到数据库连接")
	}
	return db, nil
}
