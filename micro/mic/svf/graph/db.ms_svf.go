package graph

import (
	"time"
)

/*

- 现金账户信息表

用户编号、余额、授信额度、累计充值、账户状态（1=正常 2=禁用 9=注销）、开户时间、最后交易流水、最后交易时间

- 现金交易流水表

用户编号、交易流水号、交易时间、账务日期、交易类型（充值、消费、退款、提现、冲正）、冲正交易流水号、交易金额、账户余额、业务编号、业务订单号、备注

- 账户操作流水表

用户编号、操作流水号、操作时间、类型（10=账户开户、11=禁用账户、12=启用账户、99=注销账户）





*/

// 储值账户信息
type SvfAccount struct {
	Uid        string // 用户编号
	Balance    int64  // 余额，单位分
	Accumulate int64  // 累计充值金额，单位分
	Credit     int64  // 授信金额，单位分
	Status     string // （1=正常 2=禁用 9=注销）
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// 储值账户交易
type SvfTrade struct {
	Id        string    // 交易流水号
	Uid       string    // 用户编号
	Type      string    // 交易类型（充值、消费、退款、提现、冲正）
	Amount    int64     // 交易金额
	Balance   int64     // 账户余额
	BizId     string    // 业务编号，业务系统的订单号或者交易流水
	BatchId   string    // 交易批次编号，由业务系统定义。用于标记批量交易。 批量扣款时，先记录所有交易流水，统计交易金额，对账户余额扣款，如果成功，则提交所有预操作流水
	Status    string    // 交易状态码： 00=成功 10=预操作  20=余额不足 30=账户禁用
	Remark    string    // 交易备注信息
	CreatedAt time.Time // 创建交易时间
	UpdatedAt time.Time // 交易更新时间
}

// 储值账户日志
type SvfLog struct {
	Id        string // 操作流水号
	Uid       string // 用户编号
	TradeId   string // 交易流水号
	Msg       string // 操作说明
	CreatedAt string // 创建时间
}
