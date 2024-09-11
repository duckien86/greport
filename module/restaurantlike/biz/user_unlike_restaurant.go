package restaurantlikebiz

import (
	"2ndbrand-api/common"
	restaurantmodel "2ndbrand-api/module/restaurant/model"
	restaurantlikesmodel "2ndbrand-api/module/restaurantlike/model"
	"context"
	"log"
)

type DeleteRestaurantLikeStore interface {
	FindRestaurantWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurants, error)
	Delete(context context.Context, condition map[string]interface{}) error
}
type DcrLikeCountStore interface {
	DcrLikeCount(ctx context.Context, updateData *restaurantmodel.Restaurants) error
}
type unlikeRestaurantBiz struct {
	store  DeleteRestaurantLikeStore
	dcrLCS DcrLikeCountStore
}

func NewUnlikeRestaurantBiz(store DeleteRestaurantLikeStore, dcrLCS DcrLikeCountStore) *unlikeRestaurantBiz {
	return &unlikeRestaurantBiz{store: store, dcrLCS: dcrLCS}
}

func (biz *unlikeRestaurantBiz) UnlikeRestaurant(
	ctx context.Context,
	data *restaurantlikesmodel.Like,
) error {
	restaurant, _ := biz.store.FindRestaurantWithCondition(ctx, map[string]interface{}{"id": data.RestaurantID})
	if restaurant == nil {
		return common.ErrEntityNotFound(restaurantmodel.EntityName, nil)
	}
	if restaurant.Status == int(restaurantmodel.StatusDeleted) {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}
	if err := biz.store.Delete(
		ctx,
		map[string]interface{}{
			"user_id":       data.UserID,
			"restaurant_id": data.RestaurantID,
		},
	); err != nil {
		return restaurantlikesmodel.ErrCannotUnlikeRestaurant(err)
	}
	go func() {
		if err := biz.dcrLCS.DcrLikeCount(ctx, restaurant); err != nil {
			log.Printf("Err DcrLikeCount %v", err)
		}
	}()
	return nil
}
