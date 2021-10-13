package handler

import (
	"go-zero-resource/common/common/api"
	"net/http"

	"go-zero-resource/service/oss-endpoint/cmd/api/internal/logic"
	"go-zero-resource/service/oss-endpoint/cmd/api/internal/svc"
	"go-zero-resource/service/oss-endpoint/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
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
		api.Response(w, resp, err)
	}
}
