package ginuser_test

import (
	"2ndbrand-api/common"
	"2ndbrand-api/component/appctx"
	"2ndbrand-api/component/hasher"
	userbiz "2ndbrand-api/module/user/biz"
	usermodel "2ndbrand-api/module/user/model"
	userstorage "2ndbrand-api/module/user/storage"
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initTest() appctx.AppContext {
	rootPath := "/usr/local/git_source/2ndbrand/2ndbrand-api/"
	appConfig := common.NewConfig(rootPath, "config.yml")
	appConfig.Load()
	dbCnnStr := appConfig.GetDbCnnStr()
	if dbCnnStr == "" {
		log.Println("Check file [config.yml] :: [MYSQL_CONN_STRING]")
	}
	db, err := gorm.Open(mysql.Open(dbCnnStr), &gorm.Config{})
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
