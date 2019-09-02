package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	log.Println("BEG------------------------------")

	var opt = options.Client().ApplyURI("mongodb://test:test@121.40.83.200:37017/test?authSource=admin&authMechanism=SCRAM-SHA-1")
	var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	var client, err = mongo.Connect(ctx, opt)

	log.Println("Ping ")
	// Check the connection
	err = client.Ping(ctx, nil)

	names, nerr := client.ListDatabaseNames(ctx, bson.D{})
	if nerr != nil {
		log.Fatal(nerr)
	}
	log.Println(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	log.Println("END------------------------------")
}
