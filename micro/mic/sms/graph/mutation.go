package graph

import (
	g "github.com/graphql-go/graphql"
)

func init() {
	MutationFields(_mutation)
}

var _mutation = g.Fields{
	"sms": &g.Field{
		Description: "短信服务",
		Type: g.NewObject(g.ObjectConfig{
			Name:       "SmsMutationType",
			Interfaces: nil,
			Fields: g.Fields{
				"send": &g.Field{
					Type: ResponseCodeType,
					Args: g.FieldConfigArgument{
						"mobile":     &g.ArgumentConfig{Type: g.String, DefaultValue: nil, Description: "手机号码"},
						"templateId": &g.ArgumentConfig{Type: g.String, DefaultValue: nil, Description: "模板编号"},
					},
					Resolve: func(p g.ResolveParams) (i interface{}, err error) {
						return ERR_SUCCESS, err
					},
					Description: "发送短信，仅限平台系统使用此接口，对外提供短信服务的接口单独提供。",
				},
			},
			IsTypeOf:    nil,
			Description: "短信操作类型",
		}),
		Args: nil,
		Resolve: func(p g.ResolveParams) (i interface{}, e error) {
			return "", e
		},
	},
}
