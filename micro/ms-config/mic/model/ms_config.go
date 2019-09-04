package model

import (
	"time"

	ql "github.com/graphql-go/graphql"
)

// 配置信息管理
// 字段：服务名、配置JSON内容、备注、更新时间、创建时间
type MsConfig struct {
	Name      string      `json:"name"`
	Value     string      `json:"value"`
	Remark    string      `json:"remark"`
	UpdatedAt time.Timer  `json:"updatedAt"`
	CreatedAt time.Ticker `json:"createdAt"`
}

// Graphql Query `ms_config`
var MsConfigQueryFields = ql.Fields{
	"configFindOne": &ql.Field{
		Name: "",
		Type: ConfigObject,
		Args: ql.FieldConfigArgument{
			"name": &ql.ArgumentConfig{
				Type:         ql.String,
				DefaultValue: nil,
				Description:  "服务名，根据服务名查询指定服务配置信息",
			},
		},
		Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
			name := p.Args["name"].(string)
			return &MsConfig{Name: name, Value: "{}"}, e
		},
		DeprecationReason: "",
		Description:       "单条数据查询",
	},
	"configFind": &ql.Field{
		Name: "",
		Type: ql.NewList(ConfigObject),
		Args: nil,
		Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
			return i, e
		},
		DeprecationReason: "",
		Description:       "列表查询",
	},
}

// Graphql Mutation `ms_config`
var MsConfigMutationFields = ql.Fields{
	"config": &ql.Field{
		Name: "",
		Type: ql.String,
		Args: nil,
		Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
			return i, e
		},
		DeprecationReason: "",
		Description:       "",
	},
}

var ConfigObject = ql.NewObject(ql.ObjectConfig{
	Name:       "ConfigObject",
	Interfaces: nil,
	Fields: ql.Fields{
		"name": &ql.Field{
			Name: "",
			Type: ql.String,
			Args: nil,
			Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
				return "", e
			},
			DeprecationReason: "",
			Description:       "",
		},
		"value": &ql.Field{
			Name: "",
			Type: ql.String,
			Args: nil,
			Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
				return "", e
			},
			DeprecationReason: "",
			Description:       "",
		},
		"remark": &ql.Field{
			Name: "",
			Type: ql.String,
			Args: nil,
			Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
				return "", e
			},
			DeprecationReason: "",
			Description:       "",
		},
		"updatedAt": &ql.Field{
			Name: "",
			Type: ql.String,
			Args: nil,
			Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
				return "", e
			},
			DeprecationReason: "",
			Description:       "",
		},
		"createdAt": &ql.Field{
			Name: "",
			Type: ql.DateTime,
			Args: nil,
			Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
				t := time.Now().Format(time.RFC3339)
				date, e := time.Parse(t, time.RFC3339)
				return date, e
			},
			DeprecationReason: "",
			Description:       "",
		},
	},
	IsTypeOf:    nil,
	Description: "配置信息查询",
})
