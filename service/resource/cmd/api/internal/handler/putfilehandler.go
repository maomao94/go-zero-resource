package handler

import (
	"fmt"
	"go-zero-resource/common/api"
	"go-zero-resource/common/errorx"
	"net/http"

	"go-zero-resource/service/resource/cmd/api/internal/logic"
	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func putFileHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PutFileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprint(err)))
			return
		}

		l := logic.NewPutFileLogic(r.Context(), ctx)
		_, header, err := r.FormFile("file")
		if err != nil {
			httpx.Error(w, errorx.NewDefaultError(fmt.Sprint(err)))
			return
		}
		resp, err := l.PutFile(req, header)
		if err != nil {
			httpx.Error(w, errorx.ParseError(err))
		} else {
			httpx.OkJson(w, api.OkWithData(resp))
		}
	}
}
