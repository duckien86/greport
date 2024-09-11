package middleware

import (
	"2ndbrand-api/common"
	"2ndbrand-api/component/appctx"

	"github.com/gin-gonic/gin"
)

// Hàm bắt lỗi crash của ứng dụng
// trả về như 1 lỗi thông thường thông qua cơ chế panic và recover
func Recover(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.Header("Context-Type", "application/json")
				if appErr, ok := err.(*common.AppError); ok {
					ctx.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
					// return
				}
				appErr := common.ErrInternal(err.(error))
				ctx.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
				// return
			}
		}()
		ctx.Next()
	}
}
