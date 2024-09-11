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

func UpdateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDbConn()
		var updateData restaurantmodel.RestaurantsUpdate
		// id, err := strconv.Atoi(ctx.Param("id"))
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := ctx.ShouldBind(&updateData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(db)
		requester := ctx.MustGet(common.CurrentUser).(common.Requester)
		biz := restaurantbiz.NewUpdateRestaurantBiz(store, requester)
		if err := biz.UpdateRestaurant(ctx, &updateData, int(uid.GetLocalID())); err != nil {
			panic(common.ErrCannotUpdateEntity(restaurantmodel.EntityName, err))
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessRes(true))
	}
}
