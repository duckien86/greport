package main

import (
	"greport/component/appctx"
	"greport/middleware"
	reportcontroller "greport/module/report/controller"
	"greport/module/user/transport/ginuser"

	"github.com/gin-gonic/gin"
)

type route struct {
	version *gin.RouterGroup
	server  *gin.Engine
	appCtx  appctx.AppContext
}

// NewRoute : create new route
func NewRoute(verName string, server *gin.Engine, appCtx appctx.AppContext) *route {
	server.Use(middleware.Recover(appCtx)) // apply recover middleware
	return &route{
		version: server.Group(verName),
		server:  server,
		appCtx:  appCtx,
	}
}

// AddReport: Setup report module route
func (r *route) AddReport() {
	users := r.version.Group("/greport")
	users.GET("/ping", reportcontroller.Pong(r.appCtx))
	users.POST("/ping", reportcontroller.Pong(r.appCtx))
	users.POST("/msglog", reportcontroller.GetMsgLog(r.appCtx))
}

// Setup user module route
func (r *route) AddUser() {
	users := r.version.Group("/users")
	users.POST("/login", ginuser.Login(r.appCtx))
	users.POST("/refresh-token", ginuser.RefreshToken(r.appCtx))
	users.POST("/register", ginuser.Register(r.appCtx))
	users.POST("/verify-registration", ginuser.VerifyRegistration(r.appCtx))
	// require authentication
	users.Use(middleware.RequireAuth(r.appCtx))
	users.PATCH("/:id", ginuser.UpdateUser(r.appCtx))
	users.GET("/profile/:id", ginuser.GetProfile(r.appCtx))
	users.PATCH("/change-password/:id", ginuser.UpdateUserPassword(r.appCtx))
	// users.PATCH("/:id", middleware.RequireAuth(r.appCtx), ginuser.UpdateUser(r.appCtx))
	// users.GET("/profile/:id", middleware.RequireAuth(r.appCtx), ginuser.GetProfile(r.appCtx))
	// users.PATCH("/change-password/:id", middleware.RequireAuth(r.appCtx), ginuser.UpdateUserPassword(r.appCtx))
}

// Setup admin module route
func (r *route) AddAdmin(appCtx appctx.AppContext, version *gin.RouterGroup) {
	admin := r.version.Group("/admin",
		middleware.RequireAuth(r.appCtx),
		middleware.VerifyRole(r.appCtx, "admin"),
	)
	admin.POST("/login", ginuser.Login(r.appCtx))
}
