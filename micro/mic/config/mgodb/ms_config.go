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

type Kv struct {
	Key   string                 `json:"key" bson:"key"`
	Value map[string]interface{} `json:"value" bson:"value"`
}

// 配置信息管理
// 字段：服务名、配置JSON内容、备注、更新时间、创建时间
type Config struct {
	Key       string                 `json:"key" bson:"key"`
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

// 查询配置
func (c *Config) Query(keys []string, updateConfig func(kv map[string]interface{})) {
	collection := c.Collection()
	cursor, err := collection.Find(context.TODO(), bson.M{"key": bson.M{"$in": keys}}, options.Find().SetProjection(bson.M{"key": 1, "value": 1}))
	if err != nil {
		log.Panic(err)
	}
	for cursor.Next(context.TODO()) {
		if err := cursor.Err(); err != nil {
			log.Panic(err)
		}
		var row Config
		err = cursor.Decode(&row)
		if err != nil {
			log.Panic(err)
		}
		kv := make(map[string]interface{})
		kv[row.Key] = row.Value
		log.Print("################")
		updateConfig(kv)
	}
}

// 观察配置更新
func (c *Config) Watch(updateConfig func(kv map[string]interface{})) {
	collection := c.Collection()
	pipeline := mongo.Pipeline{{
		{"$match", bson.M{"operationType": bson.M{"$in": bson.A{"insert", "update"}}}},
	}}
	// Watch 新增与更新操作
	// update: {"_id": {"_data": "825DC9F595000000012B022C0100296E5A10043306C94DE71F44B684510C3E35B9F1F646645F696400645DC896DF0319D57180B38E570004"},"operationType": "update","clusterTime": {"$timestamp":{"t":"1573516693","i":"1"}},"ns": {"db": "msde","coll": "ms_config"},"documentKey": {"_id": {"$oid":"5dc896df0319d57180b38e57"}},"updateDescription": {"updatedFields": {"value": {"default": "mongo//:112"}},"removedFields": []},"fullDocument": {"_id": {"$oid":"5dc896df0319d57180b38e57"},"key": "mongo","value": {"default": "mongo//:112"}}}
	// insert: {"_id": {"_data": "825DC9F5D4000000032B022C0100296E5A10043306C94DE71F44B684510C3E35B9F1F646645F696400645DC9F5D588A45A8498BFF5410004"},"operationType": "insert","clusterTime": {"$timestamp":{"t":"1573516756","i":"3"}},"fullDocument": {"_id": {"$oid":"5dc9f5d588a45a8498bff541"},"key": "test","value": {"k1": "value1"}},"ns": {"db": "msde","coll": "ms_config"},"documentKey": {"_id": {"$oid":"5dc9f5d588a45a8498bff541"}}}
	// delete: {"_id": {"_data": "825DC9F5F8000000012B022C0100296E5A10043306C94DE71F44B684510C3E35B9F1F646645F696400645DC9F5D588A45A8498BFF5410004"},"operationType": "delete","clusterTime": {"$timestamp":{"t":"1573516792","i":"1"}},"ns": {"db": "msde","coll": "ms_config"},"documentKey": {"_id": {"$oid":"5dc9f5d588a45a8498bff541"}}}

	changeStreamOptions := options.ChangeStream().SetBatchSize(1).SetFullDocument(options.UpdateLookup)
	changeStream, err := collection.Watch(context.TODO(), pipeline, changeStreamOptions)
	if err != nil {
		log.Fatal(err)
	}
	for changeStream.Next(context.TODO()) {
		log.Print(changeStream.Current.String())
		var row Config
		changeStream.Current.Lookup("fullDocument").Unmarshal(&row)
		kv := make(map[string]interface{})
		kv[row.Key] = row.Value
		log.Print("------------------------------")
		updateConfig(kv)
	}
}
