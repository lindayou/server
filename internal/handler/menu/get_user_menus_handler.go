package menu

import (
	"net/http"
	"server/response"

	"server/internal/logic/menu"
	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserMenusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserMenusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewGetUserMenusLogic(r.Context(), svcCtx)
		resp, err := l.GetUserMenus(&req)
		response.Response(w, resp, err)
	}
}
