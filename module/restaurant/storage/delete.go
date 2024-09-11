package restaurantstorage

import (
	"2ndbrand-api/common"
	restaurantmodel "2ndbrand-api/module/restaurant/model"
	"context"
)

func (store *sqlStore) Delete(
	context context.Context,
	id int,
) error {
	if err := store.db.Table(restaurantmodel.Restaurants{}.TableName()).
		Where("id=?", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
