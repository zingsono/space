package graph

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Graphql Schema Fields
var (
	queryFields    graphql.Fields
	mutationFields graphql.Fields
)

// 查询字段定义
func MergeQueryFields(args ...graphql.Fields) {
	if queryFields == nil {
		queryFields = make(graphql.Fields)
	}
	for _, arg := range args {
		for k, v := range arg {
			queryFields[k] = v
		}
	}
}

// 更新字段定义
func MergeMutationFields(args ...graphql.Fields) {
	if mutationFields == nil {
		mutationFields = make(graphql.Fields)
	}
	for _, arg := range args {
		for k, v := range arg {
			mutationFields[k] = v
		}
	}
}

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
		Query:        QueryType(),
		Mutation:     MutationType(),
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
var QueryType = func() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "Query",
		Interfaces:  nil,
		Fields:      queryFields,
		IsTypeOf:    nil,
		Description: "查询操作",
	})
}

// Graphql Mutation Type
var MutationType = func() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "Mutation",
		Interfaces:  nil,
		Fields:      mutationFields,
		IsTypeOf:    nil,
		Description: "更新操作",
	})
}
