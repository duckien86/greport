package userbiz

import (
	"2ndbrand-api/common"
	restaurantmodel "2ndbrand-api/module/restaurant/model"
	"context"
)

type DeleteRestaurantStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurants, error)
	Delete(context context.Context, id int) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil { // nếu không tìm thấy data
		return common.ErrInvalidRequest(err)
	}

	if oldData.Status == 0 { // nếu dữ liệu đã soft delete
		return common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, err)
	}
	return nil
}
