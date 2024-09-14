package appctx

import (
	"greport/common"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDbConn() *gorm.DB
	GetSecretKey() string
	GetAppConfig() *common.Config
}

type appCtx struct {
	db        *gorm.DB
	secretKey string
	// smsConfig map[string]string
	appConfig *common.Config
}

// App contains db connection, secret key, app config ...etc
func NewAppCtx(db *gorm.DB, secretKey string, appConfig *common.Config) *appCtx {
	return &appCtx{
		db:        db,
		secretKey: secretKey,
		appConfig: appConfig,
	}
}

func (ctx *appCtx) GetMainDbConn() *gorm.DB {
	return ctx.db
}
func (ctx *appCtx) GetSecretKey() string {
	return ctx.secretKey
}
func (ctx *appCtx) GetAppConfig() *common.Config {
	return ctx.appConfig
}
