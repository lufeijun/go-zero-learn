package redis

import (
	"context"

	"demo/gozero115/internal/svc"
	"demo/gozero115/internal/types"
	"demo/gozero115/tool/funcs"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetLogic {
	return &SetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetLogic) Set(req *types.RedisSet) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	resp = funcs.ResponseInit()

	key := req.Key
	value := req.Value
	if key == "" {
		resp.Message = "key is empty"
		return
	}

	err = l.svcCtx.RedisClient.Set(key, value)
	if err != nil {
		resp.Message = err.Error()
		return
	}

	return
}
