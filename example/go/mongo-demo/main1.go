package main

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Test struct {
	Name string
	Age  int64
}

func main() {
	log.Println("BEG------------------------------")

	var opt = options.Client().ApplyURI("mongodb://test:test@121.40.83.200:37017/test?authSource=admin&authMechanism=SCRAM-SHA-1")
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	var client, err = mongo.Connect(ctx, opt)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("client...")
	database := client.Database("test")
	collection := database.Collection("test")
	cname := collection.Name()
	log.Println("cname=" + cname)

	oneRs, err := collection.InsertOne(context.Background(), &Test{
		Name: "宋江",
		Age:  102,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(primitive.NewObjectID().Hex())
	log.Println(oneRs.InsertedID)
	log.Println(oneRs.InsertedID.(primitive.ObjectID).Hex())
	log.Println(reflect.TypeOf(oneRs.InsertedID))

	fmt.Println("Connected to MongoDB!")
	log.Println("END------------------------------")
}
