package graph

import (
	"log"

	g "github.com/graphql-go/graphql"
)

func init() {
	MutationFields(_mutation)
}

var _mutation = g.Fields{
	"member": &g.Field{
		Name: "",
		Type: g.NewObject(g.ObjectConfig{
			Name:       "MemberMutationType",
			Interfaces: nil,
			Fields: g.Fields{
				"smsRegister": &g.Field{
					Description: "短信验证注册会员",
					Type:        g.String,
					Args: g.FieldConfigArgument{
						"orgUid": &g.ArgumentConfig{Type: g.String, Description: "机构用户编号"},
						"mobile": &g.ArgumentConfig{Type: g.String, Description: "手机号码"},
						"code":   &g.ArgumentConfig{Type: g.String, Description: "短信验证码"},
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
						"orgUid":   &g.ArgumentConfig{Type: g.String, Description: "机构用户编号"},
						"loginId":  &g.ArgumentConfig{Type: g.String, DefaultValue: nil, Description: "登录账号/手机/邮箱"},
						"password": &g.ArgumentConfig{Type: g.String, DefaultValue: nil, Description: "登录密码"},
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
						"orgUid":      &g.ArgumentConfig{Type: g.String, Description: "机构用户编号"},
						"mobile":      &g.ArgumentConfig{Type: g.String, Description: "手机号码"},
						"newPassword": &g.ArgumentConfig{Type: g.String, Description: "新密码"},
						"code":        &g.ArgumentConfig{Type: g.String, Description: "短信验证码"},
					},
					Resolve: func(p g.ResolveParams) (i interface{}, e error) {
						log.Print("")
						return i, e
					},
				},
			},
			IsTypeOf:    nil,
			Description: "会员服务操作类型",
		}),
		Args:              nil,
		Resolve:           nil,
		DeprecationReason: "",
		Description:       "",
	},
}
