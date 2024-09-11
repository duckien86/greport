package restaurantlikebiz

import (
	"2ndbrand-api/common"
	restaurantlikesmodel "2ndbrand-api/module/restaurantlike/model"
	usermodel "2ndbrand-api/module/user/model"
	"context"
)

type ListUserLikedStore interface {
	GetUserLikeRestaurant(
		ctx context.Context,
		condition map[string]interface{},
		filter *restaurantlikesmodel.Filter,
		paging *common.Paging,
	) ([]common.UserPublic, error)
}

type listUserLikedBiz struct {
	store ListUserLikedStore
}

func NewListUserLikedBiz(store ListUserLikedStore) *listUserLikedBiz {
	return &listUserLikedBiz{store: store}
}

func (biz *listUserLikedBiz) ListUserLiked(
	ctx context.Context,
	filter *restaurantlikesmodel.Filter,
	paging *common.Paging,
) ([]common.UserPublic, error) {
	result, err := biz.store.GetUserLikeRestaurant(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(usermodel.EntityName, err)
	}
	return result, nil
}
