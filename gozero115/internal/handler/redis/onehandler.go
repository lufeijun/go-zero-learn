package redis

import (
	"net/http"

	"demo/gozero115/internal/logic/redis"
	"demo/gozero115/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OneHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := redis.NewOneLogic(r.Context(), svcCtx)
		resp, err := l.One()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
