package common

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const ( // define key name
	DbMysql          = "MYSQL"
	DbClickhouse     = "CLICKHOUSE"
	AppPort          = "APP_PORT"
	AppGinMode       = "APP_GIN_MODE"
	AppSecretKey     = "APP_SECRET_KEY"
	DbHost           = "HOSTNAME"
	DbPort           = "PORT"
	DbUsername       = "USERNAME"
	DbPassword       = "PASSWORD"
	DbName           = "NAME"
	SmsApiKey        = "SMS_API_KEY"
	SmsUsername      = "SMS_USERNAME"
	SmsPassword      = "SMS_PASSWORD"
	RedisHost        = "HOST"
	RedisUsername    = "USERNAME"
	RedisPassword    = "PASSWORD"
	RedisPort        = "PORT"
	RabbitmqHost     = "HOST"
	RabbitmqPort     = "PORT"
	RabbitmqUsername = "USERNAME"
	RabbitmqPassword = "PASSWORD"
)

const (
	StorageTypeEnv   = "env"
	StorageTypeRedis = "redis"
)

type Config struct {
	dirPath  string
	fileName string
}

func NewConfig(dirPath, fileName string) *Config {
	return &Config{
		dirPath:  dirPath,
		fileName: fileName,
	}
}

func (c *Config) Load(serviceName ...string) {
	viper.SetConfigFile(c.dirPath + c.fileName)
	// viper.AddConfigPath(filepath)
	// viper.AutomaticEnv()
	// Đọc cấu hình từ file YAML
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}
	c.setEnvByGroup(serviceName...)
}

// Scan config.yml file and set to environment variable
func (c *Config) setEnvByGroup(serviceName ...string) {
	if len(serviceName) < 1 {
		fmt.Println("Error reading  @ ", serviceName)
		return
	}
	// Lấy giá trị cấu hình của service
	for _, name := range serviceName {
		configs := viper.GetStringMapString(name)
		for k, v := range configs {
			k = strings.ToUpper(name + "_" + k)
			if len(os.Getenv(k)) == 0 { // Ưu tiên biến môi trường
				os.Setenv(k, v)
				log.Printf("Set Env :: %s->%s ", k, v)
			} else {
				log.Printf("Ignore by default :: %s->%s ", k, os.Getenv(k))
			}
		}
	}
}

func (c *Config) GetDbCnnStr(dbType string) string {
	databaseURI := ""

	switch dbType {
	case DbMysql:
		databaseURI = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv(dbType+"_"+DbUsername),
			os.Getenv(dbType+"_"+DbPassword),
			os.Getenv(dbType+"_"+DbHost),
			os.Getenv(dbType+"_"+DbPort),
			os.Getenv(dbType+"_"+DbName))
	case DbClickhouse:
		databaseURI = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv(DbClickhouse+"_"+DbUsername),
			os.Getenv(DbClickhouse+"_"+DbPassword),
			os.Getenv(DbClickhouse+"_"+DbHost),
			os.Getenv(DbClickhouse+"_"+DbPort),
			os.Getenv(DbClickhouse+"_"+DbName))
	}
	return databaseURI
}

func (c *Config) LoadDbCnn(dbType string) (*gorm.DB, error) {
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
		// "clickhouse://gorm:gorm@localhost:9942/gorm?dial_timeout=10s&read_timeout=20s"
		//
		// dsn := "http://:@localhost:8123/default?dial_=10s&read_timeout=20s"
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

	if c.IsDebugMode() { // set debug mode
		db = db.Debug()
	}
	log.Printf("%s DB connected \n %s", dbType, databaseURI)

	return db, nil
}

func (c *Config) GetAppPort() string {
	port := os.Getenv(AppPort)
	if len(port) < 1 {
		port = DefaultPort
		log.Println("PORT is not declare. Set default by", port)
	}
	return port
}

func (c *Config) IsDebugMode() bool {
	return os.Getenv(AppGinMode) == "debug"
}

func (c *Config) GetSecret() string {
	return os.Getenv(AppSecretKey)
}

func (c *Config) GetSmsConfig() map[string]string {
	return map[string]string{
		"sms_api_key":  os.Getenv(SmsApiKey),
		"sms_username": os.Getenv(SmsUsername),
		"sms_password": os.Getenv(SmsPassword),
	}
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) < 1 {
		value = defaultValue
	}
	return value
}
