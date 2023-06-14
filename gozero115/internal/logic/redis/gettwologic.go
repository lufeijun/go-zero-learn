package redis

import (
	"context"

	"demo/gozero115/internal/svc"
	"demo/gozero115/internal/types"
	"demo/gozero115/tool/funcs"

	"github.com/zeromicro/go-zero/core/logx"
)

type GettwoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGettwoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GettwoLogic {
	return &GettwoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GettwoLogic) Gettwo(req *types.RedisGet2) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	resp = funcs.ResponseInit()

	key := req.Name

	if key == "" {
		key = "name"
	}

	value, err := l.svcCtx.RedisClient.Get(key)
	if err != nil {
		resp.Message = err.Error()
		return
	}

	resp.Data = value

	return
}
