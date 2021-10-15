package api

const (
	SUCCESS = 10000
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func (r *Response) IsSuccess() bool {
	if r.Code == SUCCESS {
		return true
	} else {
		return false
	}
}

func Result(code int, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func Ok() *Response {
	return Result(SUCCESS, "接口调用成功", map[string]interface{}{})
}

func OkWithMessage(message string) *Response {
	return Result(SUCCESS, message, map[string]interface{}{})
}

func OkWithData(data interface{}) *Response {
	return Result(SUCCESS, "接口调用成功", data)
}

func OkWithDetailed(data interface{}, message string) *Response {
	return Result(SUCCESS, message, data)
}
