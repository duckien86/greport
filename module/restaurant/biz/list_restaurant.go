package restaurantbiz

import (
	"2ndbrand-api/common"
	restaurantmodel "2ndbrand-api/module/restaurant/model"
	"context"
)

type ListRestaurantRepo interface {
	ListRestaurant(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurants, error)
}
type listRestaurantRepo struct {
	repo ListRestaurantRepo
}

func NewListRestaurantBiz(repo ListRestaurantRepo) *listRestaurantRepo {
	return &listRestaurantRepo{
		repo: repo,
	}
}

func (biz *listRestaurantRepo) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurants, error) {
	listRestaurant, err := biz.repo.ListRestaurant(ctx, filter, paging, "User")
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}
	return listRestaurant, nil
}
