package resource

import (
	"net/http"

	"github.com/hehanpeng/go-zero-resource/gtw/internal/logic/resource"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetFileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := resource.NewGetFileLogic(r.Context(), svcCtx)
		err := l.GetFile(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
