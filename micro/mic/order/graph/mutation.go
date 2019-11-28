package graph

import (
	"log"

	"github.com/graphql-go/graphql"
)

func init() {
	MutationFields(_mutation)
}

var _mutation = graphql.Fields{
	"order": &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name:       "OrderMutationType",
			Interfaces: nil,
			Fields: graphql.Fields{
				"create": &graphql.Field{
					Type: graphql.NewObject(graphql.ObjectConfig{
						Name:       "OrderCreateResultType",
						Interfaces: nil,
						Fields: graphql.Fields{
							"orderId": &graphql.Field{Type: graphql.String, Description: "订单编号"},
							"payUrl":  &graphql.Field{Type: graphql.String, Description: "订单支付网关URL，前端收到创建订单成功的结果直接跳转此链接做支付，无需做任何拼接参数处理"},
						},
						IsTypeOf:    nil,
						Description: "订单创建结果类型",
					}),
					Args: graphql.FieldConfigArgument{
						"mallId":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String), DefaultValue: nil, Description: "商城编号"},
						"mallUid":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String), DefaultValue: nil, Description: "商城所属用户会员ID"},
						"uid":          &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String), DefaultValue: nil, Description: "下单用户会员ID"},
						"amount":       &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: nil, Description: "订单实际支付金额"},
						"points":       &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: nil, Description: "订单实际支付积分"},
						"remark":       &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: nil, Description: "订单备注信息"},
						"payResultUrl": &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: nil, Description: "收银台支付完成后跳转链接，拼接参数：?result={\"code\":\"00000\",\"msg\":\"ok\"}  ,等于00000则提示支付成功，否则提示支付异常显示异常信息"},
						"details": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.NewList(graphql.NewInputObject(graphql.InputObjectConfig{
								Name: "DetailsInputObject",
								Fields: graphql.InputObjectConfigFieldMap{
									"prodId":    &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String), DefaultValue: nil, Description: "产品编号"},
									"prodName":  &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String), DefaultValue: nil, Description: "产品名称"},
									"imgUrl":    &graphql.InputObjectFieldConfig{Type: graphql.String, DefaultValue: nil, Description: "产品图片"},
									"unitPrice": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int), DefaultValue: nil, Description: "产品单价，单位分"},
									"num":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int), DefaultValue: nil, Description: "购买数量"},
								},
								Description: "订单明细输入对象",
							}))),
							Description: "订单明细",
						},
					},
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
						return "ttt", e
					},
					Description: "创建订单 ",
				},
				// 寄出
				// 确认收货
				// 申请退款
				// 审核退款
				// 更新快递信息

			},
			IsTypeOf:    nil,
			Description: "订单操作类型",
		}),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			log.Print("OrderMutationType............")
			return "", e
		},
		Description: "订单服务操作",
	},
}
