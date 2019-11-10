package mgodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 配置信息管理
// 字段：服务名、配置JSON内容、备注、更新时间、创建时间
type Config struct {
	Key       string                 `json:"name" bson:"key"`
	Value     map[string]interface{} `json:"value" bson:"value"`
	Remark    string                 `json:"remark" bson:"remark"`
	UpdatedAt time.Time              `json:"updatedAt" bson:"updatedAt"`
	CreatedAt time.Time              `json:"createdAt" bson:"createdAt"`
}

var NewConfig = func() *Config {
	return new(Config)
}

func (*Config) Collection() *mongo.Collection {
	return Mgo.GetDatabase(DEFAULT).Collection("ms_config")
}

// 保存数据，不存在新增，存在则更新
func (c *Config) Save() (*mongo.UpdateResult, error) {
	c.UpdatedAt = time.Now().Local()
	log.Println(c)
	updateResult, err := c.Collection().UpdateOne(context.TODO(), bson.M{"key": c.Key}, bson.M{"$set": c}, options.Update().SetUpsert(true))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// ObjectID 改为string显示
	if updateResult.UpsertedID != nil {
		updateResult.UpsertedID = updateResult.UpsertedID.(primitive.ObjectID).Hex()
	}
	return updateResult, err
}

func (c *Config) FindOne(key string) (*Config, error) {
	log.Printf("执行FindOne")
	singleResult := c.Collection().FindOne(context.Background(), bson.M{"key": key})
	msConfig := &Config{}
	err := singleResult.Decode(msConfig)
	if err != nil {
		return nil, err
	}
	return msConfig, nil
}

// 观察配置更新
func (c *Config) Watch(keys []string, updateConfig func(value map[string]interface{})) {
	collection := c.Collection()
	pipeline := mongo.Pipeline{{
		{"$match", bson.M{"operationType": bson.M{"$in": bson.A{"insert", "updated"}}}},
	}}
	changeStreamOptions := options.ChangeStream().SetBatchSize(1).SetFullDocument(options.UpdateLookup)
	changeStream, err := collection.Watch(context.TODO(), pipeline, changeStreamOptions)
	if err != nil {
		log.Fatal(err)
	}
	for changeStream.Next(context.TODO()) {
		var config *Config
		changeStream.Current.Lookup("fullDocument.value").Unmarshal(config)
		updateConfig(config.Value)
	}
}
