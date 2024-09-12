//go:generate swag init
package main

import (
	"greport/common"
	"greport/component/appctx"
	"greport/middleware"
	"greport/module/report/transport/ginreport"
	"greport/module/user/transport/ginuser"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	appConfig := common.NewConfig("./", "config.yml")
	appConfig.Load("app", "clickhouse")
	// dbCnnStr := appConfig.GetDbCnnStr(common.DbClickhouse) // get db connection string
	// if dbCnnStr == "" {
	// 	log.Println("Check file [config.yml] :: [db]")
	// 	return
	// }
	secretKey := appConfig.GetSecret() // get secret key
	if secretKey == "" {
		log.Println("Check file [config.yml] ::[app] [secret_key]")
		return
	}
	db, err := appConfig.LoadDbCnn(common.DbClickhouse)
	if err != nil {
		log.Fatal(err)
	}

	// It contains db connection, secret key, app config ...etc
	appCtx := appctx.NewAppCtx(db, secretKey, appConfig)

	// Create restAPI by GIN
	r := gin.Default()                // create new gin serve
	r.Use(middleware.Recover(appCtx)) // recover middleware

	// version api
	v1 := r.Group("/v1") // create new group
	// setupUserRoute(appCtx, v1)
	// setupAdminRoute(appCtx, v1)
	setupReportRoute(appCtx, v1)

	r.Run(":" + appConfig.GetAppPort()) // listen and serve on 0.0.0.0:{port}
	// http.ListenAndServe(":"+appConfig.GetAppPort(), r)
}

func SetupAdminRoute(appCtx appctx.AppContext, version *gin.RouterGroup) {
	admin := version.Group("/admin",
		middleware.RequireAuth(appCtx),
		middleware.VerifyRole(appCtx, "admin"),
	)
	admin.POST("/login", ginuser.Login(appCtx))
}

func SetupUserRoute(appCtx appctx.AppContext, version *gin.RouterGroup) {
	users := version.Group("/users")
	users.POST("/login", ginuser.Login(appCtx))
	users.POST("/refresh-token", ginuser.RefreshToken(appCtx))
	users.POST("/register", ginuser.Register(appCtx))
	users.POST("/verify-registration", ginuser.VerifyRegistration(appCtx))
	users.PATCH("/:id", middleware.RequireAuth(appCtx), ginuser.UpdateUser(appCtx))
	users.GET("/profile/:id", middleware.RequireAuth(appCtx), ginuser.GetProfile(appCtx))
	users.PATCH("/change-password/:id", middleware.RequireAuth(appCtx), ginuser.UpdateUserPassword(appCtx))
	//	otherGroup := version.Group("/otherGroup")
	// add more route for other group
}
func setupReportRoute(appCtx appctx.AppContext, version *gin.RouterGroup) {
	users := version.Group("/greport")
	users.GET("/ping", ginreport.Pong(appCtx))
	users.POST("/ping", ginreport.Pong(appCtx))
	users.POST("/msglog", ginreport.GetMsgLog(appCtx))
}
