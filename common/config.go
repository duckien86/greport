package common

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const ( // define key name
	AppPort          = "APP_PORT"
	AppGinMode       = "APP_GIN_MODE"
	AppSecretKey     = "APP_SECRET_KEY"
	DbHost           = "DB_HOST"
	DbPort           = "DB_PORT"
	DbUsername       = "DB_USERNAME"
	DbPassword       = "DB_PASSWORD"
	DbName           = "DB_NAME"
	SmsApiKey        = "SMS_API_KEY"
	SmsUsername      = "SMS_USERNAME"
	SmsPassword      = "SMS_PASSWORD"
	RedisHost        = "REDIS_HOST"
	RedisUsername    = "REDIS_USERNAME"
	RedisPassword    = "REDIS_PASSWORD"
	RedisPort        = "REDIS_PORT"
	RabbitmqHost     = "RABBITMQ_HOST"
	RabbitmqPort     = "RABBITMQ_PORT"
	RabbitmqUsername = "RABBITMQ_USERNAME"
	RabbitmqPassword = "RABBITMQ_PASSWORD"
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

func (c *Config) Load() {
	viper.SetConfigFile(c.dirPath + c.fileName)
	// viper.AddConfigPath(filepath)
	// viper.AutomaticEnv()
	// Đọc cấu hình từ file YAML
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}
	c.setEnvByGroup("app", "db", "sms")
}

// Scan config.yml file and set to environment variable
func (c *Config) setEnvByGroup(groupName ...string) {
	if len(groupName) < 1 {
		fmt.Println("Error reading  @ ", groupName)
		return
	}
	// Lấy giá trị cấu hình
	for _, name := range groupName {
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

func (c *Config) GetDbCnnStr() string {
	databaseURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv(DbUsername),
		os.Getenv(DbPassword),
		os.Getenv(DbHost),
		os.Getenv(DbPort),
		os.Getenv(DbName))

	return databaseURI
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
