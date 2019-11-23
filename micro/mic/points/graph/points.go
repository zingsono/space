package graph

import (
	"log"

	"github.com/graphql-go/graphql"
)

func init() {

	// 查询操作
	QueryFields(graphql.Fields{
		"account": &graphql.Field{
			Name: "",
			Type: graphql.NewObject(graphql.ObjectConfig{
				Name:       "PointsAccount",
				Interfaces: nil,
				Fields: graphql.Fields{
					"id": &graphql.Field{
						Name: "",
						Type: graphql.String,
						Args: nil,
						Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
							// i = "12312"
							return i, e
						},
						DeprecationReason: "",
						Description:       "",
					},
					"banalce": &graphql.Field{
						Name:              "",
						Type:              graphql.Int,
						Args:              nil,
						Resolve:           nil,
						DeprecationReason: "",
						Description:       "",
					},
				},
				IsTypeOf:    nil,
				Description: "积分账户",
			}),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: nil,
					Description:  "账户编号",
				},
				"reamrk": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: nil,
					Description:  "账户备注信息",
				},
			},
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				log.Print(p.Context.Value("token"))
				m := make(map[string]interface{})
				m["id"] = "123"
				m["banalce"] = 321
				return m, e
			},
			DeprecationReason: "",
			Description:       "积分账户",
		},
	})

	// 更新操作
	MutationFields(graphql.Fields{
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
	})

}
