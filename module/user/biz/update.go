package userbiz

import (
	"2ndbrand-api/common"
	usermodel "2ndbrand-api/module/user/model"
	"context"
	"errors"
)

type UpdateUserStore interface {
	FindUser(context context.Context, condition map[string]interface{}, moreKeys ...string) (*usermodel.Users, error)
	Update(context context.Context, updateData *usermodel.UserUpdate, id int) error
	UpdatePassword(context context.Context, updateData *usermodel.Users, id int) error
}

type updateUserBiz struct {
	store UpdateUserStore
}

func NewUpdateUserBiz(store UpdateUserStore) *updateUserBiz {
	return &updateUserBiz{store: store}
}

func (biz *updateUserBiz) UpdateUser(
	ctx context.Context,
	updateData *usermodel.UserUpdate,
	id int,
) error {
	oldData, err := biz.store.FindUser(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	// check status
	if oldData.Status == usermodel.StatusDeleted {
		return ErrAccountHasBeenDeleted
	}
	// check phone
	if oldData.Phone == updateData.Phone {
		return ErrPhoneIsExisted
	}
	// check email
	if oldData.Email == updateData.Email {
		return ErrEmailIsExisted
	}
	if err := biz.store.Update(ctx, updateData, id); err != nil {
		return err
	}
	return nil
}

type updateUserPwdBiz struct {
	store  UpdateUserStore
	hasher Hasher
}

func NewUserUpdatePasswordBiz(store UpdateUserStore, hasher Hasher) *updateUserPwdBiz {
	return &updateUserPwdBiz{store: store, hasher: hasher}
}

// Implement change password biz
func (biz *updateUserPwdBiz) UpdateUserPassword(
	ctx context.Context,
	updateData *usermodel.UserChangePasswordReq,
	id int) error {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	// check status
	if user.Status == usermodel.StatusDeleted {
		return ErrAccountHasBeenDeleted
	}
	// check password
	if len(updateData.Password) > 0 {
		oldPassword := biz.hasher.Hash(user.Salt + updateData.Password)
		if oldPassword != user.Password {
			return ErrUsernameOrPasswordInvalid
		}
	}
	// update password
	user.Salt = common.GenSalt(20)
	user.Password = biz.hasher.Hash(user.Salt + updateData.NewPassword)
	if err := biz.store.UpdatePassword(ctx, user, id); err != nil {
		return err
	}
	return nil
}

var (
	ErrPhoneIsExisted = common.NewCustomError(
		errors.New("phone is existed"),
		"phone is existed",
		"ErrPhoneIsExisted",
	)
	ErrEmailIsExisted = common.NewCustomError(
		errors.New("email is existed"),
		"email is existed",
		"ErrEmailIsExisted",
	)
	ErrUserNotActive = common.NewCustomError(
		errors.New("user is not active"),
		"user is not active",
		"ErrUserNotActive",
	)
)
