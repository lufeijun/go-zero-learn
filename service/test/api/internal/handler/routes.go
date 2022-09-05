// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	test09 "demo/service/test/api/internal/handler/test09"
	"demo/service/test/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Errormiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/redis",
					Handler: test09.RedisHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/error/first",
					Handler: test09.ErrorHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/gorm/get",
					Handler: test09.GormHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/test09"),
	)
}
