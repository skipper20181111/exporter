package refresh

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"exporter/internal/svc"
	"exporter/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	return &RefreshLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshLogic) Refresh() (resp *types.RefreshResp, err error) {
	resp = &types.RefreshResp{
		Code: "10000",
		Msg:  "success",
	}
	l.EmailList()
	return resp, nil
}

func (l *RefreshLogic) EmailList() {
	emailList := make([]*types.EmailInfo, 0)
	filePtr, err := os.Open("etc/email.json")
	if err != nil {
		return
	}
	defer filePtr.Close()
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&emailList)
	if err != nil {
		fmt.Println("解码失败", err.Error())
	} else {
		fmt.Println("解码成功")
		l.svcCtx.LocalCache.Set(svc.EmailListKey, emailList)
	}
}
