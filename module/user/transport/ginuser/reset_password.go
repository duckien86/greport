package ginuser

import (
	"greport/common"
	"greport/component/appctx"
	"greport/component/hasher"
	"greport/component/tokenprovider/paseto"
	userbiz "greport/module/user/biz"
	usermodel "greport/module/user/model"
	userstorage "greport/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

const ResetExpiryInMin = 10

// const ExpiryInSecond = 24 * 60 * 60

func ResetPasswordRequest(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDbConn()
		var dataModel usermodel.UserLogin
		if err := ctx.ShouldBind(&dataModel); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := userstorage.NewSQLStore(db)
		// tokeProvider := jwt.NewTokenJwtProvider(appCtx.GetSecretKey())
		tokeProvider := paseto.NewPasetoProvider(appCtx.GetSecretKey())
		sha256hash := hasher.New(hasher.TypeSha256)
		biz := userbiz.NewLoginBiz(store, tokeProvider, sha256hash, ExpiryInSecond)
		token, err := biz.Login(ctx.Request.Context(), &dataModel)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessRes(token))
	}
}
