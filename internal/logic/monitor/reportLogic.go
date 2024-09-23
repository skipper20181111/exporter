package monitor

import (
	"context"
	"exporter/internal/logic/util"

	"exporter/internal/svc"
	"exporter/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	ml     *util.MountLogic
}

func NewReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReportLogic {
	return &ReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		ml:     util.NewMountLogic(ctx, svcCtx),
	}
}

func (l *ReportLogic) Report() (resp *types.EncryptResp, err error) {
	l.ml.MountReport()
	return
}
