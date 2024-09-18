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
	// Build where clause
	whereClause := sq.And{
		sq.Eq{"1": 1},
	}
	if filter.Channel != "" {
		whereClause = append(whereClause, sq.Eq{"channel": filter.Channel})
	}
	// if filter.KeyWord != "" {
	// 	whereClause = append(whereClause, sq.Eq{"key_word": filter.KeyWord})
	// }
	if filter.CampaignId != "" {
		whereClause = append(whereClause, sq.Eq{"campaign_id": filter.CampaignId})
	}
	if filter.TemplateId != "" {
		whereClause = append(whereClause, sq.Eq{"template_id": filter.TemplateId})
	}
	if filter.Sender != "" {
		whereClause = append(whereClause, sq.Eq{"sent_by": filter.Sender})
	}
	if filter.Recipient != "" {
		whereClause = append(whereClause, sq.Eq{"recipient": filter.Recipient})
	}
	if filter.MsgId != "" {
		whereClause = append(whereClause, sq.Eq{"message_id": filter.MsgId})
	}
	if len(filter.Telcos) > 0 {
		whereClause = append(whereClause, sq.Eq{"telco": filter.Telcos})
	}
	if len(filter.PoIds) > 0 {
		whereClause = append(whereClause, sq.Eq{"po_id": filter.PoIds})
	}

	// Get details
	qrDetails := sq.
		Select(
			"message_id  AS MessageId",
			"campaign_id AS  CampaignId",
			"channel AS  Channel",
			"template_id  AS TemplateId",
			"telco AS  Telco",
			"time_sent AS  TimeSent",
			"price AS  Price",
			"client_campaign_id AS  ClientCampaignId",
			"msg_body AS  MsgBody",
		).
		From("msg_log_agg_v4").
		Where(whereClause).
		Offset(uint64(paging.GetOffset())).
		Limit(uint64(paging.Limit))
	sqlStr, params := qrDetails.MustSql()
	log.Println(sqlStr, params)
	rows, err := s.db.Query(context, sqlStr, params...)
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

	// Count total
	qrTotal := sq.
		Select("COUNT(*) AS total").
		From("msg_log_agg_v4").
		Where(whereClause)
	sqlStrTotal, params := qrTotal.MustSql()
	log.Println(sqlStrTotal, params)
	if err := s.db.QueryRow(context, sqlStrTotal, params...).Scan(&paging.Total); err != nil {
		return nil, fmt.Errorf("[Query total] -> %w", err)
	}

	return &returnData, nil
}
