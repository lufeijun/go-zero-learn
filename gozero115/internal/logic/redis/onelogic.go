package redis

import (
	"context"
	"fmt"

	"demo/gozero115/internal/svc"
	"demo/gozero115/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type OneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OneLogic {
	return &OneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OneLogic) One() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	redisconfig := redis.RedisConf{
		Host: l.svcCtx.Config.JpRedisConfig.Host,
		Pass: l.svcCtx.Config.JpRedisConfig.Pass,
		Type: l.svcCtx.Config.JpRedisConfig.Type,
	}
	redisClient, err := redis.NewRedis(redisconfig)
	if err != nil {
		panic("redis 链接失败" + err.Error())
	}

	// time.Sleep(time.Second * 10)

	fmt.Println(&redisClient)

	str, err := redisClient.Get("name")
	if err != nil {
		panic("redis 链接失败" + err.Error())
	}

	fmt.Println(str)

	return
}
