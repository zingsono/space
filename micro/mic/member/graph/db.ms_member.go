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

// 实名认证(1=是 0=否)
type MemberReal string

const (
	MEMBER_REAL_Y MemberType = "1"
	MEMBER_REAL_N MemberType = "0"
)

// 会员对象
type Member struct {
	MemberId      string       // 会员编号
	LoginId       string       // 登录账号
	Password      string       // 密码，保存md5值
	Mobile        string       // 用户绑定手机，可用于登录
	Email         string       // 用户绑定邮箱,可用于登录
	Nickname      string       // 用户昵称
	Avatar        string       // 用户头像URL
	Remark        string       // 备注信息，后台管理使用
	CreatedAt     time.Time    // 注册时间
	UpdatedAt     time.Time    // 最后更新时间
	LastLoginTime time.Time    // 最后登录时间
	Status        MemberStatus // 用户状态
	Type          MemberType   // 用户类型
	Real          MemberReal   // 是否实名认证
	User          *User        // 个人用户扩展信息
	Org           *Org         // 企业机构用户扩展信息
}

// 证件类型（1=身份证）
type UserIcType string

const (
	USER_IC_TYPE_SFZ UserIcType = "1"
)

// 扩展字段个人会员信息(user)
// 真实姓名、证件类型、证件号码、证件照正面、证件照反面、联系地址
type User struct {
	RealName   string     // 真实姓名
	Province   string     // 省
	City       string     // 市
	Address    string     // 联系地址
	IcType     UserIcType `bson:"icType" json:"icType"`   // 证件类型
	IcNumber   string     `bson:"icNumber" json:icNumber` // 证件号码
	IcAddress  string     // 证件地址
	IcBegTime  time.Time  // 有效期开始
	IcEndTime  time.Time  // 有效期结束
	IcImgFront string     // 证件正面图片
	IcImgBack  string     // 证件反面图片
}

// 扩展字段机构会员信息(org)
// 证件号码、证件照片、机构名称、负责人姓名、负责人电话、联系地址
type Org struct {
	BelongOrg       string    // 所属管理会员编号，用于商户有门店的情况
	BrandName       string    // 品牌名称
	LogoUrl         string    // 品牌LOGO
	LeadName        string    // 负责人姓名
	LeadTel         string    // 负责人联系电话
	LeadPost        string    // 负责人岗位
	Address         string    // 经营地址
	ServiceTel      string    // 客服电话
	BlName          string    // 证件资料
	BlNumber        string    // 证件号
	BlImg           string    // 证件图片
	BlBegTime       time.Time // 证件有效期开始
	BlEndTime       time.Time // 证件有效期结束
	BlLegalName     string    // 法人名
	BlAddress       string    // 证件地址
	BlBusinessScope string    // 营业范围
}
