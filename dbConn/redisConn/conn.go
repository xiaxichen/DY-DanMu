package redisConn

import (
	"DY-DanMu/dbConn/config"
	"github.com/go-redis/redis/v7"
)

var _client *redis.Client

func init() {
	NewClient(config.RedisAddr, config.RedisPWD, config.RedisDB)
}

func NewClient(addr, PWD string, DBNum int) {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: PWD,   // no password set
		DB:       DBNum, // use default DB
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
	_client = RedisClient
}
func DBConn() *redis.Client {
	if _client != nil {
		return _client
	}
	NewClient(config.RedisAddr, config.RedisPWD, config.RedisDB)
	return _client
}
