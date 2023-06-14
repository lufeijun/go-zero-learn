package redis

import (
	"net/http"

	"demo/gozero115/internal/logic/redis"
	"demo/gozero115/internal/svc"
	"demo/gozero115/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RedisGet
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := redis.NewGetLogic(r.Context(), svcCtx)
		resp, err := l.Get(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
