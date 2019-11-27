package graph

import (
	g "github.com/graphql-go/graphql"
)

func init() {
	MutationFields(_mutation)
}

var _mutation = g.Fields{}
