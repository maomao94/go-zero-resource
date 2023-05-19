package wss

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/hehanpeng/go-zero-resource/mgtw/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func WsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 升级协议
		conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
			logx.WithContext(r.Context()).Infof("升级协议", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])
			return true
		}}).Upgrade(w, r, nil)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		fmt.Println("webSocket 建立连接:", conn.RemoteAddr().String())
	}
}
