package ginrestaurant

import (
	"2ndbrand-api/common"
	"2ndbrand-api/component/appctx"
	restaurantbiz "2ndbrand-api/module/restaurant/biz"
	restaurantstorage "2ndbrand-api/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteRestaurant(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// id, err := strconv.Atoi(ctx.Param("id"))
		uid, err := common.FromBase58(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := appCtx.GetMainDbConn()
		store := restaurantstorage.NewSQLStore(db)
		requester := ctx.MustGet(common.CurrentUser).(common.Requester)
		biz := restaurantbiz.NewDeleteRestaurantBiz(store, requester)
		if err := biz.DeleteRestaurant(ctx, int(uid.GetLocalID())); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessRes(true))
	}
}
