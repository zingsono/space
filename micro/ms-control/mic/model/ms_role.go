// 操作员角色
package model

type (
	// 角色编号、角色名称、备注、创建时间、更新时间、功能权限（{服务名：菜单编号数组}）、操作权限（{菜单编号：按钮编号数组}）
	// rid,title,remark,menus,operate
	MsRole struct {
		Rid       string   `json:rid`
		Title     string   `json:title`
		Menus     []string `json:menus`
		Operate   []string `json:operate`
		Remark    string   `json:remark`
		CreatedAt string   `json:createdAt`
		CreatedBy string   `json:createdBy`
		UpdatedAt string   `json:updatedAt`
		UpdatedBy string   `json:updatedBy`
	}
)
