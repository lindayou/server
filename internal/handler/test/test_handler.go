package test

import (
	"net/http"
	"server/response"

	"server/internal/logic/test"
	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func TestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TestReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := test.NewTestLogic(r.Context(), svcCtx)
		resp, err := l.Test(&req)
		response.Response(w, resp, err)
	}
}
