package middleware

import (
	"errors"
	"greport/common"
	"greport/component/appctx"
	"greport/component/tokenprovider/paseto"
	usermodel "greport/module/user/model"
	userstorage "greport/module/user/storage"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractTokenFromHeaderString(tokenString string) (string, error) {
	splitStr := strings.Split(tokenString, " ")
	if len(splitStr) < 2 || splitStr[0] != "Bearer" || strings.TrimSpace(splitStr[1]) == "" {
		return "", ErrWrongAuthHeader
	}
	return splitStr[1], nil
}

func RequireAuth(appctx appctx.AppContext) gin.HandlerFunc {
	tokenProvider := paseto.NewPasetoProvider(appctx.GetSecretKey())
	// tokenProvider := jwt.NewTokenJwtProvider(appctx.GetSecretKey())
	return func(ctx *gin.Context) {
		token, err := extractTokenFromHeaderString(ctx.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}
		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}
		db := appctx.GetMainDbConn()
		userStore := userstorage.NewSQLStore(db)
		user, err := userStore.FindUser(ctx.Request.Context(), map[string]interface{}{"id": payload.UserId})
		if err != nil {
			panic(err)
		}
		if user.Status == int(usermodel.StatusDeleted) {
			panic(common.ErrNoPermision(errors.New("user has been deleted")))
		}
		user.Mask(false)
		ctx.Set(common.CurrentUser, user)
		ctx.Next()
	}
}

var (
	ErrWrongAuthHeader = common.NewCustomError(
		errors.New("wrong authenticate header"),
		"wrong authenticate header",
		"ErrWrongAuthHeader",
	)
)
