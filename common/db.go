package common

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ClickHouse/clickhouse-go/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const ( // define key name
	DbHost     = "HOSTNAME"
	DbPort     = "PORT"
	DbUsername = "USERNAME"
	DbPassword = "PASSWORD"
	DbName     = "NAME"
)

func GetMysqlCnn(IsDebugMode bool) (*gorm.DB, error) {
	databaseURI := ""
	var db *gorm.DB
	var err error

	dbType := strings.ToUpper(DbMysql)
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

	if IsDebugMode { // set debug mode
		db = db.Debug()
	}
	log.Printf("%s DB connected \n %s", dbType, databaseURI)

	return db, nil
}

// Get ClickHouse DB Connection
func GetClickHouseCnn(IsDebugMode bool) (clickhouse.Conn, error) {
	dbType := strings.ToUpper(DbClickhouse)
	url := fmt.Sprintf("%s:%s", os.Getenv(dbType+"_"+DbHost), os.Getenv(dbType+"_"+DbPort))
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{url},
		Auth: clickhouse.Auth{
			Database: os.Getenv(dbType + "_" + DbName),
			Username: os.Getenv(dbType + "_" + DbUsername),
			Password: os.Getenv(dbType + "_" + DbPassword),
		},
		Protocol: clickhouse.Native,
		// DialTimeout:      5 * time.Second,
		// ConnMaxLifetime:  time.Hour,
		// ConnOpenStrategy: clickhouse.ConnOpenRoundRobin,
	})
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, err
	}
	return conn, nil
}
