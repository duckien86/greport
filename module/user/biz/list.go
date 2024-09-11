package userbiz

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
type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurants, error) {

	var result []restaurantmodel.Restaurants
	result, err := biz.store.ListDataWithCondition(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil

}
