package api

const (
	SUCCESS = 200
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data",omitempty`
}

func (r *Body) IsSuccess() bool {
	if r.Code == SUCCESS {
		return true
	} else {
		return false
	}
}

func Result(code int, msg string, data interface{}) *Body {
	return &Body{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func Ok() *Body {
	return Result(SUCCESS, "操作成功", map[string]interface{}{})
}

func OkWithMessage(message string) *Body {
	return Result(SUCCESS, message, map[string]interface{}{})
}

func OkWithData(data interface{}) *Body {
	return Result(SUCCESS, "操作成功", data)
}

func OkWithDetailed(data interface{}, message string) *Body {
	return Result(SUCCESS, message, data)
}

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}
