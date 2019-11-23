package model

import (
	"time"
)

// 积分交易 ms_points_trade
// 交易流水、账户编号、交易积分数、剩余积分数、交易类型、交易时间、交易备注、业务系统标识名、交易业务系统订单号
type PointsTrade struct {
	AccountNo string `bson:"accountNo" json:"accountNo"`
	Balance   int64  `bson:"balance" json:"balance"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Remark    string `bson:"remark" json:"remark"`
}
