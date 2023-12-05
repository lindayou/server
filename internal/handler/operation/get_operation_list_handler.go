package operation

import (
	"net/http"
	"server/response"

	"server/internal/logic/operation"
	"server/internal/svc"
	"server/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetOperationListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetOperationListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := operation.NewGetOperationListLogic(r.Context(), svcCtx)
		resp, err := l.GetOperationList(&req)
		response.Response(w, resp, err)
	}
}
