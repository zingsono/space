package model

type (
	//操作员用户: 用户编号、登录账号、邮箱、手机、密码（MD5）、昵称、头像、角色（数组）、备注、创建时间、创建人、更新时间、更新人
	//uid,loginId,email,mobile,password,nickname,avatar,roles,remark,createdAt,createdBy,updatedAt,updatedBy
	MsUser struct {
		Uid       string   `json:uid`
		LoginId   string   `json:loginId`
		Email     string   `json:email`
		Mobile    string   `json:mobile`
		Password  string   `json:password`
		Nickname  string   `json:nickname`
		Avatar    string   `json:avatar`
		Roles     []string `json:roles`
		Remark    string   `json:remark`
		CreatedAt string   `json:createdAt`
		CreatedBy string   `json:createdBy`
		UpdatedAt string   `json:updatedAt`
		UpdatedBy string   `json:updatedBy`
	}
)

