package graph

import (
	"time"
)

/*
积分账户集合(ms_points_account)
账户编号5位数字、发行账户编号、创建时间、状态（1=正常 2=禁用 9=注销）、积分有效期天数（天）、剩余积分、授信透支积分数、账户备注、积分比率（商家调整积分比率，直接影响商品价格）

积分可用集合ms_points_usable
记录ID_id、账户编号、充值时间、过期时间、本次充值积分数、剩余积分、预充值积分数、预扣积分数、批次号

积分交易记录(ms_points_trade)
交易流水_id、账户编号、发行账户编号（充值与回收才有值）、交易积分数、剩余积分数、交易类型、交易时间、交易备注、业务系统标识名、交易业务系统订单号、积分可用记录ID、批次号（批量充值操作标记）

积分批次号集合(ms_points_batch)
批次号、开始时间、完成时间、操作描述（如：批量赠送积分）、状态（0=创建批次 1=执行中 2=完成 =3失败）、结果简述、下载操作结果明细表格

*/

// 积分账户状态（1=正常 2=禁用 9=注销）
type AccountStatus string

const (
	POINTS_ACCOUNT_STATUS_NORMAL    AccountStatus = "1"
	POINTS_ACCOUNT_STATUS_FORBIDDEN AccountStatus = "2"
	POINTS_ACCOUNT_STATUS_INVALID   AccountStatus = "9"
)

// 积分账户信息 ms_points_account
type PointsAccount struct {
	Muid      string        `bson:"muid" json:"muid"`           // 当前用户的积分账户， muid与ouid标识一个积分账户
	Ouid      string        `bson:"ouid" json:"ouid"`           // 发行机构会员用户ID
	Points    int64         `bson:"points" json:"points"`       // 账户剩余可用积分
	Amount    int64         `bson:"amount" json:"amount"`       // 账户剩余积分可转换的现金金额
	Status    AccountStatus `bson:"status" json:"status"`       // 状态（1=正常 2=禁用 9=注销）
	ValidDays int64         `bson:"validDays" json:"validDays"` // 当前账户发行的积分有效期天数
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"` // 创建时间
	UpdatedAt time.Time     `bson:"updatedAt" json:"updatedAt"` // 更新时间
	Remark    string        `bson:"remark" json:"remark"`       // 备注信息
}

// 积分批次号集合(ms_points_batch)
type PointsBatch struct {
	Ouid    string `bson:"ouid" json:"ouid"` // 发行机构会员用户ID
	BatchId string //批次编号
	BegTime string //开始时间
	EndTime string // 结束时间
	Remark  string //批次操作备注
	Status  string //状态状态（0=创建批次 1=执行中 2=完成 =3失败）
	Result  string //操作结果
	DownUrl string //下载地址链接
}

// 积分账户预操作记录（ms_points_withhold）
/*type PointsWithhold struct {
	Id string
	BatchId string //批次编号
}*/

// 记录预操作流水，扣减账户预操作积分
// 先只记录预操作消费流水，操作完成，统计所有预操作记录总积分，一次扣除。如果是否，则所有记录回滚

// 账户可用积分 ms_points_usable
// 记录ID_id、账户编号、充值时间、过期时间、积分数、积分比率、充值交易流水
type PointsUsable struct {
	UsableId  string    `bson:"usableId" json:"usableId"` // 可用积分记录ID
	Muid      string    `bson:"muid" json:"muid"`         // 当前用户的积分账户， muid与ouid标识一个积分账户
	Ouid      string    `bson:"ouid" json:"ouid"`         // 发行机构会员用户ID
	CreatedAt time.Time `bson:"time" json:"createdAt"`    // 创建时间
	Expires   time.Time `bson:"expires" json:"expires"`   // 过期时间
	Points    int64     `bson:"points" json:"points"`     // 当前批次可用积分数
	Ratio     int64     `bson:"ratio" json:"ratio"`       // 积分比例 1分钱=多少积分
	Payouts   int64     `bson:"payouts" json:"payouts"`   // 已支出积分数
	BatchId   string    // 批次编号，可选，可用于统计某一批次积分使用情况、撤回某一个批次的积分发放
}

// 积分交易 ms_points_trade
// 交易流水、账户编号、交易积分数、剩余积分数、交易类型、交易时间、交易备注、业务系统标识名、交易业务系统订单号
type PointsTrade struct {
	Muid      string    `bson:"muid" json:"muid"` // 当前用户的积分账户， muid与ouid标识一个积分账户
	Ouid      string    `bson:"ouid" json:"ouid"` // 发行机构会员用户ID
	BatchId   string    //批次编号
	Balance   int64     `bson:"balance" json:"balance"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"` // 更新时间
	Remark    string    `bson:"remark" json:"remark"`
}
