package graph

import (
	"time"
)

/*
会员订单信息(ms_order)
订单信息（订单号、订单金额、订单、订单明细、支付状态、发货状态、支付超时时间、支付完成时间、订单创建时间、下单会员ID、商城ID）

退款优先退积分，积分不够再退支付的金额

确认订单页面，用户查询剩余积分与可抵扣金额，订单支付金额最多可使用积分的金额。

下单直接预扣除积分，如果预扣除不成功，则下单失败。


积分不走收银台，因为进入收银台已经提交订单成功，系统如果积分不足，则订单无法继续。
*/

// 会员订单信息(ms_order)
type Order struct {
	// 订单通用基本信息字段
	OrderId     string     `bson:"orderId" json:"orderId"`         // 订单编号
	Uid         string     `bson:"uid" json:"uid"`                 // 下单用户ID
	MallId      string     `bson:"mallId" json:"mallId"`           // 商城编号
	MallUid     string     `bson:"mallUid" json:"mallUid"`         // 商城用户ID ,标记此订单谁销售，也是下单用户的积分发行商户
	ProviderUid string     `bson:"providerUid" json:"providerUid"` // 供货商用户ID，标记此订单由谁供货并发货
	Status      string     `bson:"status" json:"status"`           // 订单状态(0=待支付  1=已支付  2=已退款 3=部分已退款 8=已完成 9=已删除 )
	CreatedAt   time.Time  `bson:"createdAt" json:"createdAt"`     // 订单创建时间
	FinishTime  time.Time  `bson:"finishTime" json:"finishTime"`   // 订单完成时间
	Remark      string     `bson:"remark" json:"remark"`           // 订单备注信息
	Refund      *Refund    `bson:"refund" json:"refund"`           // 退款信息
	Pay         *Pay       `bson:"pay" json:"pay"`                 // 订单支付信息
	Express     *Express   `bson:"express" json:"express"`         // 订单快递物流
	Details     []*Details `bson:"details" json:"details"`         // 订单产品明细

}

// 订单支付信息
type Pay struct {
	PayTime        time.Time `bson:"payTime" json:"payTime"`               // 支付时间
	PayExpires     time.Time `bson:"payExpires" json:"payExpires"`         // 自动支付过期时间
	OriginalAmount int64     `bson:"originalAmount" json:"originalAmount"` // 订单原始价格，没有使用积分，没有使用优惠券的价格
	PayAmount      int64     `bson:"payAmount" json:"payAmount"`           // 实际支付订单总金额，单位分
	PayPoints      int64     `bson:"payPoints" json:"payPoints"`           // 实际支付总积分 ，根据用户的积分数，查询出可抵扣的金额加上实际支付的订单金额，如果等于原价，则通过。
}

// 退款信息
type Refund struct {
	Status      string    `bson:"status" json:"status"`           // 订单申请退款状态（10=未申请退款 11=申请整单退款 12=整单退款成功 13=整单退款拒绝 21=申请部分退款  22=部分退款成功  23=部分退款拒绝）
	Amount      int64     `bson:"amount" json:"amount"`           // 订单总退款金额
	ApplyTime   time.Time `bson:"applyTime" json:"applyTime"`     // 申请退款时间
	ApplyReason string    `bson:"applyReason" json:"applyReason"` // 申请原因
	Images      []string  `bson:"images" json:"images"`           // 图片
	AuditTime   time.Time `bson:"auditTime" json:"auditTime"`     // 审核时间
	AuditResult string    `bson:"auditResult" json:"auditResult"` // 审核结果说明
	Express     *Express  `bson:"express" json:"express"`         // 退货快递信息
}

// 订单快递信息
type Express struct {
	Status         string    `bson:"status" json:"status"`                 // 快递状态（10=待发货 11=等待取件 21=配送中  29=已签收  31=拒签退货配送中  39=退货已签收 44=快递丢失）
	SenderName     string    `bson:"senderName" json:"senderName"`         // 发件人姓名
	SenderTel      string    `bson:"senderTel" json:"senderTel"`           // 发件人电话
	SenderAddress  string    `bson:"senderAddress" json:"senderAddress"`   // 发件人地址
	SenderTime     time.Time `bson:"senderTime" json:"senderTime"`         // 发货时间
	SendBackTime   time.Time `bson:"sendBackTime" json:"sendBackTime"`     // 退货签收时间
	ReceiveName    string    `bson:"receiveName" json:"receiveName"`       // 收件人姓名
	ReceiveTel     string    `bson:"receiveTel" json:"receiveTel"`         // 收件人电话
	ReceiveAddress string    `bson:"receiveAddress" json:"receiveAddress"` // 收件地址
	ReceiveTime    time.Time `bson:"receiveTime" json:"receiveTime"`       // 签收时间
	Num            string    `bson:"num" json:"num"`                       // 快递单号
	Company        string    `bson:"company" json:"Company"`               // 快递公司名称
	DeliverName    string    `bson:"deliverName" json:"deliverName"`       // 配送员姓名
	DeliverTel     string    `bson:"deliverTel" json:"deliverTel"`         // 配送员电话

}

// 订单明细
// 计算出：可使用积分的金额
type Details struct {
	ProdId         string `bson:"prodId" json:"prodId"`                 // 产品编号
	ProdName       string `bson:"prodName" json:"prodName"`             // 产品名
	UnitPrice      int64  `bson:"unitPrice" json:"unitPrice"`           // 单价
	PointsPayRatio int64  `bson:"pointsPayRatio" json:"pointsPayRatio"` // 积分占产品单价的比率，产品单价的百分之多少可以使用积分支付 如：20%，此字段值为20 ，单价为100元的产品，有20元可以使用积分支付
	Num            int    `bson:"num" json:"num"`                       // 数量
	RefundNum      string `bson:"refundNum" json:"refundNum"`           // 退货数量（等于0代表未退货）
	ImgUrl         string `bson:"imgUrl" json:"imgUrl"`                 // 产品图片
}
