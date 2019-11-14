// 功能菜单
package model

type (
	// 菜单编号、菜单名称、图标、链接、排序、备注、上级编号、服务名、状态（1=显示 0=隐藏）、操作按钮（[编号、名称、备注]）、接口数据权限（）
	// mid,title,icon,link,sort,remark,service,status,operate,permissions
	MsMenu struct {
		Mid         string        `json:"mid"`
		Title       string        `json:"title"`
		Icon        string        `json:"icon"`
		Link        string        `json:"link"`
		Sort        string        `json:"sort"`
		Service     string        `json:"service"`
		Status      string        `json:"status"`
		Operate     []Operate     `json:"operate"`
		Permissions []Permissions `json:"permissions"`
		Remark      string        `json:"remark"`
		CreatedAt   string        `json:"createdAt"`
		CreatedBy   string        `json:"createdBy"`
		UpdatedAt   string        `json:"updatedAt"`
		UpdatedBy   string        `json:"updatedBy"`
	}

	// 操作按钮
	Operate struct {
		Id    string `json:"id"`
		Label string `json:"label"`
		Tip   string `json:"tip"`
	}

	// 功能菜单数据权限
	Permissions struct {
		Query []string `json:"query"`
	}
)
