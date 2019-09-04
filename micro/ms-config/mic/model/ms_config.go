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
	"config": &ql.Field{
		Name: "",
		Type: MsConfigQueryType,
		Args: nil,
		Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
			return i, e
		},
		DeprecationReason: "",
		Description:       "配置管理",
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

var MsConfigQueryType = ql.NewObject(ql.ObjectConfig{
	Name:       "ConfigQueryType",
	Interfaces: nil,
	Fields: ql.Fields{
		"name": &ql.Field{
			Name: "",
			Type: ql.String,
			Args: nil,
			Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
				return i, e
			},
			DeprecationReason: "",
			Description:       "",
		},
		"value": &ql.Field{
			Name: "",
			Type: ql.String,
			Args: nil,
			Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
				return i, e
			},
			DeprecationReason: "",
			Description:       "",
		},
		"remark": &ql.Field{
			Name: "",
			Type: ql.String,
			Args: nil,
			Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
				return i, e
			},
			DeprecationReason: "",
			Description:       "",
		},
		"updatedAt": &ql.Field{
			Name: "",
			Type: ql.DateTime,
			Args: nil,
			Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
				return i, e
			},
			DeprecationReason: "",
			Description:       "",
		},
		"createdAt": &ql.Field{
			Name: "",
			Type: ql.DateTime,
			Args: nil,
			Resolve: func(p ql.ResolveParams) (i interface{}, e error) {
				return i, e
			},
			DeprecationReason: "",
			Description:       "",
		},
	},
	IsTypeOf:    nil,
	Description: "配置信息查询",
})
