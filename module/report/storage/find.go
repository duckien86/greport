package reportstorage

import (
	"context"
	"fmt"
	reportmodel "greport/module/report/model"
	"time"
)

func (s *sqlStore) FindAll(context context.Context, condition reportmodel.MsgLogReq, moreKeys ...string) (*[]reportmodel.MsgLogRes, error) {
	var returnData *[]reportmodel.MsgLogRes
	rows, err := s.db.Query(context, "select price, date, postcode1, postcode2 from uk_price_paid limit 10")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			price     float64
			date      time.Time
			postcode1 string
			postcode2 string
		)
		if err := rows.Scan(&price, &date, &postcode1, &postcode2); err != nil {
			return nil, err
		}
		fmt.Printf("row: price=%f, date=%s , postcode1=%s, postcode2=%s \n", price, date, postcode1, postcode2)
	}
	rows.Close()
	// return rows.Err()
	return returnData, nil
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
