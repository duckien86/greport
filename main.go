//go:generate swag init
package main

import (
	"2ndbrand-api/common"
	"2ndbrand-api/component/appctx"
	"2ndbrand-api/component/rabbitmq/workqueues"
	"2ndbrand-api/middleware"
	"2ndbrand-api/module/user/transport/ginuser"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	appConfig := common.NewConfig("./", "config.yml")
	appConfig.Load()
	dbCnnStr := appConfig.GetDbCnnStr() // get db connection string
	if dbCnnStr == "" {
		log.Println("Check file [config.yml] :: [db]")
		return
	}
	secretKey := appConfig.GetSecret() // get secret key
	if secretKey == "" {
		log.Println("Check file [config.yml] ::[app] [secret_key]")
		return
	}
	db, err := gorm.Open(mysql.Open(dbCnnStr), &gorm.Config{}) // open db connection
	if err != nil {
		log.Fatal(err)
	}
	if appConfig.IsDebugMode() { // set debug mode
		db = db.Debug()
	}
	// Create app context. It will be used in all handlers.
	// It contains db connection, secret key, app config ...etc
	appCtx := appctx.NewAppCtx(db, secretKey, appConfig)
	// start async task
	go workqueues.StartConsumer("sms", "email")

	// Create restAPI by GIN
	r := gin.Default()                // create new gin serve
	r.Use(middleware.Recover(appCtx)) // recover middleware

	// version api
	v1 := r.Group("/v1") // create new group
	setupMainRoute(appCtx, v1)
	setupAdminRoute(appCtx, v1)
	r.Run(":" + appConfig.GetAppPort()) // listen and serve on 0.0.0.0:{port}
}

func setupAdminRoute(appCtx appctx.AppContext, version *gin.RouterGroup) {
	admin := version.Group("/admin",
		middleware.RequireAuth(appCtx),
		middleware.VerifyRole(appCtx, "admin"),
	)
	admin.POST("/login", ginuser.Login(appCtx))
}

func setupMainRoute(appCtx appctx.AppContext, version *gin.RouterGroup) {
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

// func setupSwagger(version *gin.RouterGroup) {
// 	group := version.Group("/docs")
// 	docs.SwaggerInfo.BasePath = "/v1/"
// 	group.GET("/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
// }
