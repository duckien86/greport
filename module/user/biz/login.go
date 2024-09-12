package userbiz

import (
	"context"
	"errors"
	"greport/common"
	"greport/component/tokenprovider"
	usermodel "greport/module/user/model"
)

type LoginStorage interface {
	FindUser(context context.Context, condition map[string]interface{}, moreKeys ...string) (*usermodel.Users, error)
}

type loginBiz struct {
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBiz(
	storeUser LoginStorage,
	tokenProvider tokenprovider.Provider,
	hasher Hasher,
	expiry int,
) *loginBiz {
	return &loginBiz{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (biz *loginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	// 1. find user data by email
	user, _ := biz.storeUser.FindUser(ctx, map[string]interface{}{"username": data.Username})
	if user == nil {
		return nil, ErrUsernameOrPasswordInvalid
	}
	if user.Status == int(usermodel.StatusDeleted) {
		return nil, ErrAccountHasBeenDeleted
	}
	// 2. validate password
	hashPassword := biz.hasher.Hash(user.Salt + data.Password)
	if hashPassword != user.Password {
		return nil, ErrUsernameOrPasswordInvalid
	}
	payload := tokenprovider.TokenPayload{
		UserId:   user.Id,
		Username: user.Username,
		Role:     user.Role,
	}
	// 3 generate token
	token, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, err
	}
	// 4. return token
	return token, nil
}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)
	ErrAccountHasBeenDeleted = common.NewCustomError(
		errors.New("account has been deleted"),
		"account has been deleted",
		"ErrAccountHasBeenDeleted",
	)
)
