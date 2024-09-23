package email

import (
	"net/http"

	"exporter/internal/logic/email"
	"exporter/internal/svc"
	"exporter/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EasyemailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EasyEmailRes
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := email.NewEasyemailLogic(r.Context(), svcCtx)
		resp, err := l.Easyemail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
