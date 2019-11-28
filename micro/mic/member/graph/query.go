package graph

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func init() {
	QueryFields(_query)
}

type Mem struct {
	Id string `json:"memberId"`
}

var _query = graphql.Fields{
	"member": &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name:       "MemberQueryType",
			Interfaces: nil,
			Fields: graphql.Fields{
				"single": &graphql.Field{
					Type: MemberInfoType,
					Args: graphql.FieldConfigArgument{
						"uid": &graphql.ArgumentConfig{
							Type:         graphql.String,
							DefaultValue: nil,
							Description:  "会员用户ID",
						},
						"token": &graphql.ArgumentConfig{
							Type:         graphql.String,
							DefaultValue: nil,
							Description:  "会话授权Token",
						},
					},
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
						log.Print("MemberInfoType single............")
						log.Println(p.Args)
						return &Mem{Id: "2342341"}, e
					},
					Description: "查询单个会员用户信息",
				},
				"total": &graphql.Field{
					Type:        graphql.Int,
					Description: "会员总数，最大值5000，当结果集大于5000时，如果有必要，更换搜索条件，请勿依赖此字段做统计。",
					/*Args: graphql.FieldConfigArgument{

					},*/
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
						return 5000, e
					},
				},
				"list": &graphql.Field{
					Type: graphql.NewList(MemberInfoType),
					/*Args: graphql.FieldConfigArgument{

					},*/
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
						log.Print("MemberInfoType dataset............")
						log.Println(p.Args)
						var ms = []*Mem{{Id: "234234111111111111111"}}
						v, e := json.Marshal(ms)
						fmt.Println(v)
						return ms, e
					},
					Description: "会员分页查询",
				},
			},
			IsTypeOf:    nil,
			Description: "会员查询类型",
		}),
		Args: nil,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			log.Print("MemberType............")
			return "", e
		},
		DeprecationReason: "",
		Description:       "会员查询服务",
	},
}

var MemberInfoType = graphql.NewObject(graphql.ObjectConfig{
	Name:       "MemberInfoType",
	Interfaces: nil,
	Fields: graphql.Fields{
		"uid":           &graphql.Field{Type: graphql.String, Description: "会员用户ID"},
		"mobile":        &graphql.Field{Type: graphql.String, Description: "手机号"},
		"email":         &graphql.Field{Type: graphql.String, Description: "邮箱"},
		"lastLoginTime": &graphql.Field{Type: graphql.String, Description: "最后登录时间"},
		"createdAt":     &graphql.Field{Type: graphql.String, Description: "会员注册时间"},
		"memberType":    &graphql.Field{Type: graphql.String, Description: "会员类型（1=普通用户 2=机构用户）"},
		"points": &graphql.Field{
			Type: graphql.NewObject(graphql.ObjectConfig{
				Name:       "MemberPointsType",
				Interfaces: nil,
				Fields: graphql.Fields{
					"accountNo": &graphql.Field{Type: graphql.String, Description: "积分账户编号"},
					"balance":   &graphql.Field{Type: graphql.Int, Description: "积分账户余额"},
				},
				IsTypeOf:    nil,
				Description: "会员积分账户类型",
			}),
			Args: graphql.FieldConfigArgument{
				"issueUid": &graphql.ArgumentConfig{
					Type:         graphql.NewNonNull(graphql.String),
					DefaultValue: nil,
					Description:  "发行会员编号",
				},
			},
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				// 调用积分账户查询用户积分
				return "", e
			},
			Description: "会员积分账户，没有授权积分功能的服务此字段不可用",
		},
	},
	IsTypeOf:    nil,
	Description: "会员用户信息",
})
