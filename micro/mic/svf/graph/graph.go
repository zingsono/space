package graph

import (
	"github.com/graphql-go/graphql"
	"github.com/zingsono/space/micro/lib/hgraph"
)

var (
	GraphqlHttpHandler = hgraph.GraphqlHttpHandler
	QueryFields        = hgraph.MergeQueryFields
	MutationFields     = hgraph.MergeMutationFields
)

// 分页请求参数定义
var PageArgument = graphql.FieldConfigArgument{
	"limit": &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: 20, Description: "一次返回记录行数，默认20"},
	"skip":  &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: 0, Description: "跳过记录行数"},
}

// 合并参数
func Argument(args ...graphql.FieldConfigArgument) graphql.FieldConfigArgument {
	var newArgument = make(graphql.FieldConfigArgument)
	for _, item := range args {
		for k, v := range item {
			newArgument[k] = v
		}
	}
	return newArgument
}

// 响应状态码
var ResponseCodeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ResponseCode",
	Fields: graphql.Fields{
		"code": &graphql.Field{Type: graphql.String, Description: "响应码"},
		"msg":  &graphql.Field{Type: graphql.String, Description: "响应描述"},
	},
	Description: "响应状态码，‘00000’成功，其它错误码见接口描述",
})
