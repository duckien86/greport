package restaurantlikesmodel

type Filter struct {
	UserID       int `json:"-" form:"user_id"`
	RestaurantID int `json:"-" form:"restaurant_id"`
}
