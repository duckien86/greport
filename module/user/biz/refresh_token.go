package userbiz

import (
	"context"
	"greport/component/tokenprovider"
)

type refreshTokenBiz struct {
	tokenProvider tokenprovider.Provider
}

func NewRefreshTokenBiz(tokenProvider tokenprovider.Provider) *refreshTokenBiz {
	return &refreshTokenBiz{tokenProvider: tokenProvider}
}

// refresh token
func (biz *refreshTokenBiz) RefreshToken(ctx context.Context, oldRefreshToken string) (*tokenprovider.Token, error) {
	newToken, err := biz.tokenProvider.RefreshToken(oldRefreshToken)
	if err != nil {
		return nil, err
	}
	return newToken, nil
}
