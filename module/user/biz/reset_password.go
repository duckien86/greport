package userbiz

import (
	"context"
	usermodel "greport/module/user/model"
)

type ResetPasswordStore interface {
	FindUser(context context.Context, condition map[string]interface{}, moreKeys ...string) (*usermodel.Users, error)
	Update(context context.Context, updateData *usermodel.UserUpdate, id int) error
	UpdatePassword(context context.Context, updateData *usermodel.Users, id int) error
}

type resetPasswordBiz struct {
	store ResetPasswordStore
}

func NewResetPasswordBiz(store ResetPasswordStore) *resetPasswordBiz {
	return &resetPasswordBiz{store: store}
}
