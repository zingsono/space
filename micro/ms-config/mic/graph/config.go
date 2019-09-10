package graph

import (
	ql "github.com/graphql-go/graphql"

	"mic/model"
)

// Graphql Query
var ConfigQueryFields = ql.Fields{
	"config": &ql.Field{
		Type: ql.NewObject(ql.ObjectConfig{
			Name: "ConfigQueryType",
			Fields: ql.Fields{
				"total": &ql.Field{
					Type: ql.Int,
					Args: ql.FieldConfigArgument{
						"name": &ql.ArgumentConfig{Type: ql.String, Description: "服务名"},
					},
					Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
						return 0, nil
					},
					Description: "总记录数查询",
				},
				"list": &ql.Field{
					Type: ql.NewList(ql.NewObject(ql.ObjectConfig{
						Name: "ConfigListType",
						Fields: ql.Fields{
							"name": &ql.Field{
								Type:        ql.String,
								Description: "服务名",
							},
							"value": &ql.Field{
								Type:        ql.String,
								Description: "配置JSON字符串",
							},
						},
						Description: "集合对象",
					})),
					Args: ql.FieldConfigArgument{
						"limit": &ql.ArgumentConfig{
							Type:         ql.Int,
							DefaultValue: 100,
							Description:  "返回结果集数量",
						},
						"skip": &ql.ArgumentConfig{
							Type:         ql.Int,
							DefaultValue: 0,
							Description:  "跳过记录数",
						},
					},
					Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
						return i, nil
					},
					Description: "数据集合",
				},
				"info": &ql.Field{
					Type: ql.NewObject(ql.ObjectConfig{
						Name: "ConfigInfoType",
						Fields: ql.Fields{
							"name":      &ql.Field{Type: ql.String, Description: "服务名"},
							"value":     &ql.Field{Type: ql.String, Description: "配置JSON字符串"},
							"remark":    &ql.Field{Type: ql.String, Description: "配置备注描述信息"},
							"updatedAt": &ql.Field{Type: ql.String, Description: "最后更新时间"},
							"createdAt": &ql.Field{Type: ql.String, Description: "最后创建时间"},
						},
						Description: "配置信息",
					}),
					Args: ql.FieldConfigArgument{
						"name": &ql.ArgumentConfig{Type: ql.String, Description: "服务名"},
					},
					Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
						return i, nil
					},
					Description: "配置信息详情",
				},
			},
			Description: "配置信息查询类型",
		}),
		Description: "配置查询",
	},
}

// Graphql Mutation
var ConfigMutationFields = ql.Fields{
	"config": &ql.Field{
		Type: ql.NewObject(ql.ObjectConfig{
			Name: "ConfigMutationType",
			Fields: ql.Fields{
				"add": &ql.Field{
					Type: ql.Int,
					Args: ql.FieldConfigArgument{
						"name": &ql.ArgumentConfig{
							Type:         nil,
							DefaultValue: nil,
							Description:  "服务名",
						},
						"value": &ql.ArgumentConfig{
							Type:         nil,
							DefaultValue: "{}",
							Description:  "JSON格式配置信息",
						},
					},
					Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
						model.NewMsConfig("", "", "")
						model.MsConfig{}.Add("config", "{}", "备注")

						return 0, nil
					},
					Description: "新增配置",
				},
				"edit": &ql.Field{
					Name: "",
					Type: ql.Int,

					DeprecationReason: "",
					Description:       "",
				},
				"del": &ql.Field{
					Name: "",
					Type: ql.Int,
					Args: nil,

					DeprecationReason: "",
					Description:       "",
				},
				"updateStatus": &ql.Field{
					Name: "",
					Type: ql.Int,
					Args: nil,

					DeprecationReason: "",
					Description:       "",
				},
			},
			Description: "更新",
		}),
		Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
			return nil, nil
		},
		Description: "配置更新操作",
	},
}
