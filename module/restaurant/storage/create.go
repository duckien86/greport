package restaurantstorage

import (
	"2ndbrand-api/common"
	restaurantmodel "2ndbrand-api/module/restaurant/model"
	"context"
)

func (s *sqlStore) Create(context context.Context, data *restaurantmodel.RestaurantsCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
