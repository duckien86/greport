package ginreport

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
		db := appctx.GetMainDbConn()
		var reqData reportmodel.MsgLogReq
		if err := ctx.ShouldBind(&reqData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := reportstorage.NewSQLStore(db)
		// tokeProvider := jwt.NewTokenJwtProvider(appCtx.GetSecretKey())
		biz := reportbiz.NewReportBiz(store)
		data, err := biz.GetMsgLog(ctx.Request.Context(), reqData)
		log.Print(data)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessRes(data))

	}
}
