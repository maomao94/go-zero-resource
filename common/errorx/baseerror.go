package errorx

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	DefaultCode  = -20000 // 服务不可用
	AuthCode     = -20001 // 授权权限不足
	MissCode     = -40001 // 缺少必选参数
	InvalidCode  = -40002 // 非法参数
	BizCode      = -40004 // 业务处理失败
	NotFoundCode = -40005 // 记录不存在

	ErrorCodeMsg = map[int]string{
		DefaultCode:  "服务不可用",
		AuthCode:     "授权权限不足",
		MissCode:     "缺少必选参数",
		InvalidCode:  "非法参数",
		BizCode:      "业务处理失败",
		NotFoundCode: "记录不存在",
	}
	//ErrCode_value = map[string]int{
	//	"服务不可用":  20000,
	//	"授权权限不足": 20001,
	//	"缺少必选参数": 40001,
	//	"非法参数":   40002,
	//	"业务处理失败": 40004,
	//}
)

type CodeError struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type CodeErrorResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ParseError(err error) error {
	logx.ErrorStackf("error: %s;", err.Error())
	switch err.(type) {
	case *CodeError:
		return err
	default:
		return NewDefaultError(fmt.Sprintf("error: %s;", err))
	}
}

func NewDefaultError(msg string) error {
	return NewCodeErrorWithData(DefaultCode, msg, map[string]interface{}{})
}

func NewCodeError(code int) error {
	return NewCodeErrorWithData(code, "-", map[string]interface{}{})
}

func NewCodeMsgError(code int, msg string) error {
	return NewCodeErrorWithData(code, msg, map[string]interface{}{})
}

func NewCodeErrorWithData(code int, msg string, data interface{}) error {
	return &CodeError{Code: code, Msg: msg, Data: data}
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) getCode() int {
	if e.Code == 0 {
		e.Code = DefaultCode
	}
	return e.Code
}

func (e *CodeError) getMsg() string {
	if e.Msg == "-" {
		e.Msg = ErrorCodeMsg[e.Code]
	}
	return e.Msg
}

func (e *CodeError) Fail() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.getCode(),
		Msg:  e.getMsg(),
		Data: e.Data,
	}
}
