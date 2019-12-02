package graph

/* 全局接口错误码定义 */

type ResponseCode struct {
	Code string `bson:"code" json:"code"`
	Msg  string `bson:"msg" json:"msg"`
}

var (
	// 通用错误码，代表成功与系统错误
	ERR_SUCCESS = &ResponseCode{Code: "00000", Msg: "ok"}
	ERR_FAIL    = &ResponseCode{Code: "99999", Msg: "fail"}
)
