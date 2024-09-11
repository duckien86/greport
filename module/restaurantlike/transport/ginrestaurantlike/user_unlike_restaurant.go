package ginrestaurantlike

import (
	"2ndbrand-api/common"
	"2ndbrand-api/component/appctx"
	restaurantstorage "2ndbrand-api/module/restaurant/storage"
	restaurantlikebiz "2ndbrand-api/module/restaurantlike/biz"
	restaurantlikesmodel "2ndbrand-api/module/restaurantlike/model"
	restaurantlikestorage "2ndbrand-api/module/restaurantlike/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserUnlikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uidRestaurant, err := common.FromBase58(ctx.Param("id"))
		if err != nil {
			panic(err)
		}
		requester := ctx.MustGet(common.CurrentUser).(common.Requester)
		dataModel := restaurantlikesmodel.Like{
			UserID:       requester.GetUserID(),
			RestaurantID: int(uidRestaurant.GetLocalID()),
		}
		db := appCtx.GetMainDbConn()
		store := restaurantlikestorage.NewSQLStore(db)
		drcLCS := restaurantstorage.NewSQLStore(db)
		biz := restaurantlikebiz.NewUnlikeRestaurantBiz(store, drcLCS)
		if err := biz.UnlikeRestaurant(ctx, &dataModel); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessRes(true))
	}
}
