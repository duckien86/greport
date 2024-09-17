package reportbiz

import (
	"context"
	reportmodel "greport/module/report/model"
)

type ReportStorageInterface interface {
	FindAll(context context.Context, condition reportmodel.MsgLogRequest, moreKeys ...string) (*[]reportmodel.MsgLogResponse, error)
}

type reportBiz struct {
	store ReportStorageInterface
}

func NewReportBiz(store ReportStorageInterface) *reportBiz {
	return &reportBiz{
		store: store,
	}
}

// GetMsgLog: Do msg log biz
func (rb *reportBiz) GetMsgLog(ctx context.Context, reqData reportmodel.MsgLogRequest) (*[]reportmodel.MsgLogResponse, error) {
	data, err := rb.store.FindAll(ctx, reqData)
	if err != nil {
		return nil, err
	}

	return data, nil
}
