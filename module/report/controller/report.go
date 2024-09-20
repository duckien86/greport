package reportcontroller

import (
	"fmt"
	"greport/common"
	"greport/component/appctx"
	reportbiz "greport/module/report/biz"
	reportmodel "greport/module/report/model"
	reportstorage "greport/module/report/storage"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Pong: Test api connect
func Pong(appctx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, common.SimpleSuccessRes("pong"))
	}
}

// GetMsgLog: handler MsgLog request
func GetMsgLog(appctx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var filter reportmodel.MsgLogFilter
		var paging common.Paging
		// dbConn := appctx.GetClickHouseConn()
		dbConn, err := common.GetClickHouseCnn(appctx.GetAppConfig().IsDebugMode()) // Get clickhouse connection
		if err != nil {
			// log.Printf("fail to connect %w", err)
			wrappedErr := fmt.Errorf("connect DB fail %w", err)
			fmt.Println(wrappedErr)
			log.Fatal("Exit")
		}
		defer dbConn.Close() // close

		if err := ctx.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := ctx.ShouldBindQuery(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()
		if details, err := common.ValidateStruct(filter); err != nil {
			panic(common.ErrValidationData(err, details))
		}
		log.Println(filter)
		store := reportstorage.NewSQLStore(dbConn)
		biz := reportbiz.NewReportBiz(store)
		data, err := biz.GetMsgLogDetails(ctx.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.NewSuccessRes(data, paging, filter))

	}
}
