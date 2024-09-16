package reportstorage

import (
	"context"
	reportmodel "greport/module/report/model"
)

func (s *sqlStore) FindAll(context context.Context, condition reportmodel.MsgLogReq, moreKeys ...string) (*[]reportmodel.MsgLogRes, error) {
	var returnData []reportmodel.MsgLogRes
	sqlCmd := "select message_id MessageId, campaign_id CampaignId, channel Channel, template_id TemplateId, time_sent TimeSent "
	sqlCmd += "from msg_log_agg_v4 "
	sqlCmd += "where channel=?"
	sqlCmd += "limit 100"
	rows, err := s.db.Query(context, sqlCmd, "sms")
	if err != nil {
		return nil, err
	}
	var row reportmodel.MsgLogRes

	for rows.Next() {
		if err := rows.ScanStruct(&row); err != nil {
			return nil, err
		}
		returnData = append(returnData, row)
	}
	rows.Close()
	// defer s.db.Close()
	return &returnData, nil
}

// func (s *sqlStore) FindAll(context context.Context, condition reportmodel.MsgLogReq, moreKeys ...string) (*[]reportmodel.MsgLogRes, error) {
// 	var returnData *[]reportmodel.MsgLogRes

// 	if err := s.db.Table("msg_log_agg_v4").Limit(20).
// 		Select("message_id", "campaign_id", "channel", "template_id").
// 		Where("channel=?", "zalo").
// 		Find(&returnData).
// 		Error; err != nil {
// 		return nil, common.ErrDB(err)
// 	}
// 	return returnData, nil
// }
