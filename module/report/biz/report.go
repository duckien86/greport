package reportbiz

import (
	"context"
	"greport/common"
	reportmodel "greport/module/report/model"
)

type ReportStorageInterface interface {
	FindAll(context context.Context,
		filter *reportmodel.MsgLogFilter,
		paging *common.Paging,
		moreKeys ...string) (*[]reportmodel.MsgLogResponse, error)
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
func (rb *reportBiz) GetMsgLog(ctx context.Context, filter *reportmodel.MsgLogFilter, paging *common.Paging) (*[]reportmodel.MsgLogResponse, error) {
	data, err := rb.store.FindAll(ctx, filter, paging)
	if err != nil {
		return nil, err
	}

	return data, nil
}
