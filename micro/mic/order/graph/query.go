package graph

import (
	"log"

	"github.com/graphql-go/graphql"
)

func init() {
	QueryFields(_query)
}

var _query = graphql.Fields{
	"order": &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name:        "MemberQueryType",
			Interfaces:  nil,
			Fields:      graphql.Fields{},
			IsTypeOf:    nil,
			Description: "查询类型",
		}),
		Args: nil,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			log.Print("MemberType............")
			return "", e
		},
		DeprecationReason: "",
		Description:       "查询服务",
	},
}
