package gtw

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gtw/gtw/internal/logic/gtw"
	"gtw/gtw/internal/svc"
	"gtw/gtw/internal/types"
)

func MakeBucketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MakeBucketReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := gtw.NewMakeBucketLogic(r.Context(), svcCtx)
		resp, err := l.MakeBucket(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
