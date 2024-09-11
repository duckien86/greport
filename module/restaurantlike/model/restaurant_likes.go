package restaurantlikesmodel

import (
	"2ndbrand-api/common"
	"errors"
	"time"
)

const EntityName = "UserLikeRestaurant"

type Like struct {
	RestaurantID int                `json:"restaurant_id" gorm:"column:restaurant_id"`
	UserID       int                `json:"user_id" gorm:"column:user_id"`
	CreatedAt    *time.Time         `json:"-" gorm:"column:created_at"`
	User         *common.UserPublic `json:"user" gorm:"preload:false"`
}

func (Like) TableName() string {
	return "restaurant_likes"
}

func ErrCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"cannot like this restaurant",
		"ErrCannotLikeRestaurant",
	)
}
func ErrCannotUnlikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"cannot unlike this restaurant",
		"ErrCannotUnlikeRestaurant",
	)
}

var (
	ErrIsEmpty = errors.New("user_id and restaurant_id cannot be empty")
)
