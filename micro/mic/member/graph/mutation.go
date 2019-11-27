package graph

import (
	"log"

	g "github.com/graphql-go/graphql"
)

func init() {
	MutationFields(_mutation)
}

var _mutation = g.Fields{
	"smsRegister": &g.Field{
		Description: "短信验证注册会员",
		Type:        g.String,
		Args: g.FieldConfigArgument{
			"reamrk": &g.ArgumentConfig{
				Type:         g.String,
				DefaultValue: nil,
				Description:  "备注信息",
			},
		},
		Resolve: func(p g.ResolveParams) (i interface{}, e error) {
			log.Print("")
			return i, e
		},
	},
	"login": &g.Field{
		Description: "账号密码登录",
		Type:        g.String,
		Args: g.FieldConfigArgument{
			"reamrk": &g.ArgumentConfig{
				Type:         g.String,
				DefaultValue: nil,
				Description:  "备注信息",
			},
		},
		Resolve: func(p g.ResolveParams) (i interface{}, e error) {
			log.Print("")
			return i, e
		},
	},
	"smsResetPassword": &g.Field{
		Description: "短信验证重置密码",
		Type:        g.String,
		Args: g.FieldConfigArgument{
			"reamrk": &g.ArgumentConfig{
				Type:         g.String,
				DefaultValue: nil,
				Description:  "备注信息",
			},
		},
		Resolve: func(p g.ResolveParams) (i interface{}, e error) {
			log.Print("")
			return i, e
		},
	},
}
