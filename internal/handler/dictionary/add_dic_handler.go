package dictionary

import (
	"net/http"
	"server/response"

	"server/internal/logic/dictionary"
	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddDicHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddDicReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dictionary.NewAddDicLogic(r.Context(), svcCtx)
		resp, err := l.AddDic(&req)
		response.Response(w, resp, err)
	}
}
