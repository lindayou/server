package dictionary

import (
	"net/http"
	"server/response"

	"server/internal/logic/dictionary"
	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteDicDetailsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteDicDetailsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dictionary.NewDeleteDicDetailsLogic(r.Context(), svcCtx)
		resp, err := l.DeleteDicDetails(&req)
		response.Response(w, resp, err)
	}
}
