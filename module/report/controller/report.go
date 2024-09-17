package reportcontroller

import (
	"greport/common"
	"greport/component/appctx"
	reportbiz "greport/module/report/biz"
	reportmodel "greport/module/report/model"
	reportstorage "greport/module/report/storage"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pong(appctx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, common.SimpleSuccessRes("pong"))
	}
}

func GetMsgLog(appctx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqData reportmodel.MsgLogFilter
		var paging common.Paging
		dbConn := appctx.GetClickHouseConn()

		if err := ctx.ShouldBind(&reqData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := ctx.ShouldBindQuery(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()
		if details, err := common.ValidateStruct(reqData); err != nil {
			panic(common.ErrValidationData(err, details))
		}
		log.Println(reqData)
		store := reportstorage.NewSQLStore(dbConn)
		biz := reportbiz.NewReportBiz(store)
		data, err := biz.GetMsgLog(ctx.Request.Context(), &reqData, &paging)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.NewSuccessRes(data, nil, nil))

	}
}
