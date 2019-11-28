package graph

import (
	"log"

	"github.com/graphql-go/graphql"
)

func init() {
	MutationFields(_mutation)
}

var _mutation = graphql.Fields{
	"order": &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name:       "OrderMutationType",
			Interfaces: nil,
			Fields: graphql.Fields{
				"create": &graphql.Field{
					Type: graphql.String,
					Args: nil,
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
						return "ttt", e
					},
					Description: "创建订单",
				},
			},
			IsTypeOf:    nil,
			Description: "订单操作类型",
		}),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			log.Print("OrderMutationType............")
			return "", e
		},
		Description: "订单服务操作",
	},
}
