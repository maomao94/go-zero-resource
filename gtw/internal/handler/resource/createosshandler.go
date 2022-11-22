package resource

import (
	"net/http"

	"github.com/hehanpeng/go-zero-resource/gtw/internal/logic/resource"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/svc"
	"github.com/hehanpeng/go-zero-resource/gtw/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateOssHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OssCreate
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := resource.NewCreateOssLogic(r.Context(), svcCtx)
		resp, err := l.CreateOss(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
