package model

import (
	"github.com/graphql-go/graphql"
	"mic/commons"
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name:       "Query",
	Interfaces: nil,
	Fields: commons.MapMerge(
		UserQueryFields,
	),
	IsTypeOf:    nil,
	Description: "查询",
})

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name:       "Mutation",
	Interfaces: nil,
	Fields: commons.MapMerge(
		UserMutationFields,
	),
	IsTypeOf:    nil,
	Description: "更新",
})

var GraphqlSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:        QueryType,
	Mutation:     MutationType,
	Subscription: nil,
	Types:        nil,
	Directives:   nil,
	Extensions:   nil,
})
