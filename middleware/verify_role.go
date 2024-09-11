package middleware

import (
	"2ndbrand-api/common"
	"2ndbrand-api/component/appctx"
	"errors"

	"github.com/gin-gonic/gin"
)

func VerifyRole(appCtx appctx.AppContext, allowRoles ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)
		flag := false
		for _, v := range allowRoles {
			if user.GetRole() == v {
				flag = true
				break
			}
		}
		if !flag {
			panic(common.ErrNoPermision(errors.New("you have no permision")))
		}
		c.Next()
	}
}
