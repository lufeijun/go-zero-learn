package redis

import (
	"context"

	"demo/gozero115/internal/svc"
	"demo/gozero115/internal/types"
	"demo/gozero115/tool/funcs"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.RedisGet) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	resp = funcs.ResponseInit()

	key := req.Name
	if key == "" {
		key = "name"
	}

	resp.Data = key

	return
}
