package mgdb

import (
	"github.com/zingsono/space/micro/lib/hmgdb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Collection(collectionName string, opts ...*options.CollectionOptions) *mongo.Collection {
	return mgdb.Db().Collection(collectionName, opts...)
}

func SetConnectString(connectString string) {
	mgdb.SetConnectString(mgdb.DEFAULT, connectString)
}
