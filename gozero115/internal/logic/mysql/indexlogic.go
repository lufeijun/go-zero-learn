package mysql

import (
	"context"

	"demo/gozero115/internal/svc"
	"demo/gozero115/internal/types"
	"demo/gozero115/model"
	"demo/gozero115/tool/funcs"

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

// goctl model mysql datasource -url="root:123456@tcp(192.168.0.121:3306)/go-zero" -table="user"  -dir="./gozero115/model"

func (l *IndexLogic) Index() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = funcs.ResponseInit()

	var u model.User
	query := "select * from user where id=?"
	err = l.svcCtx.MysqlConn.QueryRowCtx(context.Background(), &u, query, 1)
	if err != nil {
		resp.Message = err.Error()
		return
	}

	resp.Data = u

	return
}
