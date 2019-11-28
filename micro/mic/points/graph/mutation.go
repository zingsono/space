package graph

import (
	"github.com/graphql-go/graphql"
)

func init() {
	MutationFields(_mutation)
}

var _mutation = graphql.Fields{

	"create": &graphql.Field{
		Name: "",
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"reamrk": &graphql.ArgumentConfig{
				Type:         graphql.String,
				DefaultValue: nil,
				Description:  "备注信息",
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {

		},
		DeprecationReason: "",
		Description:       "创建积分账户、返回账户编号",
	},
}
