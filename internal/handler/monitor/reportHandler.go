package monitor

import (
	"net/http"

	"exporter/internal/logic/monitor"
	"exporter/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ReportHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := monitor.NewReportLogic(r.Context(), svcCtx)
		resp, err := l.Report()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
