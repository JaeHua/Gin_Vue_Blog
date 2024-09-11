package model

import "github.com/go-redis/redis"

// redis对象结构体

type Redis struct {
	RedisHost string
	RedisPort string
	RedisAuth string
	Database  int
	Db        *redis.Client
}
