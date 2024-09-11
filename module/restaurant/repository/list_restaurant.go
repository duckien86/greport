package restaurantrepo

import (
	"2ndbrand-api/common"
	restaurantmodel "2ndbrand-api/module/restaurant/model"
	"context"
)

type ListRestaurantStore interface {
	ListDataWithCondition(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurants, error)
}
type LikeRestaurantStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}
type listRestaurantRepo struct {
	store     ListRestaurantStore
	storeLike LikeRestaurantStore
}

func NewListRestaurantRepo(store ListRestaurantStore, storeLike LikeRestaurantStore) *listRestaurantRepo {
	return &listRestaurantRepo{
		store:     store,
		storeLike: storeLike,
	}
}

func (biz *listRestaurantRepo) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurants, error) {
	var listRestaurant []restaurantmodel.Restaurants
	listRestaurant, err := biz.store.ListDataWithCondition(ctx, filter, paging, "User")
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}
	restaurantIds := make([]int, len(listRestaurant))
	for k, v := range listRestaurant {
		restaurantIds[k] = v.Id
	}
	restaurantLikes, _ := biz.storeLike.GetRestaurantLikes(ctx, restaurantIds)
	for k, v := range listRestaurant {
		listRestaurant[k].LikeCount = restaurantLikes[v.Id]
	}
	return listRestaurant, nil
}
