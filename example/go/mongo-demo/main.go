package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	options := options.Client().ApplyURI("mongodb://unionlive:unionlive@211.152.57.29:39017/unionlive?authSource=admin&authMechanism=SCRAM-SHA-1")
	client, err := mongo.Connect(context.Background(), options)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

}
