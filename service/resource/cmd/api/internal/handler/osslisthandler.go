package handler

import (
	"go-zero-resource/common/api"
	"net/http"

	"go-zero-resource/service/resource/cmd/api/internal/logic"
	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ossListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OssListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewOssListLogic(r.Context(), ctx)
		resp, err := l.OssList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, api.OkWithData(resp))
		}
	}
}
