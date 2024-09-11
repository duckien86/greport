package restaurantlikestorage

import (
	"2ndbrand-api/common"
	restaurantlikesmodel "2ndbrand-api/module/restaurantlike/model"
	"context"
	"time"

	"github.com/btcsuite/btcutil/base58"
)

func (s *sqlStore) GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error) {
	type queryData struct {
		RestaurantID int `gorm:"column:restaurant_id"`
		LikeCount    int `gorm:"column:like_count"`
	}
	var listLike []queryData

	if err := s.db.Table(restaurantlikesmodel.Like{}.TableName()).
		Select("restaurant_id, count(restaurant_id) like_count").
		Where("restaurant_id in (?)", ids).
		Group("restaurant_id").Find(&listLike).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	result := make(map[int]int)
	for _, item := range listLike {
		result[item.RestaurantID] = item.LikeCount
	}
	return result, nil
}

func (s *sqlStore) GetUserLikeRestaurant(
	ctx context.Context,
	condition map[string]interface{},
	filter *restaurantlikesmodel.Filter,
	paging *common.Paging,
) ([]common.UserPublic, error) {
	var result []restaurantlikesmodel.Like
	db := s.db.Table(restaurantlikesmodel.Like{}.TableName()).Where(condition)
	if filter != nil {
		if filter.RestaurantID > 0 {
			db = db.Where("restaurant_id= ? ", filter.RestaurantID)
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(time.DateTime, string(base58.Decode(v)))
		if err != nil {
			return nil, common.ErrDB(err)
		}
		db = db.Where("created_at <= ?", timeCreated.Format(time.DateTime))
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}
	db = db.Preload("User")
	if err := db.Limit(paging.Limit).Order("created_at desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	rtnUsers := make([]common.UserPublic, len(result))
	for k, v := range result {
		user := *v.User
		user.Mask(false)
		rtnUsers[k] = user
	}
	if len(result) > 0 {
		lastRow := result[len(result)-1]
		paging.NextCursor = base58.Encode([]byte(lastRow.CreatedAt.Format(time.DateTime)))
	}
	return rtnUsers, nil
}
