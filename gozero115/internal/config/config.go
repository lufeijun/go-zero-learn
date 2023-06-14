package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	MysqlDsn string

	Jipeng string

	JpRedisConfig JpRedis

	// RedisClient *redis.Redis
}

type JpRedis struct {
	Host string
	Pass string
	Type string
}
