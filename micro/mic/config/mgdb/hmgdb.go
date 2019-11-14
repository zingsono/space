package mgdb

import (
	"github.com/zingsono/space/micro/lib/hmgdb"
	"go.mongodb.org/mongo-driver/mongo"
)

func Collection(collectionName string) *mongo.Collection {
	return hmgdb.Db().Collection(collectionName)
}

func SetConnectString(connectString string) {
	hmgdb.set
}
