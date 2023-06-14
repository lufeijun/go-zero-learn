package logger

import (
	"context"

	"demo/gozero115/internal/svc"
	"demo/gozero115/internal/types"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IndexLogic) Index() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	// {"@timestamp":"2023-06-13T17:01:11.972+08:00","caller":"logger/indexlogic.go:30","content":"hello world logx","level":"info"}
	logx.Info("hello world logx")

	// {"@timestamp":"2023-06-13T17:01:11.972+08:00","caller":"logger/indexlogic.go:32","content":"hello world logc","level":"info","span":"aef96d4694a136c2","trace":"92561e82fed48701496ca680ac441427"}
	logc.Info(l.ctx, "hello world logc")

	logc.Error(l.ctx, "hello world logc")

	return
}
