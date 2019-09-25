package hconfig

// Mongo 服务配置
type Mongo struct {
	Uri string `json:"uri"`
	Db  string `json:"db"`
}
