package ginuser

import (
	"greport/common"
	"greport/component/appctx"
	"greport/component/hasher"
	userbiz "greport/module/user/biz"
	usermodel "greport/module/user/model"
	userstorage "greport/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDbConn()
		var data usermodel.UserUpdate // data model
		uid, err := common.FromBase58(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := ctx.ShouldBind(&data); err != nil { // get request có lỗi
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewUpdateUserBiz(store)
		if err := biz.UpdateUser(ctx.Request.Context(), &data, int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessRes(true))
	}
}

func UpdateUserPassword(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDbConn()
		var updateData usermodel.UserChangePasswordReq

		if err := ctx.ShouldBind(&updateData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		currentUser := ctx.MustGet(common.CurrentUser).(common.Requester)
		store := userstorage.NewSQLStore(db)
		hasher := hasher.New(hasher.TypeSha256)
		biz := userbiz.NewUserUpdatePasswordBiz(store, hasher)
		if err := biz.UpdateUserPassword(ctx.Request.Context(), &updateData, currentUser.GetUserID()); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessRes(true))
	}
}
