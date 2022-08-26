package gtw

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gtw/gtw/internal/logic/gtw"
	"gtw/gtw/internal/svc"
	"gtw/gtw/internal/types"
)

func PutFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PutFileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := gtw.NewPutFileLogic(r.Context(), svcCtx)
		_, fileHeader, err := r.FormFile("file")
		if err != nil {
			httpx.Error(w, err)
		}
		resp, err := l.PutFile(&req, fileHeader)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
