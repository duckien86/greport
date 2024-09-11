package restaurantlikestorage

import (
	"2ndbrand-api/common"
	restaurantmodel "2ndbrand-api/module/restaurant/model"
	restaurantlikesmodel "2ndbrand-api/module/restaurantlike/model"
	"context"
)

func (s *sqlStore) FindDataWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantlikesmodel.Like, error) {
	var returnData restaurantlikesmodel.Like
	if err := s.db.Where(condition).First(&returnData).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &returnData, nil
}

func (s *sqlStore) FindRestaurantWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurants, error) {
	var returnData restaurantmodel.Restaurants
	if err := s.db.Where(condition).First(&returnData).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &returnData, nil
}
