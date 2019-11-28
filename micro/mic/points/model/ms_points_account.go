package model

import (
	"time"
)

// 积分账户状态（1=正常 2=禁用 9=注销）
type AccountStatus string

const (
	POINTS_ACCOUNT_STATUS_NORMAL    AccountStatus = "1"
	POINTS_ACCOUNT_STATUS_FORBIDDEN AccountStatus = "2"
	POINTS_ACCOUNT_STATUS_INVALID   AccountStatus = "9"
)

// 积分账户信息 ms_points_account
// 账户编号、剩余积分、创建时间、状态（1=正常 2=禁用 9=注销）、有效期天数、可兑换金额
type PointsAccount struct {
	AccountNo    string        `bson:"accountNo" json:"accountNo"`
	BalancePoint int64         `bson:"balancePoint" json:"balancePoint"`
	Status       AccountStatus `bson:"status" json:"status"`
	ValidityDays int           `bson:"validityDays" json:"validityDays"`
	CreatedAt    time.Time     `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time     `bson:"updatedAt" json:"updatedAt"`
	Remark       string        `bson:"remark" json:"remark"`
}

// 账户可用积分 ms_points_usable
// 记录ID_id、账户编号、充值时间、过期时间、积分数、积分比率、充值交易流水
type PointsUsable struct {
	Id        string    `bson:"_id" json:"id"`
	AccountNo string    `bson:"accountNo" json:"accountNo"`
	Time      time.Time `bson:"time" json:"time"`
	Expires   time.Time `bson:"expires" json:"expires"`
	Point     int64
}
