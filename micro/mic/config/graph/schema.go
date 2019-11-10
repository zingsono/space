package graph

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Graphql Schema

// 查询字段定义
var QueryFields = FieldsMerge(
	ConfigQueryFields,
)

// 更新字段定义
var MutationFields = FieldsMerge(
	ConfigMutationFields,
)

// 参考Github更新操作文档定义 https://developer.github.com/v4/mutation/

// Http Handler
var GraphqlHttpHandler = func() *handler.Handler {
	return handler.New(&handler.Config{
		Schema:   GraphqlSchema(),
		Pretty:   true,
		GraphiQL: true,
	})
}

// Graphql Schema
var GraphqlSchema = func() *graphql.Schema {
	newSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:        QueryType,
		Mutation:     MutationType,
		Subscription: nil,
		Types:        nil,
		Directives:   nil,
		Extensions:   nil,
	})
	if err != nil {
		// 异常退出
		log.Fatal(err)
	}
	log.Printf("GraphqlSchema Load Success")
	return &newSchema
}

// Graphql Query Type
var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Query",
	Interfaces:  nil,
	Fields:      QueryFields,
	IsTypeOf:    nil,
	Description: "查询操作",
})

// Graphql Mutation Type
var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Mutation",
	Interfaces:  nil,
	Fields:      MutationFields,
	IsTypeOf:    nil,
	Description: "更新操作",
})

// Merge Graphql Fields
func FieldsMerge(args ...graphql.Fields) graphql.Fields {
	fields := make(graphql.Fields)
	for _, arg := range args {
		for k, v := range arg {
			fields[k] = v
		}
	}
	return fields
}
