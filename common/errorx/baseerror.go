package errorx

const defaultCode = 500

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

func NewCodeError(code int, msg string, data interface{}) error {
	return &CodeError{Code: code, Msg: msg, Data: data}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg, map[string]interface{}{})
}

func NewErrorMessage(code int, msg string) error {
	return NewCodeError(code, msg, map[string]interface{}{})
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Fail() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
		Data: e.Data,
	}
}