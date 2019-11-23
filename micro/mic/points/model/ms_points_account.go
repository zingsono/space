package model

import (
	"time"
)

// 积分账户信息 ms_points_account
// 账户编号、剩余积分、创建时间、状态（1=正常 2=禁用 9=注销）
type PointsAccount struct {
	AccountNo string `bson:"accountNo" json:"accountNo"`
	Balance   int64  `bson:"balance" json:"balance"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Remark    string `bson:"remark" json:"remark"`
}
