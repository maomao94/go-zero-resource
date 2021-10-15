package errorx

import "fmt"

var (
	DefaultCode = 20000 // 服务不可用
	AuthCode    = 20001 // 授权权限不足
	MissingCode = 40001 // 缺少必选参数
	InvalidCode = 40002 // 非法参数
	BizCode     = 40004 // 业务处理失败
	NotFound    = 40005 // 记录不存在

	ErrorCodeMsg = map[int]string{
		20000: "服务不可用",
		20001: "授权权限不足",
		40001: "缺少必选参数",
		40002: "非法参数",
		40004: "业务处理失败",
		40005: "记录不存在",
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
	Data interface{} `json:"data"`
}

func formatError(err error) error {
	switch err.(type) {
	case nil:
		return NewDefaultError("error is nil")
	case *CodeError:
		return err
	default:
		return NewDefaultError(fmt.Sprint(err))
	}
}

func ParseError(err error) error {
	return formatError(err)
}

func NewDefaultError(msg string) error {
	return NewCodeErrorWithData(DefaultCode, msg, map[string]interface{}{})
}

func NewCodeError(code int) error {
	return NewCodeErrorWithData(code, "", map[string]interface{}{})
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
	if len(e.Msg) == 0 {
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
