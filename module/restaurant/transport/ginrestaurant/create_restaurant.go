package ginrestaurant

import (
	"2ndbrand-api/common"
	"2ndbrand-api/component/appctx"
	restaurantbiz "2ndbrand-api/module/restaurant/biz"
	restaurantmodel "2ndbrand-api/module/restaurant/model"
	restaurantstorage "2ndbrand-api/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestUser := ctx.MustGet(common.CurrentUser).(common.Requester)
		db := appCtx.GetMainDbConn()
		var data restaurantmodel.RestaurantsCreate    // data model
		if err := ctx.ShouldBind(&data); err != nil { // get request có lỗi
			panic(common.ErrInvalidRequest(err))
		}
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		data.OwnerId = requestUser.GetUserID()
		if err := biz.CreateRestaurant(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(true)
		ctx.JSON(http.StatusOK, common.SimpleSuccessRes(data.FakeId.String()))
	}
}
