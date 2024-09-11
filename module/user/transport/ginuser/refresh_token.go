package ginuser

import (
	"2ndbrand-api/common"
	"2ndbrand-api/component/appctx"
	"2ndbrand-api/component/tokenprovider/paseto"
	userbiz "2ndbrand-api/module/user/biz"
	usermodel "2ndbrand-api/module/user/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// refresh token
func RefreshToken(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dataModel usermodel.RefreshToken
		if err := ctx.ShouldBind(&dataModel); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		tokenProvider := paseto.NewPasetoProvider(appCtx.GetSecretKey())
		refreshTokenBiz := userbiz.NewRefreshTokenBiz(tokenProvider)
		token, err := refreshTokenBiz.RefreshToken(ctx.Request.Context(), dataModel.OldToken)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessRes(token))
	}
}
