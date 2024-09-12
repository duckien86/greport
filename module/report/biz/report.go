package reportbiz

import (
	"context"
	reportmodel "greport/module/report/model"
)

type ReportStorage interface {
	FindAll(context context.Context, condition reportmodel.MsgLogReq, moreKeys ...string) (*[]reportmodel.MsgLogRes, error)
}

type reportBiz struct {
	store ReportStorage
}

func NewReportBiz(store ReportStorage) *reportBiz {
	return &reportBiz{
		store: store,
	}
}

// func (biz *loginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {

func (r *reportBiz) GetMsgLog(ctx context.Context, reqData reportmodel.MsgLogReq) (*[]reportmodel.MsgLogRes, error) {
	data, err := r.store.FindAll(ctx, reqData)
	if err != nil {
		return nil, err
	}

	return data, nil
}
