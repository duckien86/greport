package restaurantbiz

import (
	"2ndbrand-api/common"
	restaurantmodel "2ndbrand-api/module/restaurant/model"
	"context"
)

type CreateRestaurantStore interface {
	Create(context context.Context, data *restaurantmodel.RestaurantsCreate) error
}
type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(
	ctx context.Context,
	data *restaurantmodel.RestaurantsCreate,
) error {
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}
	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
	}
	return nil
}
