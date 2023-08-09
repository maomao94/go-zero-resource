package ws

import (
	"github.com/gorilla/websocket"
	"github.com/hehanpeng/go-zero-resource/mgtw/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"time"
)

func WsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 升级协议
		conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
			logx.WithContext(r.Context()).Info("升级协议", "ua: ", r.Header["User-Agent"], "referer: ", r.Header["Referer"])
			return true
		}}).Upgrade(w, r, nil)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		logx.WithContext(r.Context()).Infof("websocket 建立连接: %s", conn.RemoteAddr().String())
		currentTime := uint64(time.Now().Unix())
		client := svc.NewClientCtx(svcCtx, conn.RemoteAddr().String(), conn, currentTime)
		threading.GoSafe(func() {
			client.Read()
		})
		threading.GoSafe(func() {
			client.Write()
		})
		svcCtx.ClientManager.PublishRegister(client)
	}
}
