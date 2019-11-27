package graph

import (
	"log"

	"github.com/graphql-go/graphql"
)

func init() {
	QueryFields(_MallQueryFields)
}

var _MallQueryFields = graphql.Fields{
	"mall": &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name:       "MallQueryType",
			Interfaces: nil,
			Fields: graphql.Fields{
				"single": &graphql.Field{
					Type: MallInfoType,
					Args: Argument(MaillIdArg),
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
						log.Print("MemberInfoType single............")
						log.Println(p.Args)

						return &Mall{MallId: "123123", Name: "名称", Title: "天猫商城"}, e
					},
					Description: "商城信息查询",
				},
				"total": &graphql.Field{
					Type:        graphql.Int,
					Description: "会员总数，最大值5000，不能依赖此字段做统计。",
					Args:        Argument(MaillDatasetArg),
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
						return 5000, e
					},
				},
				"dataset": &graphql.Field{
					Type: graphql.NewList(MallInfoType),
					Args: Argument(PageArgument, MaillDatasetArg),
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
						log.Print("MemberInfoType dataset............")
						log.Println(p.Args)
						return i, e
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
		Description:       "购物商城服务",
	},
}

// 单条记录查询参数定义
var MaillIdArg = graphql.FieldConfigArgument{
	"mallId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String), DefaultValue: nil, Description: "商城编号"},
}

// 结果集查询参数定义
var MaillDatasetArg = graphql.FieldConfigArgument{
	"mallId": &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: nil, Description: "商城编号"},
	"name":   &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: nil, Description: "商城名称"},
}

// 商城信息类型定义
var MallInfoType = graphql.NewObject(graphql.ObjectConfig{
	Name:       "MallInfoType",
	Interfaces: nil,
	Fields: graphql.Fields{
		"mallId":      &graphql.Field{Type: graphql.String, Description: "商城编号"},
		"name":        &graphql.Field{Type: graphql.String, Description: "商城名"},
		"title":       &graphql.Field{Type: graphql.String, Description: "商城页面标题"},
		"description": &graphql.Field{Type: graphql.String, Description: "商城页面描述"},
		"site":        &graphql.Field{Type: graphql.String, Description: "访问商城网站链接地址"},
		"remark":      &graphql.Field{Type: graphql.String, Description: "备注信息，运营人员添加的特殊备注，用于管理系统展示"},
		"createdAt":   &graphql.Field{Type: graphql.String, Description: "创建时间"},
		"updatedAt":   &graphql.Field{Type: graphql.String, Description: "更新时间"},
		"ad": &graphql.Field{
			Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
				Name:       "MallAdType",
				Interfaces: nil,
				Fields: graphql.Fields{
					"id":     &graphql.Field{Type: graphql.String, Description: "广告当前组唯一编号"},
					"key":    &graphql.Field{Type: graphql.String, Description: "广告位标识字符串"},
					"imgUrl": &graphql.Field{Type: graphql.String, Description: "图片url"},
					"href":   &graphql.Field{Type: graphql.String, Description: "跳转链接"},
					"remark": &graphql.Field{Type: graphql.String, Description: "备注信息"},
					"rank":   &graphql.Field{Type: graphql.Int, Description: "排序"},
					"show":   &graphql.Field{Type: graphql.String, Description: "是否展示（1=展示 0=隐藏）"},
				},
				IsTypeOf:    nil,
				Description: "广告类型",
			})),
			Args: graphql.FieldConfigArgument{
				"key":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String), DefaultValue: nil, Description: "广告位标识字符串，如：banner"},
				"mallId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String), DefaultValue: nil, Description: "商城编号"},
				"show":   &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: "1", Description: "是否展示（1=展示 0=隐藏），参数为空时返回全部，默认只返回展示状态的广告"},
			},
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {

				return i, e
			},
			Description: "商城广告信息,返回广告数组",
		},
	},
	IsTypeOf:    nil,
	Description: "商城类型",
})
