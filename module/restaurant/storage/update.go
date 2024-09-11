package restaurantstorage

import (
	restaurantmodel "2ndbrand-api/module/restaurant/model"
	"context"

	"gorm.io/gorm"
)

func (store *sqlStore) Update(
	ctx context.Context,
	updateData *restaurantmodel.RestaurantsUpdate,
	id int,
) error {
	if err := store.db.Where("id=?", id).Updates(&updateData); err != nil {
		return err.Error
	}
	return nil
}
func (store *sqlStore) IncLikeCount(
	ctx context.Context,
	updateData *restaurantmodel.Restaurants,
	// id int,
) error {
	if err := store.db.
		Model(&updateData).
		// Table(restaurantmodel.Restaurants{}.TableName()).
		// Where("id = ?", up).
		Update("like_count", gorm.Expr("like_count + ?", 1)); err != nil {
		return err.Error
	}
	return nil
}

func (store *sqlStore) DcrLikeCount(
	ctx context.Context,
	updateData *restaurantmodel.Restaurants,
) error {
	if err := store.db.
		// Table(restaurantmodel.Restaurants{}.TableName()).
		Model(&updateData).
		Where("like_count > ?", 0).
		Update("like_count", gorm.Expr("like_count - ?", 1)); err != nil {
		return err.Error
	}
	return nil
}
