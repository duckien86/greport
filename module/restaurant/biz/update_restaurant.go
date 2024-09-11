package restaurantbiz

import (
	"2ndbrand-api/common"
	restaurantmodel "2ndbrand-api/module/restaurant/model"
	"context"
	"errors"
)

type UpdateRestaurantStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurants, error)
	Update(
		context context.Context,
		updateData *restaurantmodel.RestaurantsUpdate,
		id int,
	) error
}

type updateRestaurantBiz struct {
	store     UpdateRestaurantStore
	requester common.Requester
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore, requester common.Requester) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store, requester: requester}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(
	ctx context.Context,
	updateData *restaurantmodel.RestaurantsUpdate,
	id int,
) error {
	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if oldData.OwnerId != biz.requester.GetUserID() {
		return common.ErrNoPermision(nil)
	}
	if oldData.Status == 0 {
		return errors.New("data has been deleted")
	}
	if err := biz.store.Update(ctx, updateData, id); err != nil {
		return err
	}
	return nil
}
