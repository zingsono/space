package model

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func FieldsMerge(args ...graphql.Fields) graphql.Fields {
	fields := make(graphql.Fields)
	for _, arg := range args {
		for k, v := range arg {
			fields[k] = v
		}
	}
	return fields
}

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name:       "Query",
	Interfaces: nil,
	Fields: FieldsMerge(
		// 查询字段
		UserQueryFields,
	),
	IsTypeOf:    nil,
	Description: "查询",
})

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name:       "Mutation",
	Interfaces: nil,
	Fields: FieldsMerge(
		// 更新字段
		UserMutationFields,
	),
	IsTypeOf:    nil,
	Description: "更新",
})

// Graphql Schema
var GraphqlSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:        QueryType,
	Mutation:     MutationType,
	Subscription: nil,
	Types:        nil,
	Directives:   nil,
	Extensions:   nil,
})

// Http Handler
var GraphqlHttpHandler = handler.New(&handler.Config{
	Schema:   &GraphqlSchema,
	Pretty:   true,
	GraphiQL: true,
})
