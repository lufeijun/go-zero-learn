package logic

import (
	"context"
	"fmt"
	"time"

	"demo/gozero115/internal/svc"
	"demo/gozero115/internal/types"
	"demo/gozero115/tool/funcs"

	"github.com/zeromicro/go-zero/core/logx"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	fmt.Println("NewIndexLogic")
	return &IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func init() {
	fmt.Println("IndexLogic init")
}

func (l *IndexLogic) Index() (resp *types.Response, err error) {

	resp = funcs.ResponseInit()

	resp.Message = "hello world gozero"
	resp.Data = time.Now().Format("2006-01-02 15:04:05")

	return
}
