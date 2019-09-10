package model

import (
	"time"
)

// 配置信息管理
// 字段：服务名、配置JSON内容、备注、更新时间、创建时间
type MsConfig struct {
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	Remark    string    `json:"remark"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewMsConfig(name string, value string, remark string) *MsConfig {
	return &MsConfig{Name: name, Value: value, Remark: remark}
}

func (*MsConfig) CollectionName() string {
	return "ms_config"
}

// 插入单条记录
func (conf *MsConfig) Add(name string, value string, remark string) *MsConfig {
	conf.Name = name
	conf.Value = value
	conf.Remark = remark
	conf.UpdatedAt = time.Now()
	conf.CreatedAt = time.Now()
	return conf
}

// 插入一条记录，返回成功条数
func (conf *MsConfig) InserOne() (int, error) {
	conf.UpdatedAt = time.Now()
	conf.CreatedAt = time.Now()

	// config.MongoClient.Database("test").Collection("test").InsertOne()
	return 1, nil
}
