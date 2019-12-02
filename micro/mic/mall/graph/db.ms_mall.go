package graph

/**
业务说明：
1. 用户ID就是商城ID,每个用户可以申请开一个商城；

*/

// 商城站点信息集合 (ms_mall)
// 字段： 创建商城的用户、商城名称、商城备注描述、商城客服电话、客服邮箱、商城域名
type Mall struct {
	Uid         string `bson:"uid" json:"uid"`                 // 用户编号，商城所属用户
	Name        string `bson:"name" json:"name"`               // 商城名称
	Title       string `bson:"title" json:"title"`             // 商城标题，用于网页Title
	Description string `bson:"description" json:"description"` // 商城描述
	Remark      string `bson:"remark" json:"remark"`           // 备注信息
	Ico         string `bson:"ico" json:"ico"`                 // 备注信息
	LogoUrl     string `bson:"logoUrl" json:"logoUrl"`         // 商城LOGO图
	Domain      string `bson:"domain" json:"domain"`           // 商城访问域名
	Ad          []*Ad  `bson:"ad" json:"ad"`                   // 商城广告信息管理，如Banner、侧边广告
}

// 广告是否展示
type AdIsShow string

const (
	AD_SHOW AdIsShow = "1"
	AD_HIDE AdIsShow = "0"
)

// 广告
// key值为广告位名称，如：banner
type Ad struct {
	Id     string `bson:"id" json:"id"`
	Key    string `bson:"key" json:"key"`
	ImgUrl string `bson:"imgUrl" json:"imgUrl"`
	Href   string `bson:"href" json:"href"`
	Title  string `bson:"title" json:"title"`
	Remark string `bson:"remark" json:"remark"`

	// 是否展示(1=展示 0=隐藏)
	IsShow AdIsShow `bson:"show" json:"show"`
	// 正序排序
	Rank int
}
