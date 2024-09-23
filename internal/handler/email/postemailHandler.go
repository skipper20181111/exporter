package email

import (
	"net/http"

	"exporter/internal/logic/email"
	"exporter/internal/svc"
	"exporter/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PostemailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostEmailRes
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := email.NewPostemailLogic(r.Context(), svcCtx)
		resp, err := l.Postemail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
