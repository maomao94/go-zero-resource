package wsx

import "encoding/json"

type WsRequest struct {
	Seq  string         `json:"seq"`
	Cmd  string         `json:"cmd"`
	Data map[string]any `json:"data,omitempty"`
}

func (h *WsRequest) String() (str string) {
	bytes, _ := json.Marshal(h)
	str = string(bytes)
	return
}

type WsResponse struct {
	Seq      string    `json:"seq"`
	Cmd      string    `json:"cmd"`
	Response *Response `json:"response"`
}

func (h *WsResponse) String() (str string) {
	bytes, _ := json.Marshal(h)
	str = string(bytes)
	return
}

type Response struct {
	Code uint32         `json:"code"`
	Msg  string         `json:"msg"`
	Data map[string]any `json:"data"`
}

type LoginReq struct {
	token  string `json:"token"`
	AppId  uint32 `json:"appId,omitempty"`
	UserId string `json:"userId,omitempty"`
}

type HeartBeatReq struct {
	UserId string `json:"userId,omitempty"`
}
