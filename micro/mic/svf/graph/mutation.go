package graph

import (
	"github.com/graphql-go/graphql"
)

func init() {
	MutationFields(_mutation)
}

var _mutation = graphql.Fields{
	"svf": &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name:       "SvfMutationType",
			Interfaces: nil,
			Fields: graphql.Fields{
				"recharge": &graphql.Field{
					Description: "储值账户充值",
					Type:        ResponseCodeType,
					Args: graphql.FieldConfigArgument{
						"cid":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String), DefaultValue: nil, Description: "客户端充值交易流水号，同一流水号重复请求不会重复充值"},
						"uid":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String), DefaultValue: nil, Description: "会员用户ID"},
						"amount": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int), DefaultValue: nil, Description: "交易金额"},
						"reamrk": &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: nil, Description: "备注信息"},
					},
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
						return ERR_SUCCESS, e
					},
				},

				"expenditure": &graphql.Field{
					Description: "储值账户消费",
					Type:        ResponseCodeType,
					Args: graphql.FieldConfigArgument{
						"cid":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String), DefaultValue: nil, Description: "客户端充值交易流水号，同一流水号重复请求不会重复充值"},
						"uid":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String), DefaultValue: nil, Description: "会员用户ID"},
						"amount": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int), DefaultValue: nil, Description: "交易金额"},
						"reamrk": &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: nil, Description: "备注信息"},
					},
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
						return ERR_SUCCESS, e
					},
				},
			},
			IsTypeOf:    nil,
			Description: "储值账户更新类型",
		}),
		Resolve:     nil,
		Description: "储值账户更新操作",
	},
}
