package serializer

import (
	"singo/serializer/handler"
)

// 三位数错误编码为复用http原本含义
// 五位数错误编码为应用自定义错误
// 五开头的五位数错误编码为服务器端错误，比如数据库操作失败
// 四开头的五位数错误编码为客户端错误，有时候是客户端代码写错了，有时候是用户操作错误
const (
	// CodeSuccess 正常
	CodeSuccess = 200
	// CodeCheckLogin 未登录
	CodeCheckLogin = 401
	// CodeNoRightErr 未授权访问
	CodeNoRightErr = 403
	// CodeDBError 数据库操作失败
	CodeDBError = 50001
	// CodeEncryptError 加密失败
	CodeEncryptError = 50002
	//CodeParamErr 各种奇奇怪怪的参数错误
	CodeParamErr = 40001
)

var (
	ErrDatabase = handler.NewActionError(500, 100010002, "数据库错误")
	ErrParams   = handler.NewActionError(400, 501010000, "参数错误")
)

// ErrParamsMsg ErrParams with extra message
func ErrParamsMsg(msg string) error {
	return handler.NewActionError(ErrParams.Status, ErrParams.Code, msg)
}

// CheckLogin 检查登录
func CheckLogin() handler.Response {
	return handler.Response{
		Code: CodeCheckLogin,
		Msg:  "未登录",
	}
}
