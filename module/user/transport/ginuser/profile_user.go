package ginuser

import (
	"2ndbrand-api/common"
	"2ndbrand-api/component/appctx"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfile(appctx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet(common.CurrentUser).(common.Requester)
		// TODO: profile of the user must be more complex. Like hobbies, address, etc.
		// payment info, etc.
		ctx.JSON(http.StatusAccepted, common.SimpleSuccessRes(user))
	}
}
