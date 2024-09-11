package myredis

import (
	"2ndbrand-api/common"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type MyRedis interface {
	GenKey(serviceName, action string, params ...string) string
	NewClient(dbName string) *redis.Client
}

const (
	DB_DEFAULT = 0
	DB_USER    = 1 // use for all user service
	DB_ORDER   = 2 // use for all order service
)

func NewClient(dbName int) *redis.Client {
	if dbName < 0 {
		dbName = DB_DEFAULT
	}
	host := common.GetEnv(common.RedisHost, "localhost")
	port := common.GetEnv(common.RedisPort, "6379")
	pwd := common.GetEnv(common.RedisPassword, "")
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port), // Default address
		Password: pwd,                              // No password by default
		DB:       dbName,                           // Default DB
	})
}

func GenKey(serviceName, action string, params ...string) (string, error) {
	if len(serviceName) < 1 || len(action) < 1 {
		return "", errors.New("serviceName and action must be not empty")
	}
	key := serviceName + "::" + action
	for _, p := range params {
		key += "::" + p
	}
	return key, nil
}
