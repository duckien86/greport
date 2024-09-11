package restaurantmodel

import (
	"2ndbrand-api/common"
	"errors"
	"strings"
)

type RestaurantsType int

const StatusActive RestaurantsType = 1
const StatusDeleted RestaurantsType = 0
const EntityName = "Restaurant"

type Restaurants struct {
	common.SQLModel `json:",inline"`
	Name            string             `json:"name" gorm:"column:name;"`
	Addr            string             `json:"addr" gorm:"column:addr;"`
	OwnerId         int                `json:"-" gorm:"column:owner_id"`
	User            *common.UserPublic `json:"owner" gorm:"foreignKey:OwnerId;preload:false;"`
	LikeCount       int                `json:"like_count" gorm:"column:like_count;default:0"`
}

func (Restaurants) TableName() string { return "restaurants" }

func (r *Restaurants) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DBTypeRestaurant)
	if u := r.User; u != nil {
		u.Mask(false)
	}
}

type RestaurantsCreate struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Addr            string `json:"addr" gorm:"column:addr;"`
	OwnerId         int    `json:"-" gorm:"column:owner_id"`
}

func (data *RestaurantsCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

func (data *RestaurantsCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DBTypeRestaurant)
}

func (RestaurantsCreate) TableName() string { return Restaurants{}.TableName() }

type RestaurantsUpdate struct {
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
	Logo string `json:"logo" gorm:"column:logo"`
}

func (RestaurantsUpdate) TableName() string { return Restaurants{}.TableName() }

var (
	ErrNameIsEmpty = errors.New("name cannot be empty")
)
