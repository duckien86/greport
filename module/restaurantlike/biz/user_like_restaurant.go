package restaurantlikebiz

import (
	"2ndbrand-api/common"
	"2ndbrand-api/component/asyncjob"
	restaurantmodel "2ndbrand-api/module/restaurant/model"
	restaurantlikesmodel "2ndbrand-api/module/restaurantlike/model"
	"context"
)

type CreateRestaurantLikeStore interface {
	Create(context context.Context, data *restaurantlikesmodel.Like) error
	FindRestaurantWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurants, error)
}
type createRestaurantLikeBiz struct {
	store             CreateRestaurantLikeStore
	storeIncLikeCount IncLikeCountStore
}

type IncLikeCountStore interface {
	IncLikeCount(ctx context.Context, updateData *restaurantmodel.Restaurants) error
}

func NewLikeRestaurantBiz(store CreateRestaurantLikeStore, storeIncLikeCount IncLikeCountStore) *createRestaurantLikeBiz {
	return &createRestaurantLikeBiz{
		store:             store,
		storeIncLikeCount: storeIncLikeCount,
	}
}
func (biz *createRestaurantLikeBiz) LikeRestaurant(
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
	if err := biz.store.Create(ctx, data); err != nil {
		return restaurantlikesmodel.ErrCannotLikeRestaurant(err)
	}
	// side effect : increase like (async job)
	// go func() {
	job := asyncjob.NewJob(func(ctx context.Context) error {
		// go func() {
		return biz.storeIncLikeCount.IncLikeCount(ctx, restaurant)
		// }()
		// return nil
	})
	asyncjob.NewGroup(true, job).Run(ctx)
	// }()

	return nil
}
