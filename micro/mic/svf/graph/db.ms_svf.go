package graph

import (
	"time"
)

/*
业务说明：
1. 账户查询，当账户不存在时，返回0；
2. 账户充值，如果不存在则自动创建；
3. 批量充值或者消费，先获取批量编号，然后调用充值，然后提交批量操作

*/

// 储值账户状态（1=正常 2=禁用 9=注销）
type AccountStatus string

const (
	SVF_ACCOUNT_STATUS_NORMAL    AccountStatus = "1"
	SVF_ACCOUNT_STATUS_FORBIDDEN AccountStatus = "2"
	SVF_ACCOUNT_STATUS_INVALID   AccountStatus = "9"
)

// 储值账户信息 （ms_svf_account）
type SvfAccount struct {
	Uid        string        `bson:"uid" json:"uid"`               // 用户编号
	Balance    int64         `bson:"balance" json:"balance"`       // 余额，单位分
	Accumulate int64         `bson:"accumulate" json:"accumulate"` // 累计充值金额，单位分
	Credit     int64         `bson:"credit" json:"credit"`         // 授信金额，单位分
	Advance    int64         `bson:"advance" json:"advance"`       // 预扣款金额，用于交易锁定金额，单位分
	Status     AccountStatus `bson:"status" json:"status"`         // （1=正常 2=禁用 9=注销）
	CreatedAt  time.Time     `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time     `bson:"updatedAt" json:"updatedAt"`
}

// 交易类型（1=充值、2=消费、3=提现）
type TradeType string

const (
	SVF_TRADE_TYPE_T1 TradeType = "1"
	SVF_TRADE_TYPE_T2 TradeType = "2"
	SVF_TRADE_TYPE_T3 TradeType = "3"
)

// 交易状态码： 00000=成功 10010=预操作 10011=预操作撤销  10020=余额不足 10031=账户被禁用 10032=账户不存在 10041=交易已冲正（只成功交易）
type TradeCode string

const (
	SVF_TRADE_CODE_00000 TradeCode = "00000"
	SVF_TRADE_CODE_10010 TradeCode = "10010"
	SVF_TRADE_CODE_10011 TradeCode = "10011"
	SVF_TRADE_CODE_10020 TradeCode = "10020"
	SVF_TRADE_CODE_10031 TradeCode = "10031"
	SVF_TRADE_CODE_10032 TradeCode = "10032"
	SVF_TRADE_CODE_10041 TradeCode = "10041"
)

// 储值账户交易 （ms_svf_trade）
type SvfTrade struct {
	Tid          string    `bson:"tid" json:"tid"`                   // 服务端交易编号号
	Cid          string    `bson:"cid" json:"cid"`                   // 客户端交易编号号，用于标记唯一交易，建议使用业务系统订单号
	Uid          string    `bson:"uid" json:"uid"`                   // 用户编号
	BulkId       string    `bson:"bulkId" json:"bulkId"`             // 批量交易编号,非批量交易为空
	TradeType    TradeType `bson:"tradeType" json:"tradeType"`       // 交易类型（充值、消费、提现）
	Amount       int64     `bson:"amount" json:"amount"`             // 交易金额
	Balance      int64     `bson:"balance" json:"balance"`           // 账户余额
	RefundAmount int64     `bson:"refundAmount" json:"refundAmount"` // 退款金额，支持全额与部分退款
	Remark       string    `bson:"remark" json:"remark"`             // 交易备注信息
	CreatedAt    time.Time `bson:"createdAt" json:"createdAt"`       // 创建交易时间
	UpdatedAt    time.Time `bson:"updatedAt" json:"updatedAt"`       // 交易更新时间
	Code         TradeCode `bson:"code" json:"code"`                 // 交易状态码： 00000=成功 10010=预操作 10011=预操作撤销  10020=余额不足 10031=账户被禁用 10032=账户不存在 10041=交易已冲正（只成功交易）
	Msg          string    `bson:"msg" json:"msg"`                   // 交易状态码描述
}

// 批量操作交易状态（1=成功 2=失败）
type BulkStatus string

const (
	SVF_BULK_STATUS_SUCCESS BulkStatus = "1"
	SVF_BULK_STATUS_FAIL    BulkStatus = "2"
)

// 账户批量交易（ms_svf_bulk）
// 批量交易需要先创建批次号
type SvfBulk struct {
	BulkId        string     `bson:"bulkId" json:"bulkId"`               // 批量批次编号
	Uid           string     `bson:"uid" json:"uid"`                     // 用户编号
	BegTime       time.Time  `bson:"begTime" json:"begTime"`             // 创建开始时间
	EndTime       time.Time  `bson:"endTime" json:"endTime"`             // 交易完成时间
	Status        BulkStatus `bson:"status" json:"status"`               // 交易状态（1=成功 2=失败）
	Remark        string     `bson:"remark" json:"remark"`               // 交易备注
	ExportFileUrl string     `bson:"exportFileUrl" json:"exportFileUrl"` // 批量交易明细导出下载地址
}

type HandType string

// 操作类型(10=账户充值申请 11=充值审核通过 12=充值审核拒绝 21=账户消费 22=消费退款 31=账户提现申请  90=冲正交易)
const (
	SVF_HAND_TYPE_E10 HandType = "10"
	SVF_HAND_TYPE_E11 HandType = "11"
	SVF_HAND_TYPE_E12 HandType = "12"
	SVF_HAND_TYPE_E21 HandType = "21"
	SVF_HAND_TYPE_E22 HandType = "22"
	SVF_HAND_TYPE_E31 HandType = "31"
	SVF_HAND_TYPE_E90 HandType = "90"
)

// 账户操作日志(ms_svf_log)
// 记录账户所有操作审计日志
type SvfLog struct {
	LogId     string   `bson:"logId" json:"logId"`         // 操作流水号
	Uid       string   `bson:"uid" json:"uid"`             // 用户编号
	Tid       string   `bson:"tid" json:"tid"`             // 服务端交易编号
	Cid       string   `bson:"cid" json:"cid"`             // 客户端交易编号
	HandType  HandType `bson:"handType" json:"handType"`   // 操作类型(10=账户充值申请 11=充值审核通过 12=充值审核拒绝 21=账户消费 22=消费退款 31=账户提现申请  90=冲正交易)
	Message   string   `bson:"message" json:"message"`     // 操作说明
	Amount    int64    `bson:"amount" json:"amount"`       // 交易金额
	CreatedAt string   `bson:"createdAt" json:"createdAt"` // 操作时间
	CreatedBy string   `bson:"createdBy" json:"createdBy"` // 操作人用户编号
}
