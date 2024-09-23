package monitor

import (
	"net/http"

	"exporter/internal/logic/monitor"
	"exporter/internal/svc"
	"exporter/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EncryptHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EncryptRes
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := monitor.NewEncryptLogic(r.Context(), svcCtx)
		resp, err := l.Encrypt(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
