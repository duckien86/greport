package ginuser

import (
	"greport/common"
	"greport/component/appctx"
	"greport/component/hasher"
	"greport/component/verifier"
	userbiz "greport/module/user/biz"
	usermodel "greport/module/user/model"
	userstorage "greport/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterUser
// @BasePath /
// @Summary Register user
// @Tags User
// @Schemes  http
// @Accept json
// @Produce json
//
//	@Param	request-data	body	usermodel.UserCreateReq	true	"Data request"
//
// @Success 200  {string}  string  "Verify id"
// @Failure 400  {string}  string  "Invalid request data"
// @Failure 500  {string}  string  "Internal server error"
// @Router  /users/register [post]
func Register(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDbConn() // get main db connection
		var data usermodel.UserCreate
		if err := ctx.ShouldBind(&data); err != nil { // bind data from request to struct
			panic(common.ErrInvalidRequest(err))
		}
		store := userstorage.NewSQLStore(db) // create new store
		sha256hash := hasher.New(hasher.TypeSha256)
		biz := userbiz.NewRegisterBiz(store, sha256hash)
		verifyId, err := biz.RegisterUser(ctx.Request.Context(), &data) // create new biz
		if err != nil {                                                 // register user
			panic(err)
		}
		// data.Mask(true)
		ctx.JSON(http.StatusOK, common.SimpleSuccessRes(verifyId))
		// ctx.JSON(http.StatusOK, common.SimpleSuccessRes(data.FakeId.String()))
	}
}

// Verify registration
// @title Verify registration
// @BasePath /
// @Summary Verify registration
// @Tags User
// @Accept json
// @Produce json
// @Param	data	body	verifier.VerifyRequest	true	"Verify request data"
// @Success 200  string  "User id"
// @Failure 400  string  "Invalid request data"
// @Failure 500  string  "Error details"
// @Router  /users/verify-registration [post]
func VerifyRegistration(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDbConn() // get main db connection
		var dataUser usermodel.UserCreate
		var verifyData verifier.VerifyRequest

		if err := ctx.ShouldBind(&verifyData); err != nil { // bind data from request to struct
			panic(common.ErrInvalidRequest(err))
		}
		sha256hash := hasher.New(hasher.TypeSha256)
		store := userstorage.NewSQLStore(db) // create new store
		biz := userbiz.NewRegisterBiz(store, sha256hash)
		err := biz.VerifyAndCreateUser(ctx.Request.Context(), &verifyData, &dataUser) // Verify and create user
		if err != nil {
			panic(err)
		}
		dataUser.Mask(true)
		ctx.JSON(http.StatusOK, common.SimpleSuccessRes(dataUser.FakeId.String()))
	}
}
