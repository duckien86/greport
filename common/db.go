package common

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LoadDbCnn(dbType string, IsDebugMode bool) (*gorm.DB, error) {
	databaseURI := ""
	var db *gorm.DB
	var err error
	switch dbType {
	case DbMysql:
		dbType = strings.ToUpper(dbType)
		databaseURI = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv(dbType+"_"+DbUsername),
			os.Getenv(dbType+"_"+DbPassword),
			os.Getenv(dbType+"_"+DbHost),
			os.Getenv(dbType+"_"+DbPort),
			os.Getenv(dbType+"_"+DbName))

		if db, err = gorm.Open(mysql.Open(databaseURI), &gorm.Config{}); err != nil {
			log.Print(databaseURI)
			log.Fatal(err)
		}

	case DbClickhouse:
		databaseURI = fmt.Sprintf("http://%s:%s@%s:%s/%s?dial_timeout=10s&read_timeout=20s",
			os.Getenv(dbType+"_"+DbUsername),
			os.Getenv(dbType+"_"+DbPassword),
			os.Getenv(dbType+"_"+DbHost),
			os.Getenv(dbType+"_"+DbPort),
			os.Getenv(dbType+"_"+DbName))

		if db, err = gorm.Open(clickhouse.Open(databaseURI), &gorm.Config{}); err != nil {
			// if db, err = gorm.Open(clickhouse.Open(databaseURI), &gorm.Config{}); err != nil {
			log.Print(databaseURI)
			log.Fatal(err)
		}

	}

	if IsDebugMode { // set debug mode
		db = db.Debug()
	}
	log.Printf("%s DB connected \n %s", dbType, databaseURI)

	return db, nil
}
