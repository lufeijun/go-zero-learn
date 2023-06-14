package svc

import (
	"demo/gozero115/internal/config"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	MysqlConn   sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {

	// redis 连接
	redisconfig := redis.RedisConf{
		Host: c.JpRedisConfig.Host,
		Pass: c.JpRedisConfig.Pass,
		Type: c.JpRedisConfig.Type,
	}
	redisClient, err := redis.NewRedis(redisconfig)
	if err != nil {
		panic("redis 链接失败" + err.Error())
	}

	// mysql 连接
	conn := sqlx.NewMysql(c.MysqlDsn)

	return &ServiceContext{
		Config:      c,
		RedisClient: redisClient,
		MysqlConn:   conn,
	}
}
