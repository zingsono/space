package model

import (
	"github.com/graphql-go/graphql"
)

type (
	// 操作员用户: 用户编号、登录账号、邮箱、手机、密码（MD5）、昵称、头像、角色（数组）、备注、创建时间、创建人、更新时间、更新人
	// uid,loginId,email,mobile,password,nickname,avatar,roles,remark,createdAt,createdBy,updatedAt,updatedBy
	MsUser struct {
		Uid       string   `json:"uid"`
		LoginId   string   `json:"loginId"`
		Email     string   `json:"email"`
		Mobile    string   `json:"mobile"`
		Password  string   `json:"password"`
		Nickname  string   `json:"nickname"`
		Avatar    string   `json:"avatar"`
		Roles     []string `json:"roles"`
		Remark    string   `json:"remark"`
		CreatedAt string   `json:"createdAt"`
		CreatedBy string   `json:"createdBy"`
		UpdatedAt string   `json:"updatedAt"`
		UpdatedBy string   `json:"updatedBy"`
	}
)

var UserObject = graphql.NewObject(graphql.ObjectConfig{
	Name:       "User",
	Interfaces: nil,
	Fields: graphql.Fields{
		"uid":       &graphql.Field{Name: "", Type: graphql.String, Args: nil, Resolve: func(p graphql.ResolveParams) (i interface{}, e error) { return "a", e }, DeprecationReason: "", Description: ""},
		"loginId":   &graphql.Field{Name: "", Type: graphql.String, Args: nil, Resolve: func(p graphql.ResolveParams) (i interface{}, e error) { return "a", e }, DeprecationReason: "", Description: ""},
		"email":     &graphql.Field{Name: "", Type: graphql.String, Args: nil, Resolve: func(p graphql.ResolveParams) (i interface{}, e error) { return "a", e }, DeprecationReason: "", Description: ""},
		"mobile":    &graphql.Field{Name: "", Type: graphql.String, Args: nil, Resolve: func(p graphql.ResolveParams) (i interface{}, e error) { return "a", e }, DeprecationReason: "", Description: ""},
		"nickname":  &graphql.Field{Name: "", Type: graphql.String, Args: nil, Resolve: func(p graphql.ResolveParams) (i interface{}, e error) { return "a", e }, DeprecationReason: "", Description: ""},
		"avatar":    &graphql.Field{Name: "", Type: graphql.String, Args: nil, Resolve: func(p graphql.ResolveParams) (i interface{}, e error) { return "a", e }, DeprecationReason: "", Description: ""},
		"roles":     &graphql.Field{Name: "", Type: graphql.String, Args: nil, Resolve: func(p graphql.ResolveParams) (i interface{}, e error) { return "a", e }, DeprecationReason: "", Description: ""},
		"remark":    &graphql.Field{Name: "", Type: graphql.String, Args: nil, Resolve: func(p graphql.ResolveParams) (i interface{}, e error) { return "a", e }, DeprecationReason: "", Description: ""},
		"createdAt": &graphql.Field{Name: "", Type: graphql.String, Args: nil, Resolve: func(p graphql.ResolveParams) (i interface{}, e error) { return "a", e }, DeprecationReason: "", Description: ""},
		"createdBy": &graphql.Field{Name: "", Type: graphql.String, Args: nil, Resolve: func(p graphql.ResolveParams) (i interface{}, e error) { return "a", e }, DeprecationReason: "", Description: ""},
		"updatedAt": &graphql.Field{Name: "", Type: graphql.String, Args: nil, Resolve: func(p graphql.ResolveParams) (i interface{}, e error) { return "a", e }, DeprecationReason: "", Description: ""},
		"updatedBy": &graphql.Field{Name: "", Type: graphql.String, Args: nil, Resolve: func(p graphql.ResolveParams) (i interface{}, e error) { return "a", e }, DeprecationReason: "", Description: ""},
	},
	IsTypeOf:    nil,
	Description: "",
})

var UserList = graphql.NewList(UserObject)

var UserPage = graphql.NewObject(graphql.ObjectConfig{
	Name:       "UserPage",
	Interfaces: nil,
	Fields: graphql.Fields{
		"total": &graphql.Field{
			Name:              "",
			Type:              graphql.Int,
			Args:              nil,
			Resolve:           func(p graphql.ResolveParams) (i interface{}, e error) { return i, e },
			DeprecationReason: "",
			Description:       "",
		},
		"pageSize": &graphql.Field{
			Name:              "",
			Type:              graphql.Int,
			Args:              nil,
			Resolve:           func(p graphql.ResolveParams) (i interface{}, e error) { return i, e },
			DeprecationReason: "",
			Description:       "",
		},
		"pageNum": &graphql.Field{
			Name:              "",
			Type:              graphql.Int,
			Args:              nil,
			Resolve:           func(p graphql.ResolveParams) (i interface{}, e error) { return i, e },
			DeprecationReason: "",
			Description:       "",
		},
		"data": &graphql.Field{
			Name:              "",
			Type:              UserList,
			Args:              nil,
			Resolve:           func(p graphql.ResolveParams) (i interface{}, e error) { return i, e },
			DeprecationReason: "",
			Description:       "",
		},
	},
	IsTypeOf:    nil,
	Description: "",
})

// 查询操作字段定义
var UserQueryFields = graphql.Fields{
	"User": &graphql.Field{
		Name: "user",
		Type: UserObject,
		Args: nil,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {

			user := &MsUser{Uid: "123"}
			return user, e
		},
		DeprecationReason: "",
		Description:       "",
	},
	"UserPage": &graphql.Field{
		Name: "",
		Type: UserPage,
		Args: nil,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {

			return i, e
		},
		DeprecationReason: "",
		Description:       "",
	},
}

// 更新操作字段定义
var UserMutationFields = graphql.Fields{
	"User": &graphql.Field{
		Name: "",
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name:       "UserOper",
			Interfaces: nil,
			Fields: graphql.Fields{
				"insert": &graphql.Field{
					Name: "",
					Type: graphql.Int,
					Args: graphql.FieldConfigArgument{
						"nickname": &graphql.ArgumentConfig{
							Type:         graphql.String,
							DefaultValue: nil,
							Description:  "",
						},
					},
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {

						return 1, e
					},
					DeprecationReason: "",
					Description:       "",
				},
				"update": &graphql.Field{
					Name: "",
					Type: graphql.Int,
					Args: graphql.FieldConfigArgument{
						"nickname": &graphql.ArgumentConfig{
							Type:         graphql.String,
							DefaultValue: nil,
							Description:  "",
						},
					},
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {

						return 1, e
					},
					DeprecationReason: "",
					Description:       "",
				},
				"delete": &graphql.Field{
					Name: "",
					Type: graphql.Int,
					Args: graphql.FieldConfigArgument{
						"nickname": &graphql.ArgumentConfig{
							Type:         graphql.String,
							DefaultValue: nil,
							Description:  "",
						},
					},
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {

						return 1, e
					},
					DeprecationReason: "",
					Description:       "",
				},
			},
			IsTypeOf:    nil,
			Description: "",
		}),
		Args:              nil,
		Resolve:           func(p graphql.ResolveParams) (i interface{}, e error) { return i, e },
		DeprecationReason: "",
		Description:       "",
	},
}
