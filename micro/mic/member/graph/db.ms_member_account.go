package graph

// 会员账户(ms_member_account)
// 会员编号_id、服务机构会员ID、积分账户ID、现金账户ID、彩票账户ID、短信账户ID、电子券账户ID、产品账户ID、商城ID 、电子券产品账户、电子券会员账户、电子券验券账户、推荐人推荐码、当前用户推荐码

// 会员账户(ms_member_account)
// 会员与其他账户服务关联，会员对应每个发行机构存在一个账户。  在会员系统建立账户服务关联关系，可以实现多套会员系统使用账户服务。
type MemberAccount struct {
	MemberId string
	Org
}
