package graph

import (
	"log"
	"time"

	"github.com/graphql-go/graphql"
)

func init() {
	QueryFields(_query)
}

var _query = graphql.Fields{
	"svf": &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name:       "SvfQueryType",
			Interfaces: nil,
			Fields: graphql.Fields{
				"single": &graphql.Field{
					Type: SvfInfoType,
					Args: Argument(SvfIdArg),
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
						log.Println(p.Args)

						acc := &SvfAccount{
							Uid:        "100001",
							Balance:    100000,
							Accumulate: 100000,
							Credit:     100000,
							Advance:    100000,
							Status:     "1",
							CreatedAt:  time.Now(),
							UpdatedAt:  time.Now(),
						}
						return acc, e
					},
					Description: "储值账户查询",
				},
				"total": &graphql.Field{
					Type:        graphql.Int,
					Description: "会员总数，最大值5000，不能依赖此字段做统计。",
					Args:        Argument(SvfListArg),
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
						return 5000, e
					},
				},
				"list": &graphql.Field{
					Type: graphql.NewList(SvfInfoType),
					Args: Argument(PageArgument, SvfListArg),
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
						log.Print("MemberInfoType dataset............")
						log.Println(p.Args)
						return i, e
					},
					Description: "列表查询",
				},
			},
			IsTypeOf:    nil,
			Description: "账户查询类型",
		}),
		Args: nil,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			log.Print("MemberType............")
			return "", e
		},
		DeprecationReason: "",
		Description:       "储值账户服务",
	},
}

// 单条记录查询参数定义
var SvfIdArg = graphql.FieldConfigArgument{
	"uid": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String), DefaultValue: nil, Description: "会员编号"},
}

// 结果集查询参数定义
var SvfListArg = graphql.FieldConfigArgument{
	"uid": &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: nil, Description: "会员编号"},
}

var SvfInfoType = graphql.NewObject(graphql.ObjectConfig{
	Name:       "SvfInfoType",
	Interfaces: nil,
	Fields: graphql.Fields{
		"uid":        &graphql.Field{Type: graphql.String, Description: "用户编号"},
		"balance":    &graphql.Field{Type: graphql.Int, Description: "余额，单位分"},
		"accumulate": &graphql.Field{Type: graphql.Int, Description: "累计充值金额，单位分"},
		"credit":     &graphql.Field{Type: graphql.Int, Description: "授信金额，单位分"},
		"advance":    &graphql.Field{Type: graphql.Int, Description: "预扣款金额，用于交易锁定金额，单位分"},
		"status":     &graphql.Field{Type: graphql.String, Description: "账户状态（1=正常 2=禁用 9=注销）"},
		"createdAt":  &graphql.Field{Type: graphql.String, Description: "创建时间"},
		"updatedAt":  &graphql.Field{Type: graphql.String, Description: "更新时间"},
		"trades": &graphql.Field{
			Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
				Name:       "SvfTradeType",
				Interfaces: nil,
				Fields: graphql.Fields{
					"tid":          &graphql.Field{Type: graphql.String, Description: "交易编号"},
					"cid":          &graphql.Field{Type: graphql.String, Description: "客户端交易编号号，用于标记唯一交易，一般使用业务系统订单号"},
					"uid":          &graphql.Field{Type: graphql.String, Description: "用户编号"},
					"bulkId":       &graphql.Field{Type: graphql.String, Description: "批量交易编号,非批量交易为空"},
					"tradeType":    &graphql.Field{Type: graphql.String, Description: "交易类型（1=充值、2=消费、3=提现）"},
					"amount":       &graphql.Field{Type: graphql.Int, Description: "交易金额"},
					"balance":      &graphql.Field{Type: graphql.Int, Description: "账户余额"},
					"refundAmount": &graphql.Field{Type: graphql.Int, Description: "退款金额，支持全额与部分退款"},
					"remark":       &graphql.Field{Type: graphql.String, Description: "交易备注信息"},
					"createdAt":    &graphql.Field{Type: graphql.String, Description: "创建交易时间"},
					"updatedAt":    &graphql.Field{Type: graphql.String, Description: "交易更新时间"},
					"code":         &graphql.Field{Type: graphql.String, Description: "交易状态码： 00000=成功 10010=预操作 10011=预操作撤销  10020=余额不足 10031=账户被禁用 10032=账户不存在 10041=交易已冲正（只成功交易）"},
					"msg":          &graphql.Field{Type: graphql.String, Description: "交易状态码描述"},
				},
				IsTypeOf:    nil,
				Description: "账户交易记录",
			})),
			Args: Argument(PageArgument, graphql.FieldConfigArgument{
				"tid":    &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: nil, Description: "交易编号"},
				"cid":    &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: nil, Description: "客户端交易编号"},
				"uid":    &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: nil, Description: "用户编号"},
				"bulkId": &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: nil, Description: "批量操作编号"},
			}),
			Resolve: func(p graphql.ResolveParams) (i interface{}, err error) {
				trades := []*SvfTrade{&SvfTrade{
					Tid:          "123123123",
					Cid:          "123123123",
					Uid:          "123123123",
					BulkId:       "34234234",
					TradeType:    SVF_TRADE_TYPE_T2,
					Amount:       1110,
					Balance:      1110,
					RefundAmount: 1110,
					Remark:       "固定值数据",
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					Code:         SVF_TRADE_CODE_00000,
					Msg:          "ok",
				}}
				return trades, err
			},
			Description: "账户交易记录，分页查询"},
	},
	IsTypeOf:    nil,
	Description: "储值账户信息类型",
})
