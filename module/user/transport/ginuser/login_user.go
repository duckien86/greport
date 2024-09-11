package ginuser

import (
	"2ndbrand-api/common"
	"2ndbrand-api/component/appctx"
	"2ndbrand-api/component/hasher"
	"2ndbrand-api/component/tokenprovider/paseto"
	userbiz "2ndbrand-api/module/user/biz"
	usermodel "2ndbrand-api/module/user/model"
	userstorage "2ndbrand-api/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

const ExpiryInSecond = 10

// const ExpiryInSecond = 24 * 60 * 60

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
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
