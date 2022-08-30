package handler

import (
	"fmt"
	"go-zero-resource/common/errorx"
	"net/http"

	"go-zero-resource/service/resource/cmd/api/internal/logic"
	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func getFileHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetFileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprint(err)))
			return
		}

		l := logic.NewGetFileLogic(r.Context(), ctx)
		err := l.GetFile(req, w)
		if err != nil {
			httpx.Error(w, errorx.ParseError(err))
		} else {
			//httpx.OkJson(w, api.Ok())
		}
	}
}
