// WebSocket Message Server
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/websocket"
)

// 全局变量定义
var (
	config *Config
)

// ** 启动服务
func main() {
	fmt.Println("Start server......")
	// 初始化配置
	config = NowConfig()

	// 消息发送
	http.HandleFunc("/put", func(writer http.ResponseWriter, request *http.Request) {
		// NewMessage().Put()
		writer.Write([]byte("msg put"))
	})

	// 消息订阅
	http.Handle("/poll?cid=32uuid", websocket.Handler(func(ws *websocket.Conn) {
		log.Printf("RequestURI %s", ws.Request().RequestURI)
		cid := ws.Request().RequestURI
		NewMessage().Poll(cid, ws)
	}))

	// 启动服务
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", config.Port), nil))
}

type Config struct {
	Port                  int
	MongoConnectionString string
}

// 设置配置对象
func NowConfig() *Config {
	return &Config{
		Port:                  8610,
		MongoConnectionString: "mongodb://test:test@121.40.83.200:37017/test?authSource=admin&authMechanism=SCRAM-SHA-1",
	}
}

// Mongodb 消息集合：ws
type Message struct {
	// 客户端编号，根据此字段订阅消息
	Cid string `bson:"cid" json:"cid"`
	// 消息内容文本
	Msg string `bson:"msg" json:"msg"`
	// 来源用户cid
	Origin string `bson:"origin" json:"origin"`
	// 消息状态（1=已接收 0=新消息）
	Status int8 `bson:"status" json:"status"`
	// 创建时间
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	// 过期时间，过期清理
	Expires time.Time `bson:"expires" json:"expires"`
}

func NewMessage() *Message {
	return new(Message)
}

func (m *Message) Collection() *mongo.Collection {
	return Db().Collection("ws")
}

// 发送消息
func (m *Message) Put(origin, targetCid, msg string) {
	m.Cid = targetCid
	m.Origin = origin
	m.Msg = msg
	m.Status = 0
	m.CreatedAt = time.Now()
	m.Expires = time.Now()
	m.Collection().InsertOne(context.TODO(), m)
}

// 订阅消息
func (m *Message) Poll(cid string, ws *websocket.Conn) *Message {
	ctx := context.TODO()
	changeStream, err := m.Collection().Watch(ctx, &bson.M{"match": &bson.M{"cid": cid}})
	if err != nil {
		log.Print(err)
	}
	for {
		if changeStream.Next(ctx) {
			bson.Unmarshal(changeStream.Current, &m)
			// 发送Ws消息
			err := websocket.Message.Send(ws, m.Msg)
			if err != nil {
				log.Print(err)
			}
		}
	}
}

func Connect(connectionString string) *mongo.Database {
	dbName := (strings.Split((strings.Split(connectionString, "/"))[3], "?"))[0]
	if dbName == "" {
		log.Fatalf("Errror Mongodb connectionString %s", connectionString)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	options := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		log.Fatalf("Error Connect Mongodb Fail...")
	}
	database := client.Database(dbName)
	log.Printf("Mongodb connect %s", connectionString)
	return database
}

// 数据库连接全局缓存
var db *mongo.Database

// 获取数据库连接
func Db() *mongo.Database {
	if db != nil {
		return db
	}
	// mongodb://test:test@121.40.83.200:37017/test?authSource=admin&authMechanism=SCRAM-SHA-1
	connectionString := config.MongoConnectionString
	db = Connect(connectionString)
	return db
}
