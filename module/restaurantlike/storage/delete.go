package restaurantlikestorage

import (
	"2ndbrand-api/common"
	restaurantlikesmodel "2ndbrand-api/module/restaurantlike/model"
	"context"
)

func (s *sqlStore) Delete(context context.Context, condition map[string]interface{}) error {
	if err := s.db.Table(restaurantlikesmodel.Like{}.TableName()).
		Where(condition).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
