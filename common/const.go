package common

const (
	DBTypeRestaurant = 1
	DBTypeUser       = 2
)

const DefaultPort = "8080"

type Requester interface {
	GetUserID() int
	GetEmail() string
	GetRole() string
}

const (
	CurrentUser = "user"
)

type UserPublic struct {
	SQLModel   `json:",inline"`
	Last_name  string `json:"last_name" gorm:"column:last_name"`
	First_name string `json:"first_name" gorm:"column:first_name"`
	Role       string `json:"role" gorm:"column:role"`
}

func (UserPublic) TableName() string {
	return "users"
}
func (u *UserPublic) Mask(isAdminOrOwner bool) {
	u.GenUID(DBTypeUser)
}
