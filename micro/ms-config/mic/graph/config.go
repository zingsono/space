package graph

import (
	"encoding/json"
	"log"

	ql "github.com/graphql-go/graphql"

	"mic/model"
)

type Config struct {
	Total int64
	List  []model.MsConfig
	Info  model.MsConfig
	Value string
}

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
						},
						Description: "配置信息",
					}),
					Args: ql.FieldConfigArgument{
						"name": &ql.ArgumentConfig{Type: ql.NewNonNull(ql.String), Description: "服务名"},
					},
					Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
						log.Printf("执行 query config info")
						name := p.Args["name"].(string)
						msConfig, e := (&model.MsConfig{}).FindOne(name)
						return msConfig, e
					},
					Description: "配置信息详情",
				},
				"value": &ql.Field{
					Name: "",
					Type: ql.String,
					Args: ql.FieldConfigArgument{
						"name": &ql.ArgumentConfig{Type: ql.NewNonNull(ql.String), Description: "服务名"},
					},
					Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
						log.Printf("执行 query config value")
						name := p.Args["name"].(string)
						msConfig, e := (&model.MsConfig{}).FindOne(name)
						return msConfig.Value, e
					},
					DeprecationReason: "",
					Description:       "配置内容JSON字符串",
				},
			},
			Description: "配置信息查询类型",
		}),
		Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
			log.Printf("执行 query config")
			rs := &Config{}
			return rs, e
		},
		Description: "配置查询",
	},
}

// Graphql Mutation
var ConfigMutationFields = ql.Fields{
	"configSave": &ql.Field{
		Type: UpdateResultType,
		Args: ql.FieldConfigArgument{
			"name": &ql.ArgumentConfig{
				Type:         ql.NewNonNull(ql.String),
				DefaultValue: nil,
				Description:  "服务名",
			},
			"value": &ql.ArgumentConfig{
				Type:         ql.NewNonNull(ql.String),
				DefaultValue: "{}",
				Description:  "JSON格式配置信息",
			},
		},
		Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
			log.Printf("执行QL方法：%s", "config.configSave")
			value := p.Args["value"].(string)
			jsonMap := make(map[string]interface{})
			e = json.Unmarshal([]byte(value), &jsonMap)
			if e != nil {
				return nil, e
			}
			updateResult, e := (&(model.MsConfig{Name: p.Args["name"].(string), Value: jsonMap})).Save()
			if e != nil {
				return nil, e
			}
			return updateResult, nil
		},
		Description: "新增/更新配置，响应成功更新条数",
	},
}
