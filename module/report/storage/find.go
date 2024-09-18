package reportstorage

import (
	"context"
	"fmt"
	"greport/common"
	reportmodel "greport/module/report/model"
	"log"

	sq "github.com/Masterminds/squirrel"
)

func (s *sqlStore) FindAll(context context.Context, filter *reportmodel.MsgLogFilter, paging *common.Paging, moreKeys ...string) (*[]reportmodel.MsgLogResponse, error) {
	var returnData []reportmodel.MsgLogResponse
	qrDetails := sq.
		Select(
			"message_id  AS MessageId",
			"campaign_id AS  CampaignId",
			"channel AS  Channel",
			"template_id  AS TemplateId",
			"time_sent AS  TimeSent",
		).From("msg_log_agg_v4")

	queryTotal := sq.Select("COUNT(*) AS total").From("msg_log_agg_v4")

	if filter.Channel != "" {
		// @TODO: Phương án tạm thời này sẽ bị sql injection
		qrDetails = qrDetails.Where(fmt.Sprintf("channel = '%s' ", filter.Channel))
		queryTotal = queryTotal.Where(fmt.Sprintf("channel = '%s' ", filter.Channel))
	}
	qrDetails = qrDetails.Offset(uint64(paging.GetOffset()))
	qrDetails = qrDetails.Limit(uint64(paging.Limit))
	sqlStr, params := qrDetails.MustSql()
	log.Println(sqlStr, params)
	// Count total
	sqlStrTotal, _ := queryTotal.MustSql()
	if err := s.db.QueryRow(context, sqlStrTotal).Scan(&paging.Total); err != nil {
		return nil, fmt.Errorf("[Query total] -> %w", err)
	}

	// Get details
	rows, err := s.db.Query(context, sqlStr)
	if err != nil {
		return nil, err
	}
	var row reportmodel.MsgLogResponse
	for rows.Next() {
		if err := rows.ScanStruct(&row); err != nil {
			return nil, err
		}
		returnData = append(returnData, row)
	}
	rows.Close()
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
