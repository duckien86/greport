package restaurantlikestorage

import (
	"2ndbrand-api/common"
	restaurantlikesmodel "2ndbrand-api/module/restaurantlike/model"
	"context"
)

func (s *sqlStore) Create(context context.Context, data *restaurantlikesmodel.Like) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
