// 服务发现客户端
package hdiscover

// 服务注册结构体
type Instance struct {
	Id    string   `json:"id"`
	Name  string   `json:"name"`
	Host  string   `json:"host"`
	Port  string   `json:"port"`
	Watch []string `json:"watch"`
}
