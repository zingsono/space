package graph

import (
	"log"

	"github.com/graphql-go/graphql"
)

func init() {
	QueryFields(_query)
}

var _query = graphql.Fields{
	"order": &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name:       "OrderQueryType",
			Interfaces: nil,
			Fields: graphql.Fields{
				"single": &graphql.Field{
					Type: OrderInfoType,
					Args: Argument(orderIdArg),
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
						log.Println(p.Args)

						return &Order{OrderId: "123123123"}, e
					},
					Description: "商城信息查询",
				},
				"total": &graphql.Field{
					Type:        graphql.Int,
					Description: "会员总数，最大值5000，不能依赖此字段做统计。",
					Args:        Argument(orderListArg),
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
						return 5000, e
					},
				},
				"list": &graphql.Field{
					Type: graphql.NewList(OrderInfoType),
					Args: Argument(PageArgument, orderListArg),
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
		Description:       "订单查询服务",
	},
}

// 单条记录查询参数定义
var orderIdArg = graphql.FieldConfigArgument{
	"orderId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String), DefaultValue: nil, Description: "订单编号"},
}

// 结果集查询参数定义
var orderListArg = graphql.FieldConfigArgument{
	"orderId": &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: nil, Description: "订单编号"},
	"uid":     &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: nil, Description: "下单用户ID"},
	"mallId":  &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: nil, Description: "商城编号"},
}

// 订单信息类型
var OrderInfoType = graphql.NewObject(graphql.ObjectConfig{
	Name:       "OrderInfoType",
	Interfaces: nil,
	Fields: graphql.Fields{
		"orderId":        &graphql.Field{Type: graphql.String, Description: "订单编号"},
		"uid":            &graphql.Field{Type: graphql.String, Description: "下单用户ID"},
		"mallId":         &graphql.Field{Type: graphql.String, Description: "商城编号"},
		"mallUid":        &graphql.Field{Type: graphql.String, Description: "商城用户ID ,标记此订单谁销售"},
		"providerUid":    &graphql.Field{Type: graphql.String, Description: "供货商用户ID"},
		"status":         &graphql.Field{Type: graphql.String, Description: "订单状态(0=待支付  1=已支付  2=已退款 3=部分已退款 8=已完成 9=已删除 )"},
		"createdAt":      &graphql.Field{Type: graphql.String, Description: "订单创建时间(yyyyMMddHHmmss)"},
		"finishTime":     &graphql.Field{Type: graphql.String, Description: "订单完成时间(yyyyMMddHHmmss)"},
		"remark":         &graphql.Field{Type: graphql.String, Description: "订单备注信息"},
		"payTime":        &graphql.Field{Type: graphql.String, Description: "支付时间(yyyyMMddHHmmss)"},
		"payExpires":     &graphql.Field{Type: graphql.String, Description: "自动支付过期时间"},
		"supplyAmount":   &graphql.Field{Type: graphql.String, Description: "供应商看到的订单价格，根据商品价格与数量计算"},
		"originalAmount": &graphql.Field{Type: graphql.String, Description: "订单原始价格，没有使用积分，没有使用优惠券的价格"},
		"amount":         &graphql.Field{Type: graphql.String, Description: "实际支付订单总金额，单位分"},
		"points":         &graphql.Field{Type: graphql.String, Description: "实际支付总积分 ，根据用户的积分数，查询出可抵扣的金额加上实际支付的订单金额，如果等于原价，则通过。"},

		"refund": &graphql.Field{Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "RefundType",
			Fields: graphql.Fields{
				"status":      &graphql.Field{Type: graphql.String, Description: "订单申请退款状态（10=未申请退款 11=申请整单退款 12=整单退款成功 13=整单退款拒绝 21=申请部分退款  22=部分退款成功  23=部分退款拒绝）"},
				"amount":      &graphql.Field{Type: graphql.Int, Description: " 订单退款金额"},
				"points":      &graphql.Field{Type: graphql.Int, Description: "订单退款积分"},
				"applyTime":   &graphql.Field{Type: graphql.String, Description: "申请退款时间"},
				"applyReason": &graphql.Field{Type: graphql.String, Description: "申请原因"},
				"images":      &graphql.Field{Type: graphql.NewList(graphql.String), Description: "图片"},
				"auditTime":   &graphql.Field{Type: graphql.String, Description: "审核时间"},
				"auditResult": &graphql.Field{Type: graphql.String, Description: "审核结果说明"},
				"express":     &graphql.Field{Type: ExpressType, Description: "退货快递信息"},
			},
			Description: "订单退款信息类型",
		}), Description: "退款信息"},
		"express": &graphql.Field{Type: ExpressType, Description: "订单快递物流信息"},
		"details": &graphql.Field{Type: graphql.NewObject(graphql.ObjectConfig{
			Name:       "DetailsType",
			Interfaces: nil,
			Fields: graphql.Fields{
				"prodId":         &graphql.Field{Type: graphql.String, Description: "产品编号"},
				"prodName":       &graphql.Field{Type: graphql.String, Description: "产品名"},
				"unitPrice":      &graphql.Field{Type: graphql.Int, Description: "销售单价"},
				"stockPrice":     &graphql.Field{Type: graphql.Int, Description: "进货单价"},
				"pointsPayRatio": &graphql.Field{Type: graphql.Int, Description: "积分占产品单价的比率，产品单价的百分之多少可以使用积分支付 如：20%，此字段值为20 ，单价为100元的产品，有20元可以使用积分支付"},
				"num":            &graphql.Field{Type: graphql.Int, Description: "数量"},
				"refundNum":      &graphql.Field{Type: graphql.Int, Description: "退货数量（等于0代表未退货）"},
				"imgUrl":         &graphql.Field{Type: graphql.String, Description: "产品图片"},
			},
			IsTypeOf:    nil,
			Description: "订单明细类型",
		}), Description: "订单产品明细"},
	},
	IsTypeOf:    nil,
	Description: "订单信息类型",
})

// 快递信息
var ExpressType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ExpressType",
	Fields: graphql.Fields{
		"status":         &graphql.Field{Type: graphql.String, Description: "快递状态（10=待发货 11=等待取件 21=配送中  29=已签收  31=拒签退货配送中  39=退货已签收 44=快递丢失）"},
		"senderName":     &graphql.Field{Type: graphql.String, Description: "发件人姓名"},
		"senderTel":      &graphql.Field{Type: graphql.String, Description: "发件人电话"},
		"senderAddress":  &graphql.Field{Type: graphql.String, Description: "发件人地址"},
		"senderTime":     &graphql.Field{Type: graphql.String, Description: "发货时间"},
		"sendBackTime":   &graphql.Field{Type: graphql.String, Description: "退货签收时间"},
		"receiveName":    &graphql.Field{Type: graphql.String, Description: "收件人姓名"},
		"receiveTel":     &graphql.Field{Type: graphql.String, Description: "收件人电话"},
		"receiveAddress": &graphql.Field{Type: graphql.String, Description: "收件地址"},
		"receiveTime":    &graphql.Field{Type: graphql.String, Description: "签收时间"},
		"num":            &graphql.Field{Type: graphql.String, Description: "快递单号"},
		"company":        &graphql.Field{Type: graphql.String, Description: "快递公司名称"},
		"deliverName":    &graphql.Field{Type: graphql.String, Description: "配送员姓名"},
		"deliverTel":     &graphql.Field{Type: graphql.String, Description: "配送员电话"},
	},
	Description: "快递信息类型",
})
