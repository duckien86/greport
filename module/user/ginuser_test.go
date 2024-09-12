package ginuser_test

import (
	"context"
	"greport/common"
	"greport/component/appctx"
	"greport/component/hasher"
	userbiz "greport/module/user/biz"
	usermodel "greport/module/user/model"
	userstorage "greport/module/user/storage"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initTest() appctx.AppContext {
	rootPath := "/usr/local/git_source/2ndbrand/greport/"
	appConfig := common.NewConfig(rootPath, "config.yml")
	appConfig.Load()
	db, err := appConfig.LoadDbCnn(common.DbMysql)

	if err != nil {
		log.Fatal(err)
	}
	return appctx.NewAppCtx(db, appConfig.GetSecret(), appConfig)
}

func TestRegisterUser(t *testing.T) {
	appCtx := initTest()
	db := appCtx.GetMainDbConn() // get main db connection
	// generate data test
	data := usermodel.UserCreate{
		Username: "0123456789",
		Password: "123456",
		// Last_name:  "Nguyen",
		// First_name: "Van",
		Verify: "sms",
	}
	// verifier.SetSmsConfig(appConfig.GetSmsConfig()) // set sms config
	store := userstorage.NewSQLStore(db) // create new store
	sha256hash := hasher.New(hasher.TypeSha256)
	biz := userbiz.NewRegisterBiz(store, sha256hash)
	verifyId, err := biz.RegisterUser(context.Background(), &data) // create new biz

	assert.NoError(t, err)
	assert.NotEmpty(t, verifyId)
}
