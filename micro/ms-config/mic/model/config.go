package model

import (
	"time"
)

// 字段：服务名、配置JSON内容、备注、更新时间、创建时间
type MsConfig struct {
	Name      string      `json:"name"`
	Value     string      `json:"value"`
	Remark    string      `json:"remark"`
	UpdatedAt time.Timer  `json:"updatedAt"`
	CreatedAt time.Ticker `json:"createdAt"`
}



