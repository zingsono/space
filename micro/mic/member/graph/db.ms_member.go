// 会员信息数据库操作
package graph

import (
	"time"
)

// 状态（1=正常 2=禁用 9=注销）
type MemberStatus string

const (
	MEMBER_STATUS_NORMAL    MemberStatus = "1"
	MEMBER_STATUS_FORBIDDEN MemberStatus = "2"
	MEMBER_STATUS_INVALID   MemberStatus = "9"
)

// 账号类型（1=个人 2=机构）
type MemberType string

const (
	MEMBER_TYPE_USER MemberType = "1"
	MEMBER_TYPE_ORG  MemberType = "2"
)

// 实名认证(0=否 1=个人认证 2=企业认证 )
type MemberReal string

const (
	MEMBER_REAL_Y MemberType = "1"
	MEMBER_REAL_N MemberType = "0"
)

// 会员对象
type Member struct {
	Uid           string       `bson:"uid" json:"uid"`                     // 会员编号，应用于所有服务
	LoginId       string       `bson:"loginId" json:"loginId"`             // 登录账号
	Password      string       `bson:"password" json:"password"`           // 密码，保存md5值
	Mobile        string       `bson:"mobile" json:"mobile"`               // 用户绑定手机，可用于登录
	Email         string       `bson:"email" json:"email"`                 // 用户绑定邮箱,可用于登录
	Nickname      string       `bson:"nickname" json:"nickname"`           // 用户昵称
	Avatar        string       `bson:"avatar" json:"avatar"`               // 用户头像URL
	CreatedAt     time.Time    `bson:"createdAt" json:"createdAt"`         // 注册时间
	UpdatedAt     time.Time    `bson:"updatedAt" json:"updatedAt"`         // 最后更新时间
	LastLoginTime time.Time    `bson:"lastLoginTime" json:"lastLoginTime"` // 最后登录时间
	Status        MemberStatus `bson:"status" json:"status"`               // 用户状态
	Type          MemberType   `bson:"type" json:"type"`                   // 用户类型
	Real          MemberReal   `bson:"real" json:"real"`                   // 是否实名认证
	User          *MemberUser  `bson:"user" json:"user"`                   // 个人用户扩展信息
	Org           *MemberOrg   `bson:"org" json:"org"`                     // 企业机构用户扩展信息
}

// 证件类型（1=身份证）
type UserIcType string

const (
	USER_IC_TYPE_SFZ UserIcType = "1"
)

// 扩展字段个人会员信息(user)
// 真实姓名、证件类型、证件号码、证件照正面、证件照反面、联系地址
type MemberUser struct {
	RealName   string     `bson:"realName" json:"realName"`     // 真实姓名
	Province   string     `bson:"province" json:"province"`     // 省
	City       string     `bson:"city" json:"city"`             // 市
	Address    string     `bson:"address" json:"address"`       // 联系地址
	IcType     UserIcType `bson:"icType" json:"icType"`         // 证件类型
	IcNumber   string     `bson:"icNumber" json:icNumber`       // 证件号码
	IcAddress  string     `bson:"icAddress" json:"icAddress"`   // 证件地址
	IcBegTime  time.Time  `bson:"icBegTime" json:"icBegTime"`   // 有效期开始
	IcEndTime  time.Time  `bson:"icEndTime" json:"icEndTime"`   // 有效期结束
	IcImgFront string     `bson:"icImgFront" json:"icImgFront"` // 证件正面图片
	IcImgBack  string     `bson:"icImgBack" json:"icImgBack"`   // 证件反面图片
}

// 扩展字段机构会员信息(org)
// 证件号码、证件照片、机构名称、负责人姓名、负责人电话、联系地址
type MemberOrg struct {
	BelongUid       string    `bson:"belongUid" json:"belongUid"`             // 所属管理会员编号，用于商户有门店或者子账号的情况
	BrandName       string    `bson:"BrandName" json:"BrandName"`             // 品牌名称
	LogoUrl         string    `bson:"logoUrl" json:"logoUrl"`                 // 品牌LOGO
	LeadName        string    `bson:"leadName" json:"leadName"`               // 负责人姓名
	LeadTel         string    `bson:"leadTel" json:"leadTel"`                 // 负责人联系电话
	LeadPost        string    `bson:"leadPost" json:"leadPost"`               // 负责人岗位
	Address         string    `bson:"address" json:"address"`                 // 经营地址
	ServiceTel      string    `bson:"serviceTel" json:"serviceTel"`           // 客服电话
	BlName          string    `bson:"blName" json:"blName"`                   // 证件资料
	BlNumber        string    `bson:"blNumber" json:"blNumber"`               // 证件号
	BlImg           string    `bson:"blImg" json:"blImg"`                     // 证件图片
	BlBegTime       time.Time `bson:"blBegTime" json:"blBegTime"`             // 证件有效期开始
	BlEndTime       time.Time `bson:"blEndTime" json:"blEndTime"`             // 证件有效期结束
	BlLegalName     string    `bson:"blLegalName" json:"blLegalName"`         // 法人名
	BlAddress       string    `bson:"blAddress" json:"blAddress"`             // 证件地址
	BlBusinessScope string    `bson:"blBusinessScope" json:"blBusinessScope"` // 营业范围
}
