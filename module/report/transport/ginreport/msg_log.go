package ginreport

import (
	"greport/common"
	"greport/component/appctx"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pong(appctx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, common.SimpleSuccessRes("pong"))
	}
}
