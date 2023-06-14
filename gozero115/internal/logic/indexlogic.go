package logic

import (
	"context"
	"errors"
	"fmt"
	"time"
	"unsafe"

	"demo/gozero115/internal/svc"
	"demo/gozero115/internal/types"
	"demo/gozero115/tool/funcs"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
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
	fmt.Println("init")
}

func (l *IndexLogic) Index() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	// var a interface{}
	resp = funcs.ResponseInit()

	// time.Sleep(time.Second)

	// mrFinish()
	fmt.Println(unsafe.Sizeof(resp))
	fmt.Println(unsafe.Sizeof(*resp))

	resp.Data = "hello world gozero" + time.Now().Format("2006-01-02 15:04:05")

	return
}

func mrFinish() {
	err := mr.Finish(func() (err error) {
		time.Sleep(time.Millisecond * 3)
		return
	}, func() error {
		time.Sleep(time.Millisecond * 2)
		return errors.New("报错了")
	}, func() error {
		time.Sleep(time.Millisecond * 1)
		return nil
	})

	if err != nil {
		fmt.Println("ok")
	}
}
