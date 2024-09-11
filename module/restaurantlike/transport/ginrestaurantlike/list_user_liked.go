package ginrestaurantlike

import (
	"2ndbrand-api/common"
	"2ndbrand-api/component/appctx"
	restaurantlikebiz "2ndbrand-api/module/restaurantlike/biz"
	restaurantlikesmodel "2ndbrand-api/module/restaurantlike/model"
	restaurantlikestorage "2ndbrand-api/module/restaurantlike/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListUserLiked(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging
		uidRestaurant, err := common.FromBase58(ctx.Param("id"))
		filter := restaurantlikesmodel.Filter{
			RestaurantID: int(uidRestaurant.GetLocalID()),
		}
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()
		db := appCtx.GetMainDbConn()
		store := restaurantlikestorage.NewSQLStore(db)
		biz := restaurantlikebiz.NewListUserLikedBiz(store)
		result, err := biz.ListUserLiked(ctx, &filter, &paging)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.NewSuccessRes(result, paging, nil))
	}
}
