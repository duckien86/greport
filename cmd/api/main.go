//go:generate swag init
package main

import (
	"greport/common"
	"greport/component/appctx"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	appConfig := common.NewConfig("./config/", "config.dev.yml")
	// appConfig := common.NewConfig("./config/", "config.dev.yml")
	appConfig.Load("app", "clickhouse")

	secretKey := appConfig.GetSecret() // get secret key
	if secretKey == "" {
		log.Println("Check file [config.yml] ::[app] [secret_key]")
		return
	}
	// chCnn, err := common.GetClickHouseCnn(appConfig.IsDebugMode()) // Get clickhouse connection
	// if err != nil {
	// 	// log.Printf("fail to connect %w", err)
	// 	wrappedErr := fmt.Errorf("connect DB fail %w", err)
	// 	fmt.Println(wrappedErr)
	// 	log.Fatal("Exit")
	// }
	appCtx := appctx.NewAppCtx(nil, nil, secretKey, appConfig) // create app context
	server := gin.Default()                                    // create new gin serve
	r := NewRoute("v1", server, appCtx)                        // create route
	r.AddUser()                                                // add user module route
	r.AddReport()                                              // add report module route
	server.Run(":" + appConfig.GetAppPort())                   // start server
}
