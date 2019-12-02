package graph

import (
	"github.com/graphql-go/graphql"
)

func init() {
	QueryFields(_query)
}

var _query = graphql.Fields{}
