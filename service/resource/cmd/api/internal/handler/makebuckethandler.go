package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-resource/service/resource/cmd/api/internal/logic"
	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"
)

func makeBucketHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MakeBucketReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMakeBucketLogic(r.Context(), ctx)
		resp, err := l.MakeBucket(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
