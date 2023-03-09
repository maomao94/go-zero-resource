package gtw

import (
	"net/http"

	"github.com/hehanpeng/go-zero-resource/gtw/internal/logic/gtw"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PingHelloHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := gtw.NewPingHelloLogic(r.Context(), svcCtx)
		resp, err := l.PingHello()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
