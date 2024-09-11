package usermodel

import (
	"2ndbrand-api/common"
	"errors"
)

const StatusActive = 10
const StatusDeleted = 0
const EntityName = "User"
const RoleName = "user"

type Users struct {
	common.SQLModel `json:",inline"`
	Username        string `json:"username" gorm:"column:username"`
	Password        string `json:"-" gorm:"column:password"`
	Email           string `json:"email" gorm:"column:email"`
	Phone           string `json:"phone" gorm:"column:phone"`
	Last_name       string `json:"last_name" gorm:"column:last_name"`
	First_name      string `json:"first_name" gorm:"column:first_name"`
	Salt            string `json:"-" gorm:"column:salt" `
	Role            string `json:"-" gorm:"column:role"`
	// Avatar          *common.Image `json:"avatar" gorm:"column:avatar"`
}

func (Users) TableName() string { return "users" }

func (u *Users) Mask(isAdminOrOwner bool) {
	u.GenUID(common.DBTypeUser)
}

func (u *Users) GetUserID() int {
	return u.Id
}
func (u *Users) GetEmail() string {
	return u.Email
}
func (u *Users) GetRole() string {
	return u.Role
}

// @Schema(ignore=true)
type UserCreate struct {
	common.SQLModel `json:",inline"`
	Username        string `json:"username" gorm:"column:username" validate:"required,validPhone"`
	Verify          string `json:"verify_by" gorm:"-" validate:"oneof=sms email"` // confirm by sms or email
	Password        string `json:"password" gorm:"column:password" validate:""`   // TODO: validate password
	Salt            string `json:"-" gorm:"column:salt" `
	Role            string `json:"-" gorm:"column:role"` // role of user
	Phone           string `json:"-" gorm:"column:phone" validate:"excluded_with=Email validPhone"`
	Email           string `json:"-" gorm:"column:email" validate:"excluded_with=Phone email"`
	// Last_name       string `json:"last_name" gorm:"column:last_name" validate:"alphaunicode"`
	// First_name      string `json:"first_name" gorm:"column:first_name" validate:"alphaunicode"`
}

// @Schema(ignore=true)
type UserCreateReq struct {
	Username   string `json:"username" validate:"required,validPhone" example:"0914590038"`
	Last_name  string `json:"last_name" validate:"alphaunicode" example:"Nguyen"`
	First_name string `json:"first_name"  validate:"alphaunicode" example:"Van"`
	Verify     string `json:"verify_by"  example:"sms"` // confirm by sms or email
}

func (UserCreate) TableName() string {
	return Users{}.TableName()
}

func (u *UserCreate) Mask(isAdminOrOwner bool) {
	u.GenUID(common.DBTypeUser)
}

type UserLogin struct {
	Username string `json:"username" gorm:"column:username" form:"username"`
	Password string `json:"password" gorm:"column:password" form:"password"`
}
type RefreshToken struct {
	OldToken string `json:"old_token" gorm:"-" form:"old_token"`
}

type UserUpdate struct {
	Last_name  string `json:"last_name" gorm:"column:last_name" validate:"required,alphaunicode"`
	First_name string `json:"first_name" gorm:"column:first_name" validate:"required,alphaunicode"`
	Email      string `json:"email" gorm:"column:email" validate:"email"`
	Phone      string `json:"phone" gorm:"column:phone" validate:"required,cphonenum"`
}

type UserChangePasswordReq struct {
	Password    string `json:"password" gorm:"column:password" `
	NewPassword string `json:"new_password" gorm:"column:password" validate:"required"`
}

func (UserUpdate) TableName() string {
	return Users{}.TableName()
}

var (
	ErrUserNameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password is invalid"),
		"username or password is invalid",
		"ErrUserNameOrPasswordInvalid",
	)
	ErrEmailIsExisted = common.NewCustomError(errors.New(
		"email has already existed"),
		"email has already existed",
		"ErrEmailIsExisted",
	)

	ErrPhoneIsExisted = common.NewCustomError(errors.New(
		"phone number has already existed"),
		"phone number has already existed",
		"ErrPhoneIsExisted",
	)
	ErrPhoneNumberInvalid = common.NewCustomError(errors.New(
		"phone number invalid"),
		"phone number invalid",
		"ErrPhoneNumberInvalid",
	)
)
