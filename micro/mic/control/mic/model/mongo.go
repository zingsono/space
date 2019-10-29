package model

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func MongoClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("连接Mongodb失败...")
	}
	return client
}
