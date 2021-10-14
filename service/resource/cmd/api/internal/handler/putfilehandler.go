package handler

import (
	"go-zero-resource/common/api"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go-zero-resource/service/resource/cmd/api/internal/logic"
	"go-zero-resource/service/resource/cmd/api/internal/svc"
)

func putFileHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewPutFileLogic(r.Context(), ctx)
		err := l.PutFile()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, api.Ok())
		}
	}
}
