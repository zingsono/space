package mgodb

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DEFAULT = "default"
)

var Mgo = new(mgo)

type mgo struct {
	dbs sync.Map
}

func (m *mgo) GetDatabase(name string) *mongo.Database {
	db, ok := m.dbs.Load(name)
	if ok {
		return nil
	}
	return db.(*mongo.Database)
}

func (m *mgo) SetDatabase(name string, v *mongo.Database) {
	if name == "" {
		name = DEFAULT
	}
	m.dbs.Store(name, v)
}

func SetDatabase(name string, connectionString string) {
	Mgo.SetDatabase(name, connect(connectionString))
}

// 注意处理连接panic
func connect(connectionString string) *mongo.Database {
	dbName := (strings.Split((strings.Split(connectionString, "/"))[3], "?"))[0]
	if dbName == "" {
		panic(fmt.Sprintf("Errror Mongodb connectionString %s", connectionString))
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	options := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		panic(fmt.Sprintf("Errror Connect mongodb exception %s", err))
	}
	database := client.Database(dbName)
	names, err := client.ListDatabaseNames(ctx, bson.D{})
	if err != nil {
		panic(fmt.Sprintf("ListDatabaseNames Connect exception: %s", err))
	}
	log.Printf("Mongodb connect success -> %s  DatabaseNames->%s", connectionString, names)
	return database
}
