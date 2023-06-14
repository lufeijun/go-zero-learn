package redis

import (
	"net/http"

	"demo/gozero115/internal/logic/redis"
	"demo/gozero115/internal/svc"
	"demo/gozero115/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RedisSet
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := redis.NewSetLogic(r.Context(), svcCtx)
		resp, err := l.Set(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
