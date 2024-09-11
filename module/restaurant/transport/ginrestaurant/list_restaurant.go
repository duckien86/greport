package ginrestaurant

import (
	"2ndbrand-api/common"
	"2ndbrand-api/component/appctx"
	restaurantbiz "2ndbrand-api/module/restaurant/biz"
	restaurantmodel "2ndbrand-api/module/restaurant/model"
	restaurantrepo "2ndbrand-api/module/restaurant/repository"
	restaurantstorage "2ndbrand-api/module/restaurant/storage"
	restaurantlikestorage "2ndbrand-api/module/restaurantlike/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var filter restaurantmodel.Filter
		var paging common.Paging

		if err := ctx.ShouldBind(&filter); err != nil { // get request có lỗi
			panic(common.ErrInvalidRequest(err))
		}
		if err := ctx.ShouldBind(&paging); err != nil { // get request err
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill() // set default value
		db := appCtx.GetMainDbConn()
		store := restaurantstorage.NewSQLStore(db)
		storeLike := restaurantlikestorage.NewSQLStore(db)
		repo := restaurantrepo.NewListRestaurantRepo(store, storeLike)
		biz := restaurantbiz.NewListRestaurantBiz(repo)
		result, err := biz.ListRestaurant(ctx.Request.Context(), &filter, &paging)
		for i := range result {
			result[i].Mask(true)
		}
		if err != nil {
			panic(common.ErrCannotListEntity(restaurantmodel.EntityName, err))
		}
		ctx.JSON(http.StatusOK, common.NewSuccessRes(result, paging, filter))
	}
}
