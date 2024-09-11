package restaurantstorage

import (
	"2ndbrand-api/common"
	restaurantmodel "2ndbrand-api/module/restaurant/model"
	"context"
)

func (s *sqlStore) ListDataWithCondition(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurants, error) {
	var result []restaurantmodel.Restaurants
	db := s.db.Table(restaurantmodel.Restaurants{}.TableName()).Where("status in (?)", restaurantmodel.StatusActive)

	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("owner_id = ?", f.OwnerId)
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)
		if err != nil {
			return nil, common.ErrDB(err)
		}
		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}
	// load relation entiry
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}
	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}
	if len(result) > 0 {
		lastRow := result[len(result)-1]
		lastRow.Mask(false)
		paging.NextCursor = lastRow.FakeId.String()
	}
	return result, nil
}
