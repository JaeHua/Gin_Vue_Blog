package model

import (
	"backend/utils"
	"fmt"
	"github.com/go-redis/redis"
)

// redis对象结构体

type Redis struct {
	RedisHost string
	RedisPort string
	RedisAuth string
	Database  int
	Db        *redis.Client
}

// 全局变量
var RD *Redis

func InitRedis() {
	//配置实例
	rd := &Redis{
		RedisHost: utils.RdHost,
		RedisPort: utils.RdPort,
		RedisAuth: ""}
	//log.Println("rd.RedisHost", utils.RdHost)
	redisAddr := fmt.Sprintf("%s:%s", rd.RedisHost, rd.RedisPort)
	// Redis 连接对象: NewClient将客户端返回到由选项指定的Redis服务器
	rd.Db = redis.NewClient(&redis.Options{
		Addr:        redisAddr,    // redis服务ip:port
		Password:    rd.RedisAuth, // redis的认证密码
		DB:          rd.Database,  // 连接的database库
		IdleTimeout: 300,          // 默认Idle超时时间
		PoolSize:    100,          // 连接池
	})
	// 验证是否连接到redis服务端
	res, err := rd.Db.Ping().Result()
	RD = rd
	if err != nil {
		fmt.Printf("Connect Failed! Err: %v\n", err)
	} else {
		fmt.Printf("Connect Successful! Ping => %v\n", res)
	}
}

// GetRedis 暴露调用接口
func GetRedis() *Redis {
	return RD
}
