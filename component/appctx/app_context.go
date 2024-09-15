package appctx

import (
	"greport/common"

	"github.com/ClickHouse/clickhouse-go/v2"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDbConn() *gorm.DB
	GetClickHouseConn() clickhouse.Conn
	GetSecretKey() string
	GetAppConfig() *common.Config
}

type appCtx struct {
	db         *gorm.DB
	clickhouse clickhouse.Conn
	secretKey  string
	// smsConfig map[string]string
	appConfig *common.Config
}

// App contains db connection, secret key, app config ...etc
func NewAppCtx(db *gorm.DB, clickhouse clickhouse.Conn, secretKey string, appConfig *common.Config) *appCtx {
	return &appCtx{
		db:         db,
		clickhouse: clickhouse,
		secretKey:  secretKey,
		appConfig:  appConfig,
	}
}

func (ctx *appCtx) GetMainDbConn() *gorm.DB {
	return ctx.db
}
func (ctx *appCtx) GetClickHouseConn() clickhouse.Conn {
	return ctx.clickhouse
}
func (ctx *appCtx) GetSecretKey() string {
	return ctx.secretKey
}
func (ctx *appCtx) GetAppConfig() *common.Config {
	return ctx.appConfig
}
