package reportstorage

import (
	"context"
	"greport/common"
	reportmodel "greport/module/report/model"
)

func (s *sqlStore) FindAll(context context.Context, condition reportmodel.MsgLogReq, moreKeys ...string) (*[]reportmodel.MsgLogRes, error) {
	var returnData *[]reportmodel.MsgLogRes

	if err := s.db.Table("msg_log_agg_v4").Limit(20).
		Select("message_id", "campaign_id", "channel", "template_id").
		Where("channel=?", "zalo").
		Find(&returnData).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return returnData, nil
}
