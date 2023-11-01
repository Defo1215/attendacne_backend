package database

import (
	"github.com/redis/go-redis/v9"
	"time"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",               // 没有密码，默认值
		DB:           0,                // 默认DB 0
		DialTimeout:  10 * time.Second, // 连接超时时间
		ReadTimeout:  30 * time.Second, // 读超时，默认3秒， -1表示取消读超时
		WriteTimeout: 30 * time.Second, // 写超时，默认等于读超时
		PoolSize:     10,               // 连接池大小，默认为10个连接
		PoolTimeout:  30 * time.Second, // 连接池超时时间，默认为4秒
	})
}

func GetRedis() *redis.Client {
	return RedisClient
}
