package user

import (
	"net/http"

	"server/internal/logic/user"
	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func EditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EditUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewEditLogic(r.Context(), svcCtx)
		resp, err := l.Edit(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
