package user

import (
	"net/http"

	"demo/gozero115/internal/logic/user"
	"demo/gozero115/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func IndexHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewIndexLogic(r.Context(), svcCtx)
		resp, err := l.Index()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
