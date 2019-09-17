package graph

import (
	"github.com/graphql-go/graphql"
)

// Mongodb 更新结果,对应结构体：mongo.UpdateResult
var UpdateResultType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UpdateResultType",
	Fields: graphql.Fields{
		"matchedCount":  &graphql.Field{Type: graphql.Int, Description: "更新条件匹配文档数量"},
		"modifiedCount": &graphql.Field{Type: graphql.Int, Description: "更新文档数量"},
		"upsertedCount": &graphql.Field{Type: graphql.Int, Description: "upsert时插入文档数量"},
		"upsertedID":    &graphql.Field{Type: graphql.String, Description: "发生upsert时插入文档的标识符"},
	},
	Description: "Mongodb更新结果",
})
